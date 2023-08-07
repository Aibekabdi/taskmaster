package http

import (
	"log"
	"net/http"
	"taskmaster/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(c *gin.Context) {
	var (
		task models.InputTask
		err  error
	)
	if err = c.BindJSON(&task); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, httpStatus, err := h.service.CreateTask(task)
	if err != nil {
		errorResponse(c, httpStatus, err.Error())
		return
	}

	// Возвращаем ID созданной задачи
	c.JSON(http.StatusCreated, gin.H{"task_id": id})
}

func (h *Handler) updateTask(c *gin.Context) {

}

func (h *Handler) deleteTask(c *gin.Context) {

}

func (h *Handler) markTaskAsDone(c *gin.Context) {

}

func (h *Handler) getTasks(c *gin.Context) {
	tasks, err := h.service.GetTasks()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(tasks)
	c.JSON(http.StatusOK, tasks)
}
