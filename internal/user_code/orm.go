package user_code

import (
	"time"
)

const TableNameUserCode = "user_codes"

// UserCode mapped from table <user_codes>
type UserCodeORM struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Email     string    `gorm:"column:email;not null" json:"email"`
	Code      string    `gorm:"column:code;not null" json:"code"`
	Status    string    `gorm:"column:status;not null;default:'unused'" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	ExpiredAt time.Time `gorm:"column:expired_at;not null" json:"expired_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName UserCode's table name
func (*UserCodeORM) TableName() string {
	return TableNameUserCode
}
