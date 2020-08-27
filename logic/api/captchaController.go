package api

import (
	"web_app/response"
	"web_app/utils"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func Captcha(ctx *gin.Context) {

	captchaId := captcha.NewLen(4)
	ImageUrl := "/captcha/" + captchaId + ".png"
	response.OkWithData(gin.H{
		"CaptchaId": captchaId,
		"ImageUrl":  ImageUrl,
	}, ctx)

}

func CaptchaImg(ctx *gin.Context) {
	utils.ServeHTTP(ctx.Writer, ctx.Request)
}
