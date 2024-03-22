package core

import (
	"html/template"
	"os"
	"strconv"
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
		templates: template.Must(template.ParseGlob("resources/views/*.html")),
	}
}
