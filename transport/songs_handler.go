package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/services"
)

type SongHandler struct {
	songs services.SongService
}

func NewSongsHandler(songs services.SongService) *SongHandler {
	return &SongHandler{songs: songs}
}

func (h *SongHandler) RegisterRoutes(r *gin.Engine) {
	songGroup := r.Group("/songs")

	{
		songGroup.GET("/", h.GetAllSongsList)
		songGroup.GET("/:id", h.GetSongByID)
		songGroup.GET("/:id/category", h.GetSongsByCategoryiD)
		songGroup.POST("/", h.CreateSong)
		songGroup.PATCH("/:id", h.UpdateSong)
		songGroup.DELETE("/:id", h.DeleteSong)
	}

}

func (h *SongHandler) CreateSong(c *gin.Context) {
	var inputSong models.CreateSongRequest

	if err := c.ShouldBindJSON(&inputSong); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song, err := h.songs.CreateSong(inputSong)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, song)
}

func (h *SongHandler) GetSongByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song, err := h.songs.GetSongByID(uint(id))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, song)
}

func (h *SongHandler) GetAllSongsList(c *gin.Context) {
	list, err := h.songs.GetSongs()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, list)
}

func (h *SongHandler) GetSongsByCategoryiD(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	songs, err := h.songs.GetSongsByCategoryiD(uint(id))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, songs)
}

func (h *SongHandler) UpdateSong(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updateSong models.UpdateSongRequest

	if err := c.ShouldBindJSON(&updateSong); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	song, err := h.songs.UpdateSong(uint(id), updateSong)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, song)
}

func (h *SongHandler) DeleteSong(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err1 := h.songs.DeleteSong(uint(id))
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"deleted": true})
}
