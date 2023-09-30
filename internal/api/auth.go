package api

import (
	"blog/global"
	"blog/internal/service"
	"blog/pkg/app"
	"blog/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("auth info valid error: %v", errs)
		response.ToErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	err := service.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("auth info check error: %v", err)
		response.ToErrResponse(errcode.UnauthorizedAuthNotExists)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("generate token error: %v", err)
		response.ToErrResponse(errcode.UnauthorizedTokenGenError)
		return
	}

	response.ToResponse(gin.H{"token": token})
}
