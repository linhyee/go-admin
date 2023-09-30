package middleware

import (
	"blog/global"
	"blog/pkg/app"
	"blog/pkg/email"
	"blog/pkg/errcode"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				s := "panic recover error: %v"
				global.Logger.WithCallersFrames().Errorf(s, r)
				err := defaultMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出, 发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息, %v", r),
				)
				if err != nil {
					global.Logger.Panicf("send mail error: %v", err)
				}
				app.NewResponse(c).ToErrResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
