package handlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/Finanzas-3438/InverBono.git/pkg/domain/entities"
	"github.com/Finanzas-3438/InverBono.git/pkg/domain/usecases"
	"github.com/Finanzas-3438/InverBono.git/pkg/web/middleware"
	"github.com/golang-jwt/jwt/v5"
)

type ProfileHandler struct {
	Usecase usecases.ProfileUsecase
}

func NewProfileHandler(uc usecases.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{Usecase: uc}
}

// 🔄 Convierte claims del JWT en una entidad User
func getUserFromContext(r *http.Request) (*entities.User, error) {
	val := r.Context().Value(middleware.UserContextKey)
	if val == nil {
		return nil, errors.New("no se encontró el token en el contexto")
	}

	claims, ok := val.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("los claims no son válidos")
	}

	userIDRaw, ok := claims["user_id"]
	if !ok || userIDRaw == nil {
		return nil, errors.New("claim user_id faltante")
	}

	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		return nil, errors.New("claim user_id no es float64")
	}

	username, _ := claims["username"].(string)
	role, _ := claims["role"].(string)

	return &entities.User{
		ID:       uint(userIDFloat),
		Username: username,
		Role:     role,
	}, nil
}

// GET /me/profile
func (h *ProfileHandler) GetOwnProfile(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromContext(r)
	if err != nil {
		http.Error(w, "Usuario no autenticado: "+err.Error(), http.StatusUnauthorized)
		return
	}

	profile, err := h.Usecase.GetProfileByID(int(user.ID), user)
	if err != nil {
		http.Error(w, "Perfil no encontrado", http.StatusNotFound)
		return
	}

	log.Println("Perfil encontrado:", profile)

	tmpl, err := template.ParseFiles("pkg/web/views/profile.html")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, profile); err != nil {
		http.Error(w, "Error al renderizar vista", http.StatusInternalServerError)
	}
}

/*
func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	var profile entities.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Formato inválido", http.StatusBadRequest)
		return
	}

	err := h.Usecase.UpdateProfile(profile, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}
*/
