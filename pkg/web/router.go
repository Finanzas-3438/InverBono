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

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "modernc.org/sqlite"
)

const (
	port   = ":8080"
	secret = "clave-secreta"
)

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

	if err := db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatalf("Error al migrar tabla users: %v", err)
	}

	repo := &repositories.UserRepoDB{DB: db}
	authUseCase := usecases.NewAuthUseCase(repo, secret)
	authHandler := handlers.NewAuthHandler(authUseCase)

	// Rutas públicas
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pkg/web/views/login.html")
	})
	http.HandleFunc("/signup", authHandler.SignupHandler)
	http.HandleFunc("/login", authHandler.LoginWithCookie(secret))
	http.HandleFunc("/logout", authHandler.LogoutHandler)

	// Ruta protegida
	http.HandleFunc("/dashboard", middleware.JWTMiddlewareFromCookie(secret, handlers.DashboardHandler))

	// Iniciar servidor
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error al iniciar servidor HTTP: %v", err)
	}
}
