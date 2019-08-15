package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tiancai110a/go-play/goplay-bk/service"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		fmt.Println("==========", service)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
