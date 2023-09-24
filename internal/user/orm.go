package user

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type UserORM struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Email     string    `gorm:"column:email;not null" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	Source    string    `gorm:"column:source" json:"source"`
	Status    string    `gorm:"column:status;not null;default:'enabled'" json:"status"`
	Role      int32     `gorm:"column:role;not null;default:2" json:"role"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName User's table name
func (*UserORM) TableName() string {
	return TableNameUser
}
