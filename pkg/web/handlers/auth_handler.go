package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Finanzas-3438/InverBono.git/pkg/domain/entities"
	"github.com/Finanzas-3438/InverBono.git/pkg/domain/usecases"
)

type AuthHandler struct {
	AuthUseCase *usecases.AuthUseCase
}

func NewAuthHandler(auth *usecases.AuthUseCase) *AuthHandler {
	return &AuthHandler{AuthUseCase: auth}
}

func (h *AuthHandler) LoginWithCookie(secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Solicitud inválida", http.StatusBadRequest)
			return
		}

		token, err := h.AuthUseCase.Login(req.Username, req.Password)
		if err != nil {
			http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			HttpOnly: true,
			Path:     "/",
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func (h *AuthHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "pkg/web/views/signup.html")
		return
	}

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	user := entities.User{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}

	err := h.AuthUseCase.UserRepo.SaveUser(user)
	if err != nil {
		http.Error(w, "Usuario ya existe", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
