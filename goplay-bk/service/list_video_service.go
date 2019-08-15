package service

import (
	"github.com/tiancai110a/go-play/goplay-bk/model"
	"github.com/tiancai110a/go-play/goplay-bk/serializer"
)

// ListVideoService 视频列表服务

type ListVideoService struct {
	Limit  int `form:"limit"`
	Offset int `form:"start"`
}

// List 视频列表
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video

	err := model.DB.Find(&videos).
		Limit(service.Limit).Offset(service.Offset).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
