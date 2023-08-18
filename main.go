package main

import (
	"example/web-service-gin/config"
	"example/web-service-gin/controller"
	"example/web-service-gin/helper"
	"example/web-service-gin/model"
	"example/web-service-gin/repository"
	"example/web-service-gin/router"
	"example/web-service-gin/service"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	log.Info().Msg("Started Server!")

	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	errDb := db.Table("tags").AutoMigrate(&model.Tags{})
	helper.ErrorPanic(errDb)

	//Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	//Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	//Controller
	tagsController := controller.NewTagsController(tagsService)

	//Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	errServer := server.ListenAndServe()
	helper.ErrorPanic(errServer)
}
