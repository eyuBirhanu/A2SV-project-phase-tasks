package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/tasks", controllers.GetTasks)
		protected.GET("/tasks/:id", controllers.GetTask)

		admin := protected.Group("/")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.POST("/tasks", controllers.CreateTask)
			admin.PUT("/tasks/:id", controllers.UpdateTask)
			admin.DELETE("/tasks/:id", controllers.DeleteTask)
			admin.POST("/promote", controllers.PromoteUser)
		}
	}

	return r
}