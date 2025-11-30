package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/wiwiieie011/songs/services"
)

func RegisterRoutes(
	router *gin.Engine, 
	songs services.SongService, 
	category services.CategoryService , 
	user services.UserServices,
	playlist services.PlayListServices,
	playlistitem services.PlayListItemsServices,

	) {
	songHandler := NewSongsHandler(songs)
	categoryHandler := NewCategoryHanlder(category)
	userHandler := NewUserHandler(user)
	playlistHandler := NewPlayListHandler(playlist)
	playlistitemHandler := NewPlayListItemsHandler(playlistitem)

	songHandler.RegisterRoutes(router)
	categoryHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
	playlistHandler.RegisterRoutes(router)
	playlistitemHandler.RegisterRoutes(router)
}
