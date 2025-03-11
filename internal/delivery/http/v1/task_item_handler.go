package v1

import (
	"net/http"

	"github.com/amha-mersha/icog_clean_code/internal/domain"
	"github.com/amha-mersha/icog_clean_code/internal/domain/dto"
	"github.com/amha-mersha/icog_clean_code/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type taskHandler struct {
	u usecase.TaskUsecase
}

var taskCnt taskHandler

func NewTaskHandler(r *gin.RouterGroup, u usecase.TaskUsecase) {
	taskCnt = taskHandler{
		u,
	}
	taskGroup := r.Group("tasks")
	{
		taskGroup.POST("", taskCnt.uploadTaskItem)
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

func (taskCnt *taskHandler) uploadTaskItem(ctx *gin.Context) {
	var newTaskItem dto.TaskCreateDTO

	if err := ctx.ShouldBindJSON(newTaskItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := taskCnt.u.CreateTask(&newTaskItem); err != nil {
		ctx.JSON(GetHTTPErrorCode(*err), gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": newTaskItem,
	})
}


func (taskCnt *taskHandler) GetTaskByID(ctx *gin.Context)  {
	rawID:= ctx.Param("id")
	taskID, err := uuid.Parse(rawID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}

}

func (taskCnt *taskHandler) GetAllTasks(ctx *gin.Context)  {
	task_id:= ctx.Param("id")
	if task_id
}
func (taskCnt *taskHandler) UpdateTask(ctx *gin.Context)  {
	task_id:= ctx.Param("id")
	if task_id
}
func (taskCnt *taskHandler) DeleteTask(ctx *gin.Context)  {
	task_id:= ctx.Param("id")
	if task_id
}
func (taskCnt *taskHandler) UpdateTaskStatus(ctx *gin.Context)  {
	task_id:= ctx.Param("id")
	if task_id
}
func (taskCnt *taskHandler) GetTasksByStatus(ctx *gin.Context)  {
	task_id:= ctx.Param("id")
	if task_id
}

