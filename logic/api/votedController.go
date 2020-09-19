package api

import (
	"web_app/dao/redis"
	"web_app/model"
	"web_app/response"
	"web_app/utils"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func UserVoted(ctx *gin.Context) {
	//ParamBindError := errors.New("参数绑定错误")
	m := new(model.VotedData)
	err := ctx.ShouldBindJSON(&m)
	if err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			response.FailWithMessage("参数绑定错误", ctx)
			return
		}
		e := utils.RemoveTopStruct(errors.Translate(utils.Trans))
		response.FailWithData(e, ctx)
		return

	}

	conn := redis.RedisConn.Get()
	conn.Do()

	response.OkWithData(m, ctx)
}
