package main

import (
	"crud/config"
	"crud/controller"
	_ "crud/docs"
	"crud/helper"
	"crud/model"
	"crud/repository"
	"crud/router"
	"crud/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // This allows all origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	routes.Use(cors.New(config))

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
