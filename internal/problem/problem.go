package problem

import (
	"backend/internal/pkg/res"
	"backend/internal/user"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func GetProblems(ctx iris.Context) {
	status := ctx.URLParam("status")

	problems, err := getProblems(StatusType(status))

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "查询问题失败", err)
		return
	}

	res.Success(ctx, problems)
}

func GetProblem(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 32)

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("问题 ID 错误").Detail(err.Error()).Type("Param Problem"))
		return
	}

	problem, err := getProblemById(int32(id))

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "查询问题失败", err)
		return
	}

	res.Success(ctx, problem)
}

func AddProblem(ctx iris.Context) {
	claims := jwt.Get(ctx).(*user.UserClaims)
	userId := claims.UserID

	var problem ProblemORM

	err := ctx.ReadJSON(&problem)

	if err != nil {
		res.Problem(ctx, iris.StatusBadRequest, "问题参数错误", err)
		return
	}

	err = addProblem(problem.Title, problem.Description, problem.Answer, DifficultyType(problem.Difficulty), problem.Types, userId)

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "创建问题失败", err)
		return
	}

	res.Success(ctx, nil)
}

func GetProblemsByUserId(ctx iris.Context) {
	claims := jwt.Get(ctx).(*user.UserClaims)
	userId := claims.UserID

	problems, err := getProblemsByUserId(userId)

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "查询问题失败", err)
		return
	}

	res.Success(ctx, problems)
}

func UpdateProblem(ctx iris.Context) {
	claims := jwt.Get(ctx).(*user.UserClaims)
	userId := claims.UserID
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 32)

	if err != nil {
		res.Problem(ctx, iris.StatusBadRequest, "问题ID错误", err)
		return
	}

	var problem ProblemORM

	err = ctx.ReadJSON(&problem)

	if err != nil {
		res.Problem(ctx, iris.StatusBadRequest, "问题参数错误", err)
		return
	}

	err = updateProblem(int32(id), problem, userId)

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "更新问题失败", err)
		return
	}

	res.Success(ctx, nil)
}
