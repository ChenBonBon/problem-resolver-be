package user_code

import (
	"backend/internal/mail"
	"backend/internal/pkg/res"
	"crypto/rand"
	"math/big"
	"os"
	"strings"

	"github.com/kataras/iris/v12"
)

func SendCode(ctx iris.Context) {
	email := ctx.URLParam("email")

	address, err := mail.ValidateMail(email)

	if err != nil {
		res.Problem(ctx, iris.StatusBadRequest, "邮箱格式错误", err)
		return
	}

	code := []string{}

	for i := 0; i < 6; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(9))
		code = append(code, n.String())
	}

	err = addCode(address, strings.Join(code, ""))

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "生成验证码失败", err)
		return
	}

	content, err := os.ReadFile("mail/templates/code.html")

	if err != nil {
		res.Problem(ctx, iris.StatusInternalServerError, "读取模板文件失败", err)
		return
	}

	mail.SendMail(address, "欢迎登录小镇做题家", strings.Replace(string(content), "{{ code }}", strings.Join(code, ""), -1))

	res.Success(ctx, nil)
}
