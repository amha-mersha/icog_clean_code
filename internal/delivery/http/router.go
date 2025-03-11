package http

import (
	v1 "github.com/amha-mersha/icog_clean_code/internal/delivery/http/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter(engine *gin.Engine) {
	apiV1Group := engine.Group("/v1")
	v1.NewTaskHandler(apiV1Group)
}
