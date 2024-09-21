package entity

import (
	"time"

	"github.com/gaspartv/go-tibia-info-back/internal/util"
	"github.com/google/uuid"
)

type User struct {
	ID                   string     `json:"id"`
	Code                 string     `json:"code"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
	DeletedAt            *time.Time `json:"deleted_at"`
	DisabledAt           *time.Time `json:"disabled_at"`
	LastLoginAt          *time.Time `json:"last_login_at"`
	Type                 string     `json:"type"`
	Police               string     `json:"police"`
	FirstName            string     `json:"first_name"`
	LastName             string     `json:"last_name"`
	Email                string     `json:"email"`
	EmailHash            string     `json:"email_hash"`
	NationalId           string     `json:"national_id"`
	NationalIdHash       string     `json:"national_id_hash"`
	Telephone            string     `json:"telephone"`
	TelephoneHash        string     `json:"telephone_hash"`
	PasswordHash         string     `json:"password_hash"`
	BirthDate            string     `json:"birth_date,omitempty"`
	AvatarUri            string     `json:"avatar_uri,omitempty"`
	Language             string     `json:"language"`
	DarkMode             bool       `json:"dark_mode"`
	Permissions          []string   `json:"permissions"`
	IsVerified           bool       `json:"is_verified"`
	VerificationToken    string     `json:"verification_token,omitempty"`
	ResetPasswordToken   string     `json:"reset_password_token,omitempty"`
	LastPasswordChangeAt *time.Time `json:"last_password_change_at"`
	TwoFactorEnabled     bool       `json:"two_factor_enabled"`
}

type CreateUserDTO struct {
	Code        string   `json:"code"`
	Type        string   `json:"type"`
	Police      string   `json:"police"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Email       string   `json:"email"`
	NationalId  string   `json:"national_id"`
	Telephone   string   `json:"telephone"`
	Password    string   `json:"password"`
	BirthDate   string   `json:"birth_date,omitempty"`
	Language    string   `json:"language"`
	DarkMode    bool     `json:"dark_mode"`
	Permissions []string `json:"permissions"`
}

func NewUser(dto CreateUserDTO) *User {
	c := util.Crypto{}
	security := util.Security{}
	bcryptInstance := util.Bcrypt{}

	emailCrypto := c.Hash(dto.Email)
	emailSecurity, err := security.Encrypt(dto.Email)
	if err != nil {
		panic(err)
	}

	nationalIdCrypto := c.Hash(dto.NationalId)
	nationalIdSecurity, err := security.Encrypt(dto.NationalId)
	if err != nil {
		panic(err)
	}

	telephoneCrypto := c.Hash(dto.Telephone)
	telephoneSecurity, err := security.Encrypt(dto.Telephone)
	if err != nil {
		panic(err)
	}

	saltRounds := 10

	password := dto.Password
	passwordHash, err := bcryptInstance.Hash(password, saltRounds)

	if err != nil {
		panic(err)
	}

	return &User{
		ID:                   uuid.New().String(),
		Code:                 dto.Code,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
		DeletedAt:            nil,
		DisabledAt:           nil,
		LastLoginAt:          nil,
		Type:                 dto.Type,
		Police:               dto.Police,
		FirstName:            dto.FirstName,
		LastName:             dto.LastName,
		Email:                emailCrypto,
		EmailHash:            emailSecurity,
		NationalId:           nationalIdCrypto,
		NationalIdHash:       nationalIdSecurity,
		Telephone:            telephoneCrypto,
		TelephoneHash:        telephoneSecurity,
		PasswordHash:         passwordHash,
		BirthDate:            dto.BirthDate,
		AvatarUri:            "https://www.gravatar.com/avatar/",
		Language:             dto.Language,
		DarkMode:             dto.DarkMode,
		Permissions:          dto.Permissions,
		IsVerified:           false,
		VerificationToken:    "1243-1243-1243-1243",
		ResetPasswordToken:   "1243-1243-1243-1243",
		LastPasswordChangeAt: nil,
		TwoFactorEnabled:     false,
	}
}
