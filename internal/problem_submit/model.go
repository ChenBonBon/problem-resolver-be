package problem_submit

import "backend/internal/db"

func addProblemSubmit(problemId int32, userId int32, content string) error {
	problemSubmit := ProblemSubmitORM{
		ProblemID: problemId,
		Content:   content,
		CreatedBy: userId,
	}

	result := db.Postgres.Create(&problemSubmit)

	return result.Error
}
