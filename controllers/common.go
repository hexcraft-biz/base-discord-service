package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hexcraft-biz/base-discord-service/config"
	"github.com/hexcraft-biz/controller"
)

type Common struct {
	*controller.Prototype
	Config config.ConfigInterface
}

func NewCommon(cfg config.ConfigInterface) *Common {
	return &Common{
		Prototype: controller.New("common", nil),
		Config:    cfg,
	}
}

func (ctrl *Common) NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
	}
}

func (ctrl *Common) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": http.StatusText(http.StatusOK)})
	}
}

func (ctrl *Common) ServerInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		guild, err := ctrl.Config.GetDG().Guild(ctrl.Config.GetDiscordServerID())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err})
		} else {
			c.JSON(http.StatusOK, guild)
		}
	}
}

func (ctrl *Common) WebsocketStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := ctrl.Config.GetDG().GatewayBot()
		if err == nil {
			c.JSON(http.StatusOK, resp)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		}
	}
}
