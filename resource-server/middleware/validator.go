package middleware

/*
参数检验中间件
*/

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// 定义 Validator
var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("regexp", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(fl.Param())
		return re.MatchString(fl.Field().String())
	})
}

// ValidateStruct 验证结构体中的字段是否符合要求
func ValidateStruct(v interface{}) error {
	// 使用 validator 验证 v
	err := validate.Struct(v)
	if err != nil {
		// 如果验证出错，将验证错误的信息进行格式化处理
		validationErrors := err.(validator.ValidationErrors)
		var errMessages []string
		// 定义标签到中文描述的映射
		tagMessages := map[string]string{
			"required": "是必填的",
			"email":    "不是有效的邮箱地址",
		}
		for _, fieldError := range validationErrors {
			// 获取对应的中文描述，如果没有则使用默认的英文描述
			message, exists := tagMessages[fieldError.Tag()]
			if !exists {
				message = "格式有误"
			}
			errMessages = append(errMessages, fmt.Sprintf("字段 '%s' 验证失败，原因: %s", fieldError.Field(), message))
		}

		return fmt.Errorf("参数验证错误: %s", errMessages)
	}
	return nil
}
