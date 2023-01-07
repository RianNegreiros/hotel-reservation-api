package handlers

import (
	"go-lodging/pkg/config"
	"go-lodging/pkg/models"
	"go-lodging/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation is the reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

// SingleRooms is the single rooms page handler
func (m *Repository) SingleRooms(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "single-rooms.page.tmpl", &models.TemplateData{})
}

// Suites is the suites page handler
func (m *Repository) Suites(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "suites.page.tmpl", &models.TemplateData{})
}

// Availability is the availability page handler
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability is the post availability page handler
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post availability page"))
}
