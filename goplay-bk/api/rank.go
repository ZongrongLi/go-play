package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tiancai110a/go-play/goplay-bk/service"
)

// DailyRank 每日排行
func DailyRank(c *gin.Context) {
	service := service.DailyRankService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
