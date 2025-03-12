package unit_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	v1 "github.com/amha-mersha/icog_clean_code/internal/delivery/http/v1"
	"github.com/amha-mersha/icog_clean_code/internal/domain"
	"github.com/amha-mersha/icog_clean_code/internal/domain/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TaskHandlerTestSuite struct {
	suite.Suite
	router  *gin.Engine
	usecase *MockTaskUsecase
	handler v1.TaskHandler
}

type MockTaskUsecase struct {
	mock.Mock
}

func (m *MockTaskUsecase) CreateTask(task *dto.TaskCreateDTO) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUsecase) GetTask(id uuid.UUID) (*domain.TaskItem, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.TaskItem), args.Error(1)
}

func (m *MockTaskUsecase) ListTasks() ([]domain.TaskItem, error) {
	args := m.Called()
	return args.Get(0).([]domain.TaskItem), args.Error(1)
}

func (m *MockTaskUsecase) UpdateTask(task *dto.TaskUpdateDTO) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUsecase) DeleteTask(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskUsecase) GetTaskByStatus(status string) ([]domain.TaskItem, error) {
	args := m.Called(status)
	return args.Get(0).([]domain.TaskItem), args.Error(1)
}

func (suite *TaskHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.usecase = new(MockTaskUsecase)
	suite.handler = v1.NewTaskHandler(suite.usecase)
	suite.router = gin.Default()
	suite.router.POST("/api/v1/tasks", suite.handler.UploadTaskItem)
	suite.router.GET("/api/v1/tasks/:id", suite.handler.GetTaskByID)
	suite.router.GET("/api/v1/tasks", suite.handler.GetAllTasks)
	suite.router.PUT("/api/v1/tasks", suite.handler.UpdateTask)
	suite.router.DELETE("/api/v1/tasks/:id", suite.handler.DeleteTask)
	suite.router.GET("/api/v1/tasks?status", suite.handler.GetTasksByStatus)
}

func (suite *TaskHandlerTestSuite) TestUploadTaskItem_Success() {
	taskCreateDTO := dto.TaskCreateDTO{
		Title:       "Test Task",
		Description: "Test Description",
		Deadline:    time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	suite.usecase.On("CreateTask", mock.Anything).Return(nil)

	w := performRequest(suite.router, "POST", "/api/v1/tasks", taskCreateDTO)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "message")
}

func (suite *TaskHandlerTestSuite) TestUploadTaskItem_BindError() {
	w := performRequest(suite.router, "POST", "/api/v1/tasks", `{"title": 12345}`)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "error")
}

func (suite *TaskHandlerTestSuite) TestGetTaskByID_Success() {
	taskID := uuid.New()
	task := &domain.TaskItem{
		ID:          taskID,
		Title:       "Test Task",
		Description: "Test Description",
	}

	suite.usecase.On("GetTask", taskID).Return(task, nil)

	w := performRequest(suite.router, "GET", "/api/v1/tasks/"+taskID.String(), nil)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "message")
}

func (suite *TaskHandlerTestSuite) TestGetTaskByID_InvalidID() {
	w := performRequest(suite.router, "GET", "/api/v1/tasks/invalid-id", nil)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid task ID format")
}

func (suite *TaskHandlerTestSuite) TestGetTaskByID_Error() {
	taskID := uuid.New()
	suite.usecase.On("GetTask", taskID).Return(&domain.TaskItem{}, errors.New("some error"))

	w := performRequest(suite.router, "GET", "/api/v1/tasks/"+taskID.String(), nil)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "some error")
}

func (suite *TaskHandlerTestSuite) TestGetAllTasks_Success() {
	tasks := []domain.TaskItem{
		{ID: uuid.New(), Title: "Task 1", Description: "Description 1"},
		{ID: uuid.New(), Title: "Task 2", Description: "Description 2"},
	}
	suite.usecase.On("ListTasks").Return(tasks, nil)

	w := performRequest(suite.router, "GET", "/api/v1/tasks", nil)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Task 1")
	assert.Contains(suite.T(), w.Body.String(), "Task 2")
}

func (suite *TaskHandlerTestSuite) TestGetAllTasks_Error() {
	suite.usecase.On("ListTasks").Return(nil, errors.New("some error"))

	w := performRequest(suite.router, "GET", "/api/v1/tasks", nil)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *TaskHandlerTestSuite) TestGetAllTasks_CustomError() {
	customErr := &domain.CustomeError{Message: "Custom error"}
	suite.usecase.On("ListTasks").Return(nil, customErr)

	w := performRequest(suite.router, "GET", "/api/v1/tasks", nil)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

// func (suite *TaskHandlerTestSuite) TestUpdateTask_Success() {
// 	updateDTO := dto.TaskUpdateDTO{
// 		ID:          uuid.New(),
// 		Title:       "Updated Task",
// 		Description: "Updated Description",
// 		Deadline:    time.Now(),
// 	}
//
// 	suite.usecase.On("UpdateTask", &updateDTO).Return(nil)
//
// 	w := performRequest(suite.router, "PUT", "/api/v1/tasks", updateDTO)
//
// 	assert.Equal(suite.T(), http.StatusCreated, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "task updated successfully")
// }

// func (suite *TaskHandlerTestSuite) TestUpdateTask_BindJSONError() {
// 	w := performRequest(suite.router, "PUT", "/api/v1/tasks", "invalid json")
//
// 	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "error")
// }
//
// func (suite *TaskHandlerTestSuite) TestUpdateTask_Error() {
// 	updateDTO := dto.TaskUpdateDTO{Title: "Updated Task", Description: "Updated Description"}
// 	suite.usecase.On("UpdateTask", &updateDTO).Return(errors.New("some error"))
//
// 	w := performRequest(suite.router, "PUT", "/api/v1/tasks", updateDTO)
//
// 	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "error")
// }
//
// func (suite *TaskHandlerTestSuite) TestUpdateTask_CustomError() {
// 	updateDTO := dto.TaskUpdateDTO{Title: "Updated Task", Description: "Updated Description"}
// 	customErr := &domain.CustomeError{Message: "Custom error"}
// 	suite.usecase.On("UpdateTask", &updateDTO).Return(customErr)
//
// 	w := performRequest(suite.router, "PUT", "/api/v1/tasks", updateDTO)
//
// 	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "Custom error")
// }

func performRequest(router *gin.Engine, method, url string, body interface{}) *httptest.ResponseRecorder {
	var jsonBody []byte
	if body != nil {
		var err error
		jsonBody, err = json.Marshal(body)
		if err != nil {
			panic(err)
		}
	}

	req, _ := http.NewRequest(method, url, bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestTaskHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskHandlerTestSuite))
}
