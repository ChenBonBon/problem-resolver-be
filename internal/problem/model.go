package problem

import (
	"backend/internal/db"
	"time"

	"github.com/lib/pq"
)

type StatusType string

const (
	Enabled  StatusType = "enabled"
	Disabled StatusType = "disabled"
)

type DifficultyType string

const (
	Easy   DifficultyType = "easy"
	Medium DifficultyType = "medium"
	Hard   DifficultyType = "hard"
)

type SolveStatusType string

const (
	Unsolved   SolveStatusType = "unsolved"
	Processing SolveStatusType = "processing"
	Solved     SolveStatusType = "solved"
)

type ProblemExample struct {
	Id          int    `json:"id"`
	ProblemID   int    `json:"problemId"`
	Input       string `json:"input"`
	Output      string `json:"output"`
	Explanation string `json:"explanation"`
}

type ProblemListItem struct {
	Id     int             `json:"id"`
	Title  string          `json:"title"`
	Types  pq.StringArray  `gorm:"column:types;type:character varying(255)[]" json:"types"`
	Status SolveStatusType `json:"status"`
	// Answers    int             `json:"answers"`
	// PassRate   int             `json:"passRate"`
	Difficulty DifficultyType `json:"difficulty"`
}

type ProblemItem struct {
	Id              int              `json:"id"`
	Title           string           `json:"title"`
	Description     string           `json:"description"`
	Difficulty      DifficultyType   `json:"difficulty"`
	Types           pq.StringArray   `gorm:"column:types;type:character varying(255)[]" json:"types"`
	ProblemExamples []ProblemExample `gorm:"foreignKey:ProblemID" json:"examples"`
	Answer          string           `json:"answer"`
}

type UserProblemListItem struct {
	Id         int            `json:"id"`
	Title      string         `json:"title"`
	Types      pq.StringArray `gorm:"column:types;type:character varying(255)[]" json:"types"`
	Status     StatusType     `json:"status"`
	Difficulty DifficultyType `json:"difficulty"`
	CreatedAt  string         `json:"createdAt"`
}

type ProblemTypeListItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type ProblemType struct {
	ProblemListItem
	Status    StatusType `json:"status"`
	CreatedAt string     `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt string     `json:"updatedAt"`
	UpdatedBy string     `json:"updatedBy"`
}

type ProblemSubmit struct {
	Submit string `json:"submit"`
}

func addProblem(title string, description string, answer string, difficulty DifficultyType, types pq.Int32Array, createdBy int32) error {
	problem := ProblemORM{
		Title:       title,
		Description: description,
		Difficulty:  string(difficulty),
		Types:       types,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
		Answer:      answer,
	}

	result := db.Postgres.Debug().Create(&problem)

	return result.Error
}

func getProblems(status StatusType) ([]ProblemListItem, error) {
	var problems []ProblemListItem

	subQuery := db.Postgres.Table("problems as p, problem_types as pt").Select("ARRAY [pt.name] as types").Where("pt.id = ANY (p.types)")

	if status == "" {
		result := db.Postgres.Model(&ProblemORM{}).Find(&problems)

		return problems, result.Error
	}

	result := db.Postgres.Table("(?) as sub ,problems as p", subQuery).Select("p.id", "p.title", "sub.types", "COALESCE(up.status, 'unsolved') as status", "difficulty").Joins("left join user_problems as up on up.problem_id = p.id").Where("p.status = ?", status).Find(&problems)

	return problems, result.Error
}

func getProblemsByUserId(userId int32) ([]UserProblemListItem, error) {
	var problems []UserProblemListItem

	subQuery := db.Postgres.Table("problems as p, problem_types as pt").Select("ARRAY [pt.name] as types").Where("pt.id = ANY (p.types)")

	result := db.Postgres.Table("(?) as sub ,problems as p", subQuery).Select("p.id", "p.title", "sub.types", "difficulty", "p.status", "p.created_at").Where("p.created_by = ?", userId).Find(&problems)

	return problems, result.Error
}

func updateProblem(id int32, items ProblemORM, userId int32) error {
	problem := ProblemORM{}
	nameSet := []string{
		"updated_by",
		"updated_at",
	}

	if items.Title != "" {
		nameSet = append(nameSet, "title")
	}

	if items.Description != "" {
		nameSet = append(nameSet, "description")
	}

	if items.Difficulty != "" {
		nameSet = append(nameSet, "difficulty")
	}

	if items.Status != "" {
		nameSet = append(nameSet, "status")
	}

	if items.Types != nil {
		nameSet = append(nameSet, "types")
	}

	if items.Answer != "" {
		nameSet = append(nameSet, "answer")
	}

	items.UpdatedBy = userId
	items.UpdatedAt = time.Now()

	result := db.Postgres.Model(&problem).Select(nameSet).Where("id = ?", id).Updates(items)

	return result.Error
}

func getProblemById(id int32) (ProblemItem, error) {
	var problem ProblemItem
	examples := []ProblemExample{}

	subQuery := db.Postgres.Table("problems as p, problem_types as pt").Select("ARRAY [pt.name] as types").Where("pt.id = ANY (p.types)")

	result := db.Postgres.Table("(?) as sub, problems as p", subQuery).Select("p.id", "p.title", "p.description", "p.answer", "sub.types", "difficulty").Where("p.id = ?", id).Find(&problem)

	if result.Error != nil {
		return problem, result.Error
	}

	result = db.Postgres.Table("problem_examples").Select("id", "input", "output", "explanation").Where("problem_id = ?", id).Find(&examples)

	if result.Error != nil {
		return problem, result.Error
	}

	problem.ProblemExamples = examples

	return problem, result.Error
}
