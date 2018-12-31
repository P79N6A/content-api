package routes

import (
	"code.byted.org/learning_fe/pathfinder-api/handler"

	"github.com/gin-gonic/gin"
)

// Hello ...
func Content(r *gin.Engine) {
	r.GET("/contents/:contentId", handler.GetContent)
	r.GET("/contents", handler.ListContent)
	r.POST("/contents", handler.CreateContent)
	r.PATCH("/contents/:contentId", handler.UpdateContent)
	r.DELETE("/contents", handler.DeleteContent)
}