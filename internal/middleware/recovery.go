package middleware

import (
	"fmt"
	"time"

	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/global"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/app"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/email"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSettings.Host,
		Port:     global.EmailSettings.Port,
		IsSSL:    global.EmailSettings.IsSSL,
		UserName: global.EmailSettings.UserName,
		Password: global.EmailSettings.Password,
		From:     global.EmailSettings.From,
	})

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err: %v"
				global.Logger.WithCallersFrames().Errorf(s, err)

				err := defailtMailer.SendMail(
					global.EmailSettings.To,
					fmt.Sprintf("例外拋出, 發生時間: %d", time.Now().Unix()),
					fmt.Sprintf("錯誤訊息: %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
