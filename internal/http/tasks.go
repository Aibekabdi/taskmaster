package http

import (
	"net/http"
	"taskmaster/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(c *gin.Context) {
	var (
		task models.Task
		err  error
	)
	if err = c.BindJSON(&task); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
	}
}

func (h *Handler) updateTask(c *gin.Context) {

}

func (h *Handler) deleteTask(c *gin.Context) {

}

func (h *Handler) markTaskAsDone(c *gin.Context) {

}

func (h *Handler) getTasks(c *gin.Context) {

}
