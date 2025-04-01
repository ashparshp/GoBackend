package config

import "html/template"

// AppConfig holds the application config
type AppConfig struct {
	UseCahce bool
	TemplateCache map[string]*template.Template
}