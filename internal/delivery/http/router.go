package delivery

import (
	"fmt"

	v1 "github.com/amha-mersha/icog_clean_code/internal/delivery/http/v1"
	"github.com/amha-mersha/icog_clean_code/internal/repository"
	"github.com/amha-mersha/icog_clean_code/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(database *gorm.DB, routerVersion string, port int) {
	router := gin.Default()
	taskRepository := repository.NewTaskRepo(database)
	taskUseCase := usecase.NewTaskUseCase(&taskRepository)
	taskController := v1.NewTaskHandler(&taskUseCase)

	taskRouter := router.Group("/api/" + routerVersion + "/tasks")
	taskRouter.POST("", taskController.UploadTaskItem)
	taskRouter.GET("", taskController.GetAllTasks)
	taskRouter.GET("/:id", taskController.GetTaskByID)
	taskRouter.PUT("", taskController.UpdateTask)
	taskRouter.DELETE("/:id", taskController.DeleteTask)
	taskRouter.GET("?status", taskController.GetTasksByStatus)

	router.Run(fmt.Sprintf(":%v", port))
}
