package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var Trans ut.Translator

func init() {
	zh := zh.New()
	uni := ut.New(zh, zh)
	Trans, _ = uni.GetTranslator("zh")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zh_translations.RegisterDefaultTranslations(v, Trans)
	}
}
