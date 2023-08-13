package service

import (
	"context"
	"errors"
	"net/http"
	"taskmaster/internal/models"
	"taskmaster/internal/repository"
	mock_repository "taskmaster/internal/repository/mocks"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateTask(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockTask, task models.InputTask)
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         models.InputTask
		expected     error
		HTTPstatus   int
	}{
		{
			name: "success to create task",
			mockBehavior: func(r *mock_repository.MockTask, task models.InputTask) {
				fixedTime := time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC)
				r.EXPECT().CreateTask(gomock.Any(), task, fixedTime, time.Now()).Return("", 0, nil)
			},
			args: models.InputTask{
				Title:    "test",
				ActiveAt: "2023-10-10",
			},
			expected:   nil,
			HTTPstatus: 0,
		},
		{
			name: "invalid time",
			mockBehavior: func(r *mock_repository.MockTask, task models.InputTask) {
			},
			args: models.InputTask{
				Title:    "test",
				ActiveAt: "2023-100-10",
			},
			expected:   errors.New("invalid time"),
			HTTPstatus: http.StatusBadRequest,
		},
		{
			name:         "time is less than current",
			mockBehavior: func(r *mock_repository.MockTask, task models.InputTask) {},
			args: models.InputTask{
				Title:    "test",
				ActiveAt: "2023-07-10",
			},
			expected:   errors.New("invalid time"),
			HTTPstatus: http.StatusBadRequest,
		},
		{
			name:         "invalid title",
			mockBehavior: func(r *mock_repository.MockTask, task models.InputTask) {},
			args: models.InputTask{
				Title:    "   ",
				ActiveAt: "2023-07-10",
			},
			expected:   errors.New("invalid title"),
			HTTPstatus: http.StatusBadRequest,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockTask(c)
			test.mockBehavior(repo, test.args)

			repos := &repository.Repository{Task: repo}
			service := newTaskService(repos)

			_, status, err := service.CreateTask(context.TODO(), test.args)
			// Проверьте результат
			assert.Equal(t, test.expected, err)
			assert.Equal(t, test.HTTPstatus, status)

		})
	}
}

func TestUpdateTask(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockTask, task models.InputTask, id string)
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		inputTask    models.InputTask
		inputID      string
		expected     error
		HTTPstatus   int
	}{
		{
			name: "success to create task",
			mockBehavior: func(r *mock_repository.MockTask, task models.InputTask, id string) {
				r.EXPECT().UpdateTask(gomock.Any(), task, id, gomock.Any()).Return(http.StatusNoContent, nil)
			},
			inputTask: models.InputTask{
				Title:    "test",
				ActiveAt: "2023-10-10",
			},
			inputID:    primitive.NewObjectID().Hex(),
			expected:   nil,
			HTTPstatus: http.StatusNoContent,
		},
		{
			name: "invalid id",
			mockBehavior: func(r *mock_repository.MockTask, task models.InputTask, id string) {
				fixedTime := time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC)
				r.EXPECT().UpdateTask(gomock.Any(), task, id, fixedTime).Return(http.StatusBadRequest, errors.New("invalid id"))
			},
			inputTask: models.InputTask{
				Title:    "test",
				ActiveAt: "2023-10-10",
			},
			inputID:    "sadasfasd",
			expected:   errors.New("invalid id"),
			HTTPstatus: http.StatusBadRequest,
		},
		{
			name: "invalid title",
			mockBehavior: func(r *mock_repository.MockTask, task models.InputTask, id string) {
			},
			inputTask: models.InputTask{
				Title:    "               ",
				ActiveAt: "2023-10-10",
			},
			inputID:    primitive.NewObjectID().Hex(),
			expected:   errors.New("invalid title"),
			HTTPstatus: http.StatusBadRequest,
		},
		{
			name: "invalid time",
			mockBehavior: func(r *mock_repository.MockTask, task models.InputTask, id string) {
			},
			inputTask: models.InputTask{
				Title:    "test",
				ActiveAt: "2023-100-10",
			},
			inputID:    primitive.NewObjectID().Hex(),
			expected:   errors.New("invalid time"),
			HTTPstatus: http.StatusBadRequest,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockTask(c)
			test.mockBehavior(repo, test.inputTask, test.inputID)

			repos := &repository.Repository{Task: repo}
			service := newTaskService(repos)

			status, err := service.UpdateTask(context.TODO(), test.inputTask, test.inputID)
			// Проверьте результат
			assert.Equal(t, test.expected, err)
			assert.Equal(t, test.HTTPstatus, status)
		})
	}
}
