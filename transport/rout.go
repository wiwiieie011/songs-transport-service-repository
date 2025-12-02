package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wiwiieie011/songs/services"
)

func RegisterRoutes(
	router *gin.Engine, 
	songs services.SongService, 
	category services.CategoryService , 
	user services.UserServices,
	playlist services.PlayListServices,
	playlistitem services.PlayListItemsServices,
	log *logrus.Logger,
	) {
	songHandler := NewSongsHandler(songs,log)
	categoryHandler := NewCategoryHanlder(category , log)
	userHandler := NewUserHandler(user,log)
	playlistHandler := NewPlayListHandler(playlist, log)
	playlistitemHandler := NewPlayListItemsHandler(playlistitem, log)

	songHandler.RegisterRoutes(router)
	categoryHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
	playlistHandler.RegisterRoutes(router)
	playlistitemHandler.RegisterRoutes(router)
}
