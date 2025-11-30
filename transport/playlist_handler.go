package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/services"
)

type PlayListHandler struct {
	playlist services.PlayListServices
}

func NewPlayListHandler(playlist services.PlayListServices) *PlayListHandler {
	return &PlayListHandler{playlist: playlist}
}

func (h *PlayListHandler) RegisterRoutes(r *gin.Engine) {
	playlistGroup := r.Group("/playlist")
	{
		playlistGroup.POST("/", h.Create)
		playlistGroup.GET("/", h.GetList)
		playlistGroup.GET("/:id", h.GetByID)
		playlistGroup.PATCH("/:id",h.UpdateByID)
		playlistGroup.DELETE("/:id", h.DeleteByID)
	}
}

func (h *PlayListHandler) Create(r *gin.Context) {
	var inputPlaylist models.CreatePlayListRequest

	if err := r.ShouldBindJSON(&inputPlaylist); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playlist, err := h.playlist.CreatePlayList(inputPlaylist)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.IndentedJSON(http.StatusCreated, playlist)
}

func (h *PlayListHandler) GetList(r *gin.Context) {
	playlist, err := h.playlist.GetAllPlaylists()
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.IndentedJSON(http.StatusOK, playlist)
}

func (h *PlayListHandler) GetByID(r *gin.Context) {
	idStr := r.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playlist, err := h.playlist.GetPlaylistByID(uint(id))
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.IndentedJSON(http.StatusOK, playlist)
}

func (h *PlayListHandler) UpdateByID(r *gin.Context) {
	idStr := r.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var upPlaylist models.UpdatePlayListRequest
	if err := r.ShouldBindJSON(&upPlaylist); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playlist, err := h.playlist.UpdatePlaylistByID(uint(id), upPlaylist)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.IndentedJSON(http.StatusOK, playlist)
}

func (h *PlayListHandler) DeleteByID(r *gin.Context) {
	idStr := r.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err1 := h.playlist.DeletePlayList(uint(id))
	if err1 != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}

	r.IndentedJSON(http.StatusOK, gin.H{"status": true})
}
