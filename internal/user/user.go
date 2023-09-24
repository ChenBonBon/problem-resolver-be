package user

import (
	"backend/internal/db"
	"backend/internal/mail"
	"backend/internal/pkg/res"
	"backend/pkg/token"
	"context"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

type User struct {
	Username string `json:"username"`
	NewToken string `json:"token"`
}

type UserClaims struct {
	UserID int32 `json:"userId"`
}

type ForgetClaims struct {
	Email string `json:"email"`
}

func AddUser(email string) (UserORM, error) {
	emailArr := strings.Split(email, "@")
	user, err := addUserByCode(email, emailArr[0])

	return user, err
}

func GetUser(ctx iris.Context) {
	claims := jwt.Get(ctx).(*UserClaims)

	user, err := getUserById(claims.UserID)

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "获取用户信息失败", err)
		return
	}

	token, err := RefreshToken(user.ID)

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "Token生成失败", err)
		return
	}

	res.Success(ctx, User{
		Username: user.Name,
		NewToken: token,
	})
}

func RefreshToken(id int32) (string, error) {
	userClaims := UserClaims{
		UserID: id,
	}

	sigKey := os.Getenv("TOKEN_SIG_KEY")
	expiredTime := os.Getenv("TOKEN_EXPIRED_TIME")

	expiredArr := strings.Split(expiredTime, "*")
	expired := 1

	for _, value := range expiredArr {

		res, err := strconv.Atoi(value)

		if err != nil {
			return "", err
		}

		expired *= res
	}

	token, err := token.Generate(userClaims, []byte(sigKey), time.Duration(expired)*time.Second)

	return token, err
}

func ForgetPassword(ctx iris.Context) {
	username := ctx.URLParam("username")

	address, err := mail.ValidateMail(username)

	if err != nil {
		user, err := getUserByName(username)

		if err != nil {
			if user.ID == 0 {
				res.Problem(ctx, iris.StatusBadRequest, "该用户不存在", err)
				return
			}

			res.Problem(ctx, iris.StatusInternalServerError, "获取用户信息失败", err)
			return
		}

		address = user.Email
	}

	forgetClaims := ForgetClaims{
		Email: address,
	}

	sigKey := os.Getenv("TOKEN_SIG_KEY")
	expiredTime := os.Getenv("TOKEN_EXPIRED_TIME")

	expiredArr := strings.Split(expiredTime, "*")
	expired := 1

	for _, value := range expiredArr {

		num, err := strconv.Atoi(value)

		if err != nil {
			res.Problem(ctx, iris.StatusInternalServerError, "Token生成失败", err)
			return
		}

		expired *= num
	}

	forgetToken, err := token.Generate(forgetClaims, []byte(sigKey), time.Duration(expired)*time.Second)

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "Token生成失败", err)
		return
	}

	content, err := os.ReadFile("mail/templates/forget.html")
	protocol := os.Getenv("PROTOCOL")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	stokens, err := token.GenerateStoken(forgetToken)

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "SToken生成失败", err)
		return
	}

	redisCtx := context.Background()

	redisErr := db.Redis.Set(redisCtx, stokens[0], forgetToken, 0).Err()

	if redisErr != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "Redis写入失败", err)
		return
	}

	mail.SendMail(address, "欢迎登录小镇做题家", strings.Replace(string(content), "{{ url }}", strings.Join([]string{protocol, "://", host, ":", port, "/reset-password?stoken=", stokens[0]}, ""), -1))

	res.Success(ctx, nil)
}

func ResetPassword(ctx iris.Context) {
	var resetPassword PasswordWithStoken

	err := ctx.ReadJSON(&resetPassword)

	if err != nil {
		res.Problem(ctx, iris.StatusBadRequest, "参数错误", err)
		return
	}

	redisCtx := context.Background()
	token, err := db.Redis.Get(redisCtx, resetPassword.Stoken).Result()

	if err != nil {
		res.Problem(ctx, iris.StatusBadRequest, "Stoken不正确", err)
		return
	}

	sigKey := os.Getenv("TOKEN_SIG_KEY")

	verifiedToken, err := jwt.Verify(jwt.HS256, []byte(sigKey), []byte(token))

	if err != nil {
		res.Problem(ctx, iris.StatusBadRequest, "Stoken不正确", err)
		return
	}

	var claims ForgetClaims
	verifiedToken.Claims(&claims)

	email := claims.Email

	err = updateUserPassword(email, resetPassword.Password)

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "更新密码失败", err)
		return
	}

	res.Success(ctx, nil)
}
