package http

import (
	"taskmaster/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "taskmaster/docs"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	api := router.Group("/api")
	{
		todolist := api.Group("/todo-list")
		{
			tasks := todolist.Group("/tasks")
			{
				tasks.POST("/", h.createTask)
				tasks.PUT("/:id", h.updateTask)
				tasks.DELETE("/:id", h.deleteTask)
				tasks.PUT("/:id/done", h.markTaskAsDone)
				tasks.GET("/", h.getTasks)
			}
		}
	}
	return router
}
