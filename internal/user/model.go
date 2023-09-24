package user

import (
	"backend/internal/db"
	"backend/pkg/encrypt"
	"time"
)

type PasswordWithStoken struct {
	Password string `json:"password"`
	Stoken   string `json:"stoken"`
}

type UserIdAndName struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func GetUserByEmail(email string) (UserIdAndName, error) {
	var user UserIdAndName

	result := db.Postgres.Model(&UserORM{}).Select("id, name").Where("email = ?", email).First(&user)

	return user, result.Error
}

func GetUserByPassword(username string, password string) (UserIdAndName, error) {
	var user UserIdAndName

	result := db.Postgres.Model(&UserORM{}).Select("id, name").Where("name = ? AND password = ?", username, password).First(&user)

	return user, result.Error
}

func addUserByCode(email string, username string) (UserORM, error) {
	user := UserORM{
		Email:  email,
		Name:   username,
		Source: "Code",
	}

	result := db.Postgres.Create(&user)

	return user, result.Error
}

func getUserById(id int32) (UserORM, error) {
	user := UserORM{}

	result := db.Postgres.First(&user, id)

	return user, result.Error
}

func getUserByName(name string) (UserORM, error) {
	user := UserORM{}

	result := db.Postgres.Where("name = ?", name).First(&user)

	return user, result.Error
}

func updateUserPassword(email string, password string) error {
	var user UserORM

	result := db.Postgres.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return result.Error
	}

	user.Password = encrypt.Md5(password)
	user.UpdatedAt = time.Now()

	result = db.Postgres.Save(&user)

	return result.Error
}
