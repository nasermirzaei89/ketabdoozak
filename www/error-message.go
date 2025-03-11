package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"net/http"
)

func sendErrorMessage(w http.ResponseWriter, r *http.Request, msg string, statusCode int) {
	w.Header().Set("HX-Retarget", "body")
	w.Header().Set("HX-Reswap", "beforeend")
	templ.Handler(templates.ErrorMessage(msg), templ.WithStatus(statusCode)).ServeHTTP(w, r)
}
