package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

var (
	uni *ut.UniversalTranslator
	//validate *validator.Validate
	trans ut.Translator
)

func Init() {
	//注册翻译器
	zhInstance := zh.New()
	uni = ut.New(zhInstance, zhInstance)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	//注册翻译器
	err := zhtranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
}

// Translate 翻译错误信息
func Translate(err error) string {
	var result = make([]string, 0)
	errorArr := err.(validator.ValidationErrors)
	for _, err := range errorArr {
		result = append(result, err.Translate(trans))
	}

	return strings.Join(result, ",")
}
