package login

import (
	"backend/internal/pkg/res"
	"backend/internal/user"
	"backend/internal/user_code"
	"backend/pkg/encrypt"

	"time"

	"github.com/kataras/iris/v12"
)

func LoginWithPassword(ctx iris.Context) {
	var login UsernameWithPassword

	err := ctx.ReadJSON(&login)

	if err != nil {
		res.Problem(ctx, iris.StatusBadRequest, "用户名或密码错误", err)
		return
	}

	userInfo, err := user.GetUserByPassword(login.Username, encrypt.Md5(login.Password))

	if err != nil {
		if userInfo.ID == 0 {
			res.Problem(ctx, iris.StatusBadRequest, "用户名或密码错误", err)
			return
		} else {
			res.Problem(ctx, iris.StatusInternalServerError, "获取用户信息失败", err)
			return
		}
	}

	token, err := user.RefreshToken(userInfo.ID)

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "Token生成失败", err)
		return
	}

	res.Success(ctx, UsernameWithToken{
		Username: userInfo.Name,
		Token:    token,
	})
}

func LoginWithCode(ctx iris.Context) {
	var login EmailWithCode

	err := ctx.ReadJSON(&login)

	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	userCode, err := user_code.GetCode(login.Email, login.Code)

	if err != nil {
		res.Problem(ctx, iris.StatusBadRequest, "邮箱或验证码错误", err)
		return
	}

	if userCode.Used {
		res.Problem(ctx, iris.StatusBadRequest, "验证码已被使用", err)
		return
	}

	if userCode.ExpiredAt.Before(time.Now()) {
		res.Problem(ctx, iris.StatusBadRequest, "邮箱或验证码已过期", err)
		return
	}

	err = user_code.UpdateCode(login.Email, login.Code)

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "更新验证码状态异常", err)
		return
	}

	userInfo, err := user.GetUserByEmail(login.Email)

	if err != nil {
		if userInfo.ID == 0 {
			newUser, err := user.AddUser(login.Email)

			if err != nil {
				res.Problem(ctx, iris.StatusInternalServerError, "创建用户失败", err)
				return
			}

			userInfo.ID = newUser.ID
			userInfo.Name = newUser.Name
		} else {
			res.Problem(ctx, iris.StatusInternalServerError, "获取用户信息失败", err)
			return
		}
	}

	token, err := user.RefreshToken(userInfo.ID)

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "Token生成失败", err)
		return
	}

	res.Success(ctx, UsernameWithToken{
		Username: userInfo.Name,
		Token:    token,
	})
}

func Logout(ctx iris.Context) {
	err := ctx.Logout()
	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "登出失败", err)
		return
	}

	res.Success(ctx, nil)
}
