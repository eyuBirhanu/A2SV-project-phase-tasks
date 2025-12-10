package main

import (
	"library_management/concurrency"
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

	reservationHandler := concurrency.NewReservationHandler(libraryService, 5)
	reservationHandler.StartWorkers(3)

	libraryController := controllers.NewLibraryController(libraryService, reservationHandler)

	libraryController.Run()
}
