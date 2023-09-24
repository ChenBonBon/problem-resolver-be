package main

import (
	"backend/internal/db"
	"backend/internal/login"
	"backend/internal/problem"
	"backend/internal/problem_submit"
	"backend/internal/problem_type"
	"backend/internal/user"
	"backend/internal/user_code"
	"backend/pkg/logger"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func main() {
	ac := logger.MakeAccessLog()
	defer ac.Close()

	app := iris.New()

	app.UseRouter(ac.Handler)

	db.ConnectPostgres()
	db.ConnectRedis()

	sigKey := os.Getenv("TOKEN_SIG_KEY")

	verifier := jwt.NewVerifier(jwt.HS256, sigKey)
	verifier.WithDefaultBlocklist()
	verifyMiddleware := verifier.Verify(func() interface{} {
		return new(user.UserClaims)
	})

	app.OnErrorCode(iris.StatusUnauthorized, func(ctx iris.Context) {
		err := ctx.GetErr()

		ctx.StopWithProblem(iris.StatusUnauthorized, iris.NewProblem().Title("Token验证失败").Detail(err.Error()).Type("Unauthorized Problem"))
	})

	app.Get("/user", user.GetUser).Use(verifyMiddleware)
	app.Delete("/logout", login.Logout).Use(verifyMiddleware)

	app.Get("/code", user_code.SendCode)

	loginParty := app.Party("/login")
	{
		loginParty.Post("/code", login.LoginWithCode)
		loginParty.Post("/password", login.LoginWithPassword)
	}

	problemParty := app.Party("/problems")
	{
		problemParty.Get("", problem.GetProblems)
		problemParty.Get("/{id}", problem.GetProblem)
		problemParty.Get("/types", problem_type.GetProblemTypes)
		problemParty.Post("/{id}", problem_submit.SubmitProblem).Use(verifyMiddleware)
	}

	userParty := app.Party("/users")
	userParty.Use(verifyMiddleware)
	{
		userParty.Get("/forget", user.ForgetPassword)
		userParty.Put("/reset", user.ResetPassword)
		userParty.Get("/problems", problem.GetProblemsByUserId)
		userParty.Post("/problems", problem.AddProblem)
		userParty.Put("/problems/:id", problem.UpdateProblem)
	}

	app.Listen(":8080")
}
