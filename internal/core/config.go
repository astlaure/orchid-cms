package core

import (
	"html/template"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func GetDebugConfig() bool {
	value, err := strconv.ParseBool(os.Getenv("DEBUG"))

	if err != nil {
		return false
	}

	return value
}

func GetRenderer() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("resources/views/**/*.html")),
	}
}

func GetValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}
