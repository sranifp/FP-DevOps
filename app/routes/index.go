package routes

import (
	"FP-DevOps/view"

	"github.com/gin-gonic/gin"
)

func Index(route *gin.Engine, index view.IndexView) {
	routes := route.Group("")
	{
		routes.GET("/", index.Index)
	}
}
