package problem_type

import (
	"backend/internal/db"
)

type StatusType string

const (
	Enabled  StatusType = "enabled"
	Disabled StatusType = "disabled"
)

func getProblemTypes(status StatusType) ([]ProblemTypeORM, error) {
	var problemTypes []ProblemTypeORM

	if status == "" {
		result := db.Postgres.Select("id", "name").Find(&problemTypes)

		return problemTypes, result.Error
	}
	result := db.Postgres.Select("id", "name").Where("status = ?", status).Find(&problemTypes)

	return problemTypes, result.Error
}
