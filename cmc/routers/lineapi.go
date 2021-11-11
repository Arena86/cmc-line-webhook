package routers

import (
	"cmc/cmc/controllers"

	"github.com/gin-gonic/gin"
)

func LineAPIRoute(r *gin.RouterGroup) {
	r.POST("/webhook", controllers.WebHook)
}
