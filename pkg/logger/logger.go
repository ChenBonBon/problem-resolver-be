package logger

import (
	"backend/pkg/fs"
	"log/slog"
	"os"
	"time"

	"github.com/kataras/iris/v12/middleware/accesslog"
)

const logDirPath = "logs"

func MakeAccessLog() *accesslog.AccessLog {
	if !fs.Exist(logDirPath) {
		err := fs.CreateDir(logDirPath)

		if err != nil {
			slog.Error("创建日志目录失败", err.Error())
			panic(err)
		}
	}
	now := time.Now()
	filename := now.Format("2006-01-02")
	ac := accesslog.File(logDirPath + "/" + filename + ".log")
	ac.AddOutput(os.Stdout)

	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.Async = false
	ac.IP = true
	ac.BytesReceivedBody = true
	ac.BytesSentBody = true
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = true
	ac.RequestBody = true
	ac.ResponseBody = false
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler

	ac.SetFormatter(&accesslog.JSON{
		Indent:    "  ",
		HumanTime: true,
	})

	return ac
}
