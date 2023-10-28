package handlers

import (
	"KatheCode/bookings/pkg/config"
	"KatheCode/bookings/pkg/models"
	"KatheCode/bookings/pkg/render"
	"net/http"
)

// Repository used by handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

    intMap := map[string]int{"number": 7}

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{IntMap: intMap})

}

func (m *Repository) AboutMe(w http.ResponseWriter, r *http.Request) {
	// perform logic
	stringMap := make(map[string]string)
	stringMap["name"] = "Yoshi"

	intMap := map[string]int{"phone": 7}

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap, IntMap: intMap})
}
