package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/waldrey/go-expert/apis/configs"
	_ "github.com/waldrey/go-expert/apis/docs"
	"github.com/waldrey/go-expert/apis/internal/entity"
	"github.com/waldrey/go-expert/apis/internal/infra/database"
	"github.com/waldrey/go-expert/apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Go Expert API Example
// @version 1
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name Waldrey S. Silva
// @contact.url https://waldrey.com
// @contact.email eu@waldrey.com

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", config.DBName)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	userDB := database.NewUser(db)
	productHandler := handlers.NewProductHandler(productDB)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.WithValue("jwt", config.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", config.JWTExpiresIn))
	// r.Use(LogRequest)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
