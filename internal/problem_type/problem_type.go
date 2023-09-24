package problem_type

import (
	"backend/internal/pkg/res"

	"github.com/kataras/iris/v12"
)

func GetProblemTypes(ctx iris.Context) {
	status := ctx.URLParam("status")

	problemTypes, err := getProblemTypes(StatusType(status))

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "获取问题分类失败", err)
		return
	}

	res.Success(ctx, problemTypes)
}
