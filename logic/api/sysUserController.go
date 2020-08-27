package api

import (
	"time"
	"web_app/middleware"
	"web_app/model"
	"web_app/request"
	"web_app/response"
	"web_app/service"
	"web_app/settings"

	"github.com/dchest/captcha"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func QueryListUser(c *gin.Context) {

	var pageInfo model.PageInfo
	c.ShouldBindJSON(&pageInfo)

	err, userList, total := service.QueryListUser(pageInfo)
	if err != nil {
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(response.ResponsePage{
		List:     userList,
		Total:    total,
		PageSize: pageInfo.PageSize,
		Page:     pageInfo.Page,
	}, c)

}

func Login(c *gin.Context) {

	var user request.RegisterAndLoginStruct
	c.ShouldBindJSON(&user)
	if user.Captcha == "" || user.CaptchaId == "" {
		response.FailWithMessage("请输入验证码", c)
		return
	}
	if captcha.VerifyString(user.CaptchaId, user.Captcha) {
		if err, u := service.Login(&user); err != nil {
			response.FailWithMessage("账号或密码错误", c)
			return
		} else {
			tokenNext(c, *u)
		}
	} else {
		response.FailWithMessage("验证码错误", c)

	}

}

//登录以后签发jwt
func tokenNext(c *gin.Context, user model.SysUser) {
	j := &middleware.SingleKeyConf{
		[]byte(settings.Conf.SingleKey), // 唯一签名
	}
	clams := model.MyClaims{
		Id:        user.Id,
		LoginName: user.Loginname,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
			Issuer:    settings.Conf.SingleKey,               //签名的发行者
		},
	}
	token, err := j.CreateJWT(clams)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
	} else {
		response.OkWithData(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
	}
}

func LoadUserMaxOrderNum(ctx *gin.Context) {
	user := new(model.SysUser)
	ctx.ShouldBindJSON(&user)
	map1 := service.LoadUserMaxOrderNum(user)
	response.OkWithData(map1, ctx)
}
