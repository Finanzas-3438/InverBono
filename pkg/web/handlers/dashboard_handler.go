package handlers

import (
	"html/template"
	"net/http"

	"github.com/Finanzas-3438/InverBono.git/pkg/web/middleware"
	"github.com/golang-jwt/jwt/v5"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserContextKey).(jwt.MapClaims)

	data := struct {
		Username string
		Role     string
	}{
		Username: claims["username"].(string),
		Role:     claims["role"].(string),
	}

	tmpl := template.Must(template.ParseFiles("pkg/web/views/dashboard.html"))
	tmpl.Execute(w, data)
}
