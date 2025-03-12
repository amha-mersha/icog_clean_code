package v1

import (
	"errors"
	"net/http"

	"github.com/amha-mersha/icog_clean_code/internal/domain"
	"github.com/amha-mersha/icog_clean_code/internal/domain/dto"
	"github.com/amha-mersha/icog_clean_code/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskHandler struct {
	u usecase.TaskUsecase
}

var taskCnt TaskHandler

func NewTaskHandler(u usecase.TaskUsecase) TaskHandler {
	return TaskHandler{
		u,
	}
}
func GetHTTPErrorCode(err domain.CustomeError) int {
	switch err.ErrCode() {
	case domain.ERR_BAD_REQUEST:
		return http.StatusBadRequest
	case domain.ERR_UNAUTHORIZED:
		return http.StatusUnauthorized
	case domain.ERR_FORBIDDEN:
		return http.StatusForbidden
	case domain.ERR_NOT_FOUND:
		return http.StatusNotFound
	case domain.ERR_CONFLICT:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func (taskCnt *TaskHandler) UploadTaskItem(ctx *gin.Context) {
	var newTaskItem dto.TaskCreateDTO

	if err := ctx.ShouldBindJSON(&newTaskItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := taskCnt.u.CreateTask(&newTaskItem); err != nil {
		var customErr *domain.CustomeError
		if errors.As(err, &customErr) {
			ctx.JSON(GetHTTPErrorCode(*customErr), gin.H{"error": customErr.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": newTaskItem,
	})
}

func (taskCnt *TaskHandler) GetTaskByID(ctx *gin.Context) {
	rawID := ctx.Param("id")
	taskID, err := uuid.Parse(rawID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}
	task, errUc := taskCnt.u.GetTask(taskID)
	if errUc != nil {
		var customErr *domain.CustomeError
		if errors.As(errUc, &customErr) {
			ctx.JSON(GetHTTPErrorCode(*customErr), gin.H{"error": customErr.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errUc.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": task,
	})
}

func (h *TaskHandler) GetAllTasks(ctx *gin.Context) {
	tasks, err := h.u.ListTasks()
	if err != nil {
		var customErr *domain.CustomeError
		if errors.As(err, &customErr) {
			ctx.JSON(GetHTTPErrorCode(*customErr), gin.H{"error": customErr.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (taskCnt *TaskHandler) UpdateTask(ctx *gin.Context) {
	var newTaskItem dto.TaskUpdateDTO

	if err := ctx.ShouldBindJSON(&newTaskItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := taskCnt.u.UpdateTask(&newTaskItem); err != nil {
		var customErr *domain.CustomeError
		if errors.As(err, &customErr) {
			ctx.JSON(GetHTTPErrorCode(*customErr), gin.H{"error": customErr.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "task updated successfuly",
	})
}

func (taskCnt *TaskHandler) DeleteTask(ctx *gin.Context) {
	rawID := ctx.Param("id")
	taskID, err := uuid.Parse(rawID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}
	errUc := taskCnt.u.DeleteTask(taskID)
	if errUc != nil {
		var customErr *domain.CustomeError
		if errors.As(errUc, &customErr) {
			ctx.JSON(GetHTTPErrorCode(*customErr), gin.H{"error": customErr.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errUc.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "successfuly deleted",
	})
}

func (taskCnt *TaskHandler) GetTasksByStatus(ctx *gin.Context) {
	statusID := ctx.Query("status")
	if !dto.ValidStatus(statusID) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task status"})
		return
	}
}
