package routers

import (
	"scaffold-demo/routers/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	auth.RegisterSubRoutes(apiGroup)
}
