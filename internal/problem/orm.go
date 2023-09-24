package problem

import (
	"time"

	"github.com/lib/pq"
)

const TableNameProblem = "problems"

// Problem mapped from table <problems>
type ProblemORM struct {
	ID          int32         `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Title       string        `gorm:"column:title;not null" json:"title"`
	Description string        `gorm:"column:description;not null" json:"description"`
	Difficulty  string        `gorm:"column:difficulty;not null" json:"difficulty"`
	Status      string        `gorm:"column:status;not null;default:'enabled'" json:"status"`
	Answer      string        `gorm:"column:answer" json:"answer"`
	Types       pq.Int32Array `gorm:"column:types;type:integer[]" json:"types"`
	CreatedAt   time.Time     `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	CreatedBy   int32         `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedAt   time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy   int32         `gorm:"column:updated_by" json:"updated_by"`
}

// TableName Problem's table name
func (*ProblemORM) TableName() string {
	return TableNameProblem
}
