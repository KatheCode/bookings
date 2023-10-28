package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// Holds the application configuration
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
	InProd bool
	Session *scs.SessionManager
}
