package api

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

func reqValidator(c *gin.Context, value interface{})  validator.FieldError {
	c.ShouldBind(value)
	validate := validator.New()
	err := validate.Struct(value)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}
	return nil
}

