package database

import (
	"database/sql"
	"encoding/json"
	"github.com/gaspartv/go-tibia-info-back/internal/entity"
)

type UserDB struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}

func (database *UserDB) Create(user *entity.User) (string, error) {
	stmt, err := database.db.Prepare("INSERT INTO users (" +
		"id," +
		"code," +
		"created_at," +
		"updated_at," +
		"deleted_at," +
		"disabled_at," +
		"last_login_at," +
		"type," +
		"police," +
		"first_name," +
		"last_name," +
		"email," +
		"email_hash," +
		"national_id," +
		"national_id_hash," +
		"telephone," +
		"telephone_hash," +
		"password_hash," +
		"birth_date," +
		"avatar_uri," +
		"language," +
		"dark_mode," +
		"permissions," +
		"is_verified," +
		"verification_token," +
		"reset_password_token," +
		"last_password_change_at," +
		"two_factor_enabled" +
		") VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		return "", err
	}
	defer stmt.Close()

	jsonPermissions, err := json.Marshal(user.Permissions)
	if err != nil {
		return "", err
	}

	if _, err := stmt.Exec(
		user.ID,
		user.Code,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
		user.DisabledAt,
		user.LastLoginAt,
		user.Type,
		user.Police,
		user.FirstName,
		user.LastName,
		user.Email,
		user.EmailHash,
		user.NationalId,
		user.NationalIdHash,
		user.Telephone,
		user.TelephoneHash,
		user.PasswordHash,
		user.BirthDate,
		user.AvatarUri,
		user.Language,
		user.DarkMode,
		jsonPermissions,
		user.IsVerified,
		user.VerificationToken,
		user.ResetPasswordToken,
		user.LastPasswordChangeAt,
		user.TwoFactorEnabled,
	); err != nil {
		return "", err
	}

	return "created successful", nil
}

func (database *UserDB) Delete(id string) (string, error) {
	stmt, err := database.db.Prepare("DELETE FROM users WHERE id =?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return "", err
	}

	return "deleted successful", nil
}

func (database *UserDB) Update(id string, user *entity.User) (string, error) {
	stmt, err := database.db.Prepare("UPDATE users SET name =?, email =?, password =? WHERE id =?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	return "updated successful", nil
}

func (database *UserDB) Get() (*entity.User, error) {
	stmt, err := database.db.Prepare("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()

	var user entity.User
	if err := row.Scan(&user.ID, &user.Code, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.DisabledAt, &user.LastLoginAt, &user.Type, &user.Police, &user.FirstName, &user.LastName, &user.EmailHash, &user.NationalIdHash, &user.TelephoneHash, &user.BirthDate, &user.AvatarUri, &user.Language, &user.DarkMode, &user.Permissions, &user.LastPasswordChangeAt, &user.TwoFactorEnabled); err != nil {
		return nil, err
	}

	return &user, nil
}

func (database *UserDB) List() ([]entity.User, error) {
	stmt, err := database.db.Prepare("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Code, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.DisabledAt, &user.LastLoginAt, &user.Type, &user.Police, &user.FirstName, &user.LastName, &user.EmailHash, &user.NationalIdHash, &user.TelephoneHash, &user.BirthDate, &user.AvatarUri, &user.Language, &user.DarkMode, &user.Permissions, &user.LastPasswordChangeAt, &user.TwoFactorEnabled)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (database *UserDB) VerifyUnique(email string) bool {
	stmt, err := database.db.Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		return false
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)

	var user entity.User
	err = row.Scan(
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
	)
	if err != nil {
		return false
	}
	return true
}
