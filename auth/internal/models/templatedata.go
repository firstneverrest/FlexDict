package models

type TemplateData struct {
	Data      map[string]interface{}
	CSRFToken string
	Error     string
}
