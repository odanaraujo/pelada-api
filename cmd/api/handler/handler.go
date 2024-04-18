package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/pelada-api/internal/application/controller"
)

func InitHandler(r *gin.RouterGroup, controller controller.Playercontroller) {
	r.POST("/players/confirm-presence", controller.ConfirmePresenceController)
}
