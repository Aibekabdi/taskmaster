package http

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"taskmaster/internal/models"
	"taskmaster/internal/service"
	mock_service "taskmaster/internal/service/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestHTTP_CreateTask(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTask, task models.InputTask, mockID primitive.ObjectID)

	tests := []struct {
		name                 string
		inputBody            string
		inputTask            models.InputTask
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Created status",
			inputBody: `{ "title":"test", "activeAt":"2023-10-10"}`,
			inputTask: models.InputTask{
				Title:    "test",
				ActiveAt: "2023-10-10",
			},
			mockBehavior: func(s *mock_service.MockTask, task models.InputTask, mockID primitive.ObjectID) {
				s.EXPECT().CreateTask(gomock.Any(), task).Return(mockID.Hex(), 0, nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:                 "No title",
			inputBody:            `{"activeAt":"2023-10-10"}`,
			mockBehavior:         func(s *mock_service.MockTask, task models.InputTask, mockID primitive.ObjectID) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"status":400,"msg":"invalid input body"}`,
		},
		{
			name:      "Server error",
			inputBody: `{"title": "  " , "activeAt":"2023-10-10"}`,
			inputTask: models.InputTask{
				Title:    "  ",
				ActiveAt: "2023-10-10",
			},
			mockBehavior: func(s *mock_service.MockTask, task models.InputTask, mockID primitive.ObjectID) {
				s.EXPECT().CreateTask(gomock.Any(), task).Return("", 500, errors.New("something went wrong"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"msg":"something went wrong"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockTask(c)
			mockID := primitive.NewObjectID()
			if test.expectedResponseBody == "" {
				test.expectedResponseBody = fmt.Sprintf(`{"task_id":"%s"}`, mockID.Hex())
			}
			test.mockBehavior(repo, test.inputTask, mockID)

			services := &service.Service{Task: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/tasks", handler.createTask)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/tasks", bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)
			t.Log(w.Body.String())
			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHTTP_updateTask(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTask, task models.InputTask, mockID primitive.ObjectID)
	tests := []struct {
		name                 string
		inputBody            string
		inputTask            models.InputTask
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Updated status",
			inputBody: `{ "title":"test", "activeAt":"2023-10-10"}`,
			inputTask: models.InputTask{
				Title:    "test",
				ActiveAt: "2023-10-10",
			},
			mockBehavior: func(s *mock_service.MockTask, task models.InputTask, mockID primitive.ObjectID) {
				s.EXPECT().UpdateTask(gomock.Any(), task, mockID.Hex()).Return(http.StatusNoContent, nil)
			},
			expectedStatusCode:   http.StatusNoContent,
			expectedResponseBody: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockTask(c)
			mockID := primitive.NewObjectID()
			test.mockBehavior(repo, test.inputTask, mockID)

			services := &service.Service{Task: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.PUT("/tasks/:id", handler.updateTask)

			// Create Request
			w := httptest.NewRecorder()
			url := fmt.Sprintf("/tasks/%s", mockID.Hex())
			req := httptest.NewRequest("PUT", url, bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)
			t.Log(w.Body.String())
			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}

}
