package user_code

import (
	"backend/internal/db"
	"time"
)

type CodeUsedAndExpiredAt struct {
	Used      bool      `gorm:"column:used" json:"used"`
	ExpiredAt time.Time `gorm:"column:expired_at" json:"expired_at"`
}

func addCode(email string, code string) error {
	userCode := UserCodeORM{
		Email: email,
		Code:  code,
	}

	result := db.Postgres.Create(&userCode)

	return result.Error
}

func GetCode(email string, code string) (CodeUsedAndExpiredAt, error) {
	var userCode CodeUsedAndExpiredAt

	result := db.Postgres.Model(&UserCodeORM{}).Select("used", "expired_at").Where("email = ? AND code = ?", email, code).First(&userCode)

	return userCode, result.Error
}

func UpdateCode(email string, code string) error {
	var userCode UserCodeORM

	result := db.Postgres.Where("email = ? AND code = ?", email, code).First(&userCode)

	if result.Error != nil {
		return result.Error
	}

	userCode.Status = "used"
	userCode.UpdatedAt = time.Now()

	result = db.Postgres.Save(&userCode)

	return result.Error
}
