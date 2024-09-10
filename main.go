package main

import (
	"GoUrlShortener/internal/dao/impl"
	"GoUrlShortener/internal/db"
	"GoUrlShortener/internal/handlers"
	"GoUrlShortener/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	db.InitDBPool()
	defer db.CloseDBPool()
	dbPool := db.GetDBPool()
	shortURLDao := impl.NewShortURLImpl(dbPool)
	shortURLService := services.NewShortURLService(shortURLDao)
	shortURLHandler := handlers.NewShortURLHandler(shortURLService)

	router := gin.Default()

	router.POST("/shorten", shortURLHandler.CreateShortURL)
	router.GET("/:shortCode", shortURLHandler.RedirectToOriginalURL)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Printf("Server started at %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
