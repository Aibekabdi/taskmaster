package http

import (
	"net/http"
	"taskmaster/internal/models"

	"github.com/gin-gonic/gin"
)

// @Summary Create task
// @Description Create a new task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body models.InputTask true "Task info"
// @Success 201
// @Failure 400,500 {object} errorResponseStruct
// @Router /api/todo-list/tasks [post]
func (h *Handler) createTask(c *gin.Context) {
	var (
		task models.InputTask
		err  error
	)
	if err = c.BindJSON(&task); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, httpStatus, err := h.service.CreateTask(c.Request.Context(), task)
	if err != nil {
		ErrorResponse(c, httpStatus, err.Error())
		return
	}

	// Возвращаем ID созданной задачи
	c.JSON(http.StatusCreated, gin.H{"task_id": id})
}

// @Summary Update task
// @Description Update task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Param task body models.InputTask true "models.InputTask info"
// @Success 204
// @Failure 404
// @Failure 400,500 {object} errorResponseStruct
// @Router /api/todo-list/tasks/{id} [put]
func (h *Handler) updateTask(c *gin.Context) {
	id := c.Param("id")
	var (
		task models.InputTask
		err  error
	)
	if err = c.BindJSON(&task); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	httpStatus, err := h.service.UpdateTask(c.Request.Context(), task, id)
	if err != nil {
		ErrorResponse(c, httpStatus, err.Error())
		return
	}

	c.Status(httpStatus)
}

// @Summary Delete task
// @Description Delete task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Success 204
// @Failure 404
// @Failure 400,500 {object} errorResponseStruct
// @Router /api/todo-list/tasks/{id} [delete]
func (h *Handler) deleteTask(c *gin.Context) {
	id := c.Param("id")
	httpStatus, err := h.service.Task.DeleteTask(c.Request.Context(), id)
	if err != nil {
		ErrorResponse(c, httpStatus, err.Error())
		return
	}

	c.Status(httpStatus)
}

// @Summary Mark task as done
// @Description Mark a task as done by its ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Success 204
// @Failure 400,500 {object} errorResponseStruct
// @Failure 404
// @Router /api/todo-list/tasks/{id}/done [put]
func (h *Handler) markTaskAsDone(c *gin.Context) {
	id := c.Param("id")
	httpStatus, err := h.service.MarkTaskAsDone(c.Request.Context(), id)
	if err != nil {
		ErrorResponse(c, httpStatus, err.Error())
		return
	}

	c.Status(httpStatus)
}

// @Summary Get tasks
// @Description Get a list of tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param status query string false "Task status" Enums(active, done)
// @Success 200 {array} models.InputTask
// @Failure 400,500 {object} errorResponseStruct
// @Router /api/todo-list/tasks [get]
func (h *Handler) getTasks(c *gin.Context) {
	status := c.DefaultQuery("status", "active")

	tasks, err := h.service.GetTasks(c.Request.Context(), status)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, tasks)
}
