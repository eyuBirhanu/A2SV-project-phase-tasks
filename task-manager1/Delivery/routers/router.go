package routers

import (
	"task-manager/Delivery/controllers"
	infrastructure "task-manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(ctrl *controllers.TaskController, jwtService infrastructure.JWTService) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", ctrl.Register)
		api.POST("/login", ctrl.Login)

		protected := api.Group("/tasks")
		protected.Use(infrastructure.AuthMiddleware(jwtService))
		{
			protected.GET("/", ctrl.GetTasks)
			protected.POST("/", ctrl.CreateTask)
		}
	}

	return r
}
