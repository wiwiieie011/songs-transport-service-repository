package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/wiwiieie011/songs/services"
)

func RegisterRoutes(router *gin.Engine, songs services.SongService, category services.CategoryService) {
	songHandler := NewSongsHandler(songs)
	categoryHandler := NewCategoryHanlder(category)

	songHandler.RegisterRoutes(router)
	categoryHandler.RegisterRoutes(router)
}
