package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/services"
)

type SongHandler struct {
	songs  services.SongService
	logger *logrus.Logger
}

func NewSongsHandler(songs services.SongService, logger *logrus.Logger) *SongHandler {
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
		h.logger.WithError(err).Warn("error invalid json")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song, err := h.songs.CreateSong(inputSong)
	if err != nil {
		h.logger.WithError(err).Error("create song failed")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.WithField("song: ", song).Info("create user succes")
	c.IndentedJSON(http.StatusOK, song)
}

func (h *SongHandler) GetSongByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.WithError(err).Warn("failed parse id")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song, err := h.songs.GetSongByID(uint(id))
	if err != nil {
		h.logger.WithError(err).Error("song found fail")
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	h.logger.WithField("song_id: ", song.ID).Info("song found succes")
	c.IndentedJSON(http.StatusOK, song)
}

func (h *SongHandler) GetAllSongsList(c *gin.Context) {
	list, err := h.songs.GetSongs()
	if err != nil {
		h.logger.WithError(err).Error("list launch failed")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("song list get succes")
	c.IndentedJSON(http.StatusOK, list)
}

func (h *SongHandler) GetSongsByCategoryiD(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.WithError(err).Warn("parse id or id invalid")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	songs, err := h.songs.GetSongsByCategoryiD(uint(id))
	if err != nil {
		h.logger.WithError(err).Error("error found songs by category")
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("songs category found succes")
	c.IndentedJSON(http.StatusOK, songs)
}

func (h *SongHandler) UpdateSong(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.WithError(err).Warn("parse id or id invalid")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updateSong models.UpdateSongRequest

	if err := c.ShouldBindJSON(&updateSong); err != nil {
		h.logger.WithError(err).Warn("error json format")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	song, err := h.songs.UpdateSong(uint(id), updateSong)
	if err != nil {
		h.logger.WithError(err).Error("eror update song")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.WithField("song_id: ", song.ID).Info("song updated succes")
	c.IndentedJSON(http.StatusOK, song)
}

func (h *SongHandler) DeleteSong(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.WithError(err).Warn("parse id or id invalid")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err1 := h.songs.DeleteSong(uint(id))
	if err1 != nil {
		h.logger.WithError(err).Error("delete song failed")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	h.logger.Info("song deleted succes")
	c.IndentedJSON(http.StatusNoContent, gin.H{"deleted": true})
}
