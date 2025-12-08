package main

import (
	"task-manager/Delivery/controllers"
	"task-manager/Delivery/routers"
	infrastructure "task-manager/Infrastructure"
	repositories "task-manager/Repositories"
	usecases "task-manager/Usecases"
)

func main() {
	jwtService := infrastructure.NewJWTService()

	userRepo := repositories.NewUserRepository()
	taskRepo := repositories.NewTaskRepository()

	userUsecase := usecases.NewUserUsecase(userRepo, jwtService)
	taskUsecase := usecases.NewTaskUsecase(taskRepo)

	taskController := controllers.NewTaskController(taskUsecase, userUsecase)

	r := routers.SetupRouter(taskController, jwtService)

	r.Run(":8080")
}
