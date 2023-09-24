package problem_submit

import (
	"backend/internal/pkg/res"
	"backend/internal/user"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func SubmitProblem(ctx iris.Context) {
	claims := jwt.Get(ctx).(*user.UserClaims)
	userId := claims.UserID
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 32)

	if err != nil {
		res.Problem(ctx, iris.StatusBadRequest, "问题ID错误", err)
		return
	}

	submit := ProblemSubmitORM{}

	err = ctx.ReadJSON(&submit)

	if err != nil {
		res.Problem(ctx, iris.StatusBadRequest, "提交参数错误", err)
		return
	}

	err = addProblemSubmit(int32(id), userId, submit.Content)

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "提交问题失败", err)
		return
	}

	res.Success(ctx, nil)
}
