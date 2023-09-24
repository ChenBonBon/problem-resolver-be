package res

import "github.com/kataras/iris/v12"

type SuccessBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(ctx iris.Context, data interface{}) {
	ctx.JSON(SuccessBody{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

func Problem(ctx iris.Context, code int, msg string, err error) {
	ctx.StopWithProblem(code, iris.NewProblem().Title(msg).Detail(err.Error()))
}
