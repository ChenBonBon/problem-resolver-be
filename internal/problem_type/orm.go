package problem_type

import (
	"time"
)

const TableNameProblemType = "problem_types"

// ProblemType mapped from table <problem_types>
type ProblemTypeORM struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Status    string    `gorm:"column:status;not null:default:'enabled'" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	CreatedBy int32     `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy int32     `gorm:"column:updated_by" json:"updated_by"`
}

// TableName ProblemType's table name
func (*ProblemTypeORM) TableName() string {
	return TableNameProblemType
}
