package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gaspartv/go-tibia-info-back/internal/database"
	middlewares "github.com/gaspartv/go-tibia-info-back/internal/midleware"
	"github.com/gaspartv/go-tibia-info-back/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

type Env struct {
	DatabaseURL  string `validate:"required"`
	Port         string `validate:"required"`
	JwtSecretKey string `validate:"required"`
}

func validateEnv(env Env) error {
	validate := validator.New()
	if err := validate.Struct(env); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	var env Env
	env.DatabaseURL = os.Getenv("DATABASE_URL")
	env.Port = os.Getenv("PORT")
	env.JwtSecretKey = os.Getenv("JWT_SECRET_KEY")

	if err := validateEnv(env); err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userDB := database.NewUserDB(db)
	userService := service.NewUserService(*userDB)

	authService := service.NewAuthService(db)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           300,
	})
	c.Use(corsOptions.Handler)

	c.Post("/api-v1/user", userService.Create)
	c.With(middlewares.JwtMiddleware).Get("/api-v1/user", userService.Get)
	c.With(middlewares.JwtMiddleware).Delete("/api-v1/user/{id}", userService.Delete)
	c.With(middlewares.JwtMiddleware).Put("/api-v1/user/{id}", userService.Update)

	c.Post("/login", authService.Login)

	fmt.Println("Server is running on port", os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), c); err != nil {
		panic(err)
	}
}
