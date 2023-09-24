package problem_submit

import (
	"time"
)

const TableNameProblemSubmit = "problem_submits"

// ProblemSubmit mapped from table <problem_submits>
type ProblemSubmitORM struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProblemID int32     `gorm:"column:problem_id;not null" json:"problem_id"`
	Content   string    `gorm:"column:content;not null" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	CreatedBy int32     `gorm:"column:created_by;not null" json:"created_by"`
}

// TableName ProblemSubmit's table name
func (*ProblemSubmitORM) TableName() string {
	return TableNameProblemSubmit
}
