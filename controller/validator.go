package controller

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

/* Compared to the default validator in Gin,
this Validator lib adds globalization support for validation error messages.*/

// Global translator instance
var trans ut.Translator

// InitTrans initializes the validator translator
func InitTrans(locale string) (err error) {
	// Customize Gin's validator engine so we can plug in our translator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Register custom validation rules here if needed
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // Chinese translator
		enT := en.New() // English translator

		// The first argument is the fallback locale
		// The following arguments are the supported locales (you can pass multiple)
		// uni := ut.New(zhT, zhT) works as well
		uni := ut.New(enT, zhT, enT)

		// The locale usually comes from the HTTP 'Accept-Language' header
		var ok bool
		// We can also call uni.FindTranslator(...) with multiple locales to search
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// Register the translator
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

// removeTopStruct trims struct names (e.g., ParamVoteData.field) from validator error keys.
func removeTopStruct(fields map[string]string) map[string]string {
	if len(fields) == 0 {
		return fields
	}

	res := make(map[string]string, len(fields))
	for field, msg := range fields {
		parts := strings.SplitN(field, ".", 2)
		if len(parts) == 2 {
			res[parts[1]] = msg
			continue
		}
		res[field] = msg
	}
	return res
}
