package main

import (
	"library_management/controllers"
	"library_management/models"
	"library_management/services"
)

func main() {
	libraryService := services.NewLibrary()

	sampleMember := models.Member{
		ID:   1,
		Name: "Eyu Birhanu",
	}
	libraryService.Members[sampleMember.ID] = sampleMember

	libraryController := controllers.NewLibraryController(libraryService)

	libraryController.Run()
}
