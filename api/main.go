package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm" // https://github.com/jinzhu/gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors" // https://github.com/rs/cors
)

var db *gorm.DB

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	// Connect DB
	log.Printf("db: %v", cfg.DatabaseURL)
	d, err := gorm.Open("mysql", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	db = d
	defer db.Close()

	db.AutoMigrate(&Memo{})

	// Routing
	r := mux.NewRouter()
	// http://localhost:15000/memo
	r.HandleFunc("/memo", post).Methods("POST")
	// http://localhost:15000/memo
	r.HandleFunc("/memo", get).Methods("GET")
	// http://localhost:15000/memo/id
	r.HandleFunc("/memo/{id:[0-9]+}", put).Methods("PUT")
	// http://localhost:15000/memo/id
	r.HandleFunc("/memo/{id:[0-9]+}", delete).Methods("DELETE")

	// Set cors
	c := cors.New(cors.Options{
		// AllowedOrigins: []string{"http://localhost:8001"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))
}
