package middleware

import (

"github.com/gin-gonic/gin"
"github.com/gin-gonic/gin/binding"
"github.com/go-playground/locales/en"
"github.com/go-playground/locales/zh"
"github.com/go-playground/locales/zh_Hant_TW"
"github.com/go-playground/universal-translator"
vd "github.com/go-playground/validator/v10"
entr "github.com/go-playground/validator/v10/translations/en"
zhtr "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*vd.Validate)
		if ok {
			switch locale{
			case "zh":
				zhtr.RegisterDefaultTranslations(v, trans)
			case "en":
				entr.RegisterDefaultTranslations(v,trans)
			default:
				zhtr.RegisterDefaultTranslations(v, trans)
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}
