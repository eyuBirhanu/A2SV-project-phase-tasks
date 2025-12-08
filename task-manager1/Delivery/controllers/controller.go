package controllers

import (
	"net/http"
	domain "task-manager/Domain"
	usecases "task-manager/Usecases"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUsecase usecases.TaskUsecase
	userUsecase usecases.UserUsecase
}

func NewTaskController(t usecases.TaskUsecase, u usecases.UserUsecase) *TaskController {
	return &TaskController{taskUsecase: t, userUsecase: u}
}

func (tc *TaskController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tc.userUsecase.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (tc *TaskController) Login(c *gin.Context) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := tc.userUsecase.Login(creds.Email, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tc.taskUsecase.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, _ := tc.taskUsecase.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}
