package BaseRouter

import (
	"web_app/logic/api"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(group *gin.RouterGroup) {
	BaseRouter := group.Group("base")
	{
		BaseRouter.POST("login", api.Login)
		BaseRouter.POST("captcha", api.Captcha)
		BaseRouter.GET("captcha/:captchaId", api.CaptchaImg)
		BaseRouter.POST("userVoted", api.UserVoted)
	}
}
