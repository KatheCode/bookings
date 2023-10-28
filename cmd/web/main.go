package main

import (
	"KatheCode/bookings/pkg/config"
	"KatheCode/bookings/pkg/handlers"
	"KatheCode/bookings/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	// change this to trur when it's on Prod
	app.InProd = false

	app.Session = createSession()

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("canot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	log.Println("Starting server in port:", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func createSession() *scs.SessionManager {
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd

	return session
}
