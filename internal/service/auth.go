package service

import (
	"blog/global"
	"errors"
)

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func CheckAuth(param *AuthRequest) error {
	if param.AppKey == global.JWTSetting.AppKey && param.AppSecret == global.JWTSetting.AppSecret {
		return nil
	}
	return errors.New("auth fail")
}
