package web

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Finanzas-3438/InverBono.git/pkg/domain/entities"
	"github.com/Finanzas-3438/InverBono.git/pkg/domain/usecases"
	"github.com/Finanzas-3438/InverBono.git/pkg/repositories"
	"github.com/Finanzas-3438/InverBono.git/pkg/web/handlers"
	"github.com/Finanzas-3438/InverBono.git/pkg/web/middleware"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "modernc.org/sqlite"
)

const (
	port   = ":8080"
	secret = "clave-secreta"
)

type Router struct {
	*mux.Router
	routes []string
}

type Handler func(w http.ResponseWriter, r *http.Request)

func StartServer() {
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		log.Fatalf("Error creando carpeta data/: %v", err)
	}

	sqlDB, err := sql.Open("sqlite", "data/dev.db")
	if err != nil {
		log.Fatalf("Error al abrir base de datos SQLite: %v", err)
	}

	db, err := gorm.Open(sqlite.New(sqlite.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a SQLite con GORM: %v", err)
	}

	// Migraciones
	db.AutoMigrate(&entities.User{}, &entities.Profile{})

	// Repositorios y casos de uso
	userRepo := &repositories.UserRepoDB{DB: db}
	profileRepo := &repositories.ProfileRepoDB{DB: db}

	profileUseCase := usecases.NewProfileUseCase(profileRepo, userRepo)
	authUseCase := usecases.NewAuthUseCase(userRepo, secret, profileRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authUseCase)
	profileHandler := handlers.NewProfileHandler(profileUseCase)

	// Router principal
	router := mux.NewRouter()

	// Archivos estáticos (opcional)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Rutas públicas
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pkg/web/views/auth/login.html")
	}).Methods("GET")

	router.HandleFunc("/signup", authHandler.SignupHandler).Methods("POST")
	router.HandleFunc("/login", authHandler.LoginWithCookie(secret)).Methods("POST")
	router.HandleFunc("/logout", authHandler.LogoutHandler).Methods("POST")

	// Ruta protegida simple
	router.HandleFunc("/dashboard", middleware.JWTMiddlewareFromCookie(secret, handlers.DashboardHandler)).Methods("GET")
	router.HandleFunc("/signup", authHandler.SignupHandler).Methods("GET", "POST")

	// ✅ Nueva vista perfil basada en usuario autenticado
	router.HandleFunc("/me", middleware.JWTMiddlewareFromCookie(secret, profileHandler.GetOwnProfile)).Methods("GET")

	// Servidor
	log.Printf("✅ Servidor escuchando en http://localhost%s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Error al iniciar servidor HTTP: %v", err)
	}

	r := Router{}
	setupStaticRoutes(r)
}

func (r *Router) Static(route string, folder string) {
	r.PathPrefix(route).Handler(http.StripPrefix(route, http.FileServer(http.Dir(folder))))
}

func (r *Router) Get(path string, h Handler) {
	// path = "GET " + path
	// r.AddRoute(path, h)
	r.Router.HandleFunc(path, h).Methods(http.MethodGet)
}

func setupStaticRoutes(r Router) {
	r.Static("/fonts", "public/fonts/")
	r.Get("/js/app.js", Handler(handlers.MakeGzipHandler(handlers.AppJS)))
	r.Get("/js/app.css", Handler(handlers.MakeGzipHandler(handlers.AppCSS)))
	r.Static("/css", "public/css/")
	r.Static("/img", "public/img/")
	r.Static("/js", "public/js/")
	r.Static("/favicon", "public/favicon/")
	r.Static("/user_uploads", "storage/user_uploads/")
}
