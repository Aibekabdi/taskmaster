package http

import (
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

	id, httpStatus, err := h.service.CreateTask(c.Request.Context(), task)
	if err != nil {
		errorResponse(c, httpStatus, err.Error())
		return
	}

	// Возвращаем ID созданной задачи
	c.JSON(http.StatusCreated, gin.H{"task_id": id})
}

func (h *Handler) updateTask(c *gin.Context) {
	id := c.Param("id")
	var (
		task models.InputTask
		err  error
	)
	if err = c.BindJSON(&task); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	httpStatus, err := h.service.UpdateTask(c.Request.Context(), task, id)
	if err != nil {
		errorResponse(c, httpStatus, err.Error())
		return
	}

	c.Status(httpStatus)
}

func (h *Handler) deleteTask(c *gin.Context) {
	id := c.Param("id")
	httpStatus, err := h.service.Task.DeleteTask(c.Request.Context(), id)
	if err != nil {
		errorResponse(c, httpStatus, err.Error())
		return
	}

	c.Status(httpStatus)
}

func (h *Handler) markTaskAsDone(c *gin.Context) {
	id := c.Param("id")
	httpStatus, err := h.service.MarkTaskAsDone(c.Request.Context(), id)
	if err != nil {
		errorResponse(c, httpStatus, err.Error())
		return
	}

	c.Status(httpStatus)
}

func (h *Handler) getTasks(c *gin.Context) {
	status := c.DefaultQuery("status", "active")

	tasks, err := h.service.GetTasks(c.Request.Context(), status)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, tasks)
}
