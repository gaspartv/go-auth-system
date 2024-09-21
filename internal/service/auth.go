package service

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gaspartv/go-tibia-info-back/internal/util"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gaspartv/go-tibia-info-back/internal/entity"
	handlerError "github.com/gaspartv/go-tibia-info-back/internal/handleError"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db *sql.DB
}

type TokenMessage struct {
	Token string `json:"token"`
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (handler *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	var auth entity.Auth

	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if auth.Email == "" {
		handlerError.Exec(w, "param email is required", http.StatusBadRequest)
		return
	}

	if auth.Password == "" {
		handlerError.Exec(w, "param password is required", http.StatusBadRequest)
		return
	}

	stmt, err := handler.db.Prepare("SELECT * FROM users WHERE email =?")
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer stmt.Close()

	c := util.Crypto{}
	row := stmt.QueryRow(c.Hash(auth.Email))

	var user entity.User
	if err := row.Scan(
		&user.ID,
		&user.Code,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
		&user.DisabledAt,
		&user.LastLoginAt,
		&user.Type,
		&user.Police,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.EmailHash,
		&user.NationalId,
		&user.NationalIdHash,
		&user.Telephone,
		&user.TelephoneHash,
		&user.PasswordHash,
		&user.BirthDate,
		&user.AvatarUri,
		&user.Language,
		&user.DarkMode,
		&user.Permissions,
		&user.IsVerified,
		&user.VerificationToken,
		&user.ResetPasswordToken,
		&user.LastPasswordChangeAt,
		&user.TwoFactorEnabled,
	); err != nil {
		fmt.Println("a: ", err)
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	decodedHash, err := base64.StdEncoding.DecodeString(user.PasswordHash)
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword(decodedHash, []byte(auth.Password))
	if err != nil {
		handlerError.Exec(w, "email or password incorrect", http.StatusBadRequest)
		return
	}

	claims := jwt.MapClaims{
		"username": user.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := TokenMessage{
		Token: tokenString,
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}
}
