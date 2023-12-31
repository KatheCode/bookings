package render

import (
	"KatheCode/bookings/pkg/config"
	"KatheCode/bookings/pkg/models"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a

}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
    return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested templated from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get temaplate from tempkate cache")
	}

	buf := new(bytes.Buffer)

    td = AddDefaultData(td)
    
	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser:", err)
		return
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// get all of the files named .page.tmpl from template folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return cache, err
	}

	// range through all files ending with .page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}
	return cache, nil

}
