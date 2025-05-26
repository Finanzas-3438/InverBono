package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Finanzas-3438/InverBono.git/pkg/domain/entities"
	"github.com/Finanzas-3438/InverBono.git/pkg/domain/usecases"
	"github.com/golang-jwt/jwt"
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

		// Buscar usuario
		user, err := h.AuthUseCase.UserRepo.GetByUsername(req.Username)
		if err != nil || user.Password != req.Password {
			http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
			return
		}

		// Crear token con user_id
		claims := jwt.MapClaims{
			"user_id":  user.ID,
			"username": user.Username,
			"role":     user.Role,
			"exp":      time.Now().Add(24 * time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := token.SignedString([]byte(secret))
		if err != nil {
			http.Error(w, "Error al firmar el token", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    signedToken,
			HttpOnly: true,
			Path:     "/",
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func (h *AuthHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10 MB límite de memoria para campos del form

	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "pkg/web/views/auth/signup.html")
		return
	}

	// Necesario para leer los valores del formulario
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al leer formulario", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	role := r.FormValue("role")
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")

	log.Printf("Creando usuario: %s", username)
	if username == "" || password == "" || role == "" {
		http.Error(w, "Campos requeridos", http.StatusBadRequest)

		return
	}
	user := entities.User{
		Username: username,
		Password: password,
		Role:     role,
	}

	if err := h.AuthUseCase.UserRepo.SaveUser(&user); err != nil {
		http.Error(w, "Usuario ya existe", http.StatusConflict)
		return
	}

	log.Printf("Usuario creado: %d", user.ID)
	profile := entities.Profile{
		UserID:    int(user.ID),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	if err := h.AuthUseCase.ProfileRepo.SaveProfile(profile); err != nil {
		http.Error(w, "Error al crear el perfil", http.StatusInternalServerError)
		return
	}

	log.Printf("Usuario y perfil creados para: %s", username)
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
