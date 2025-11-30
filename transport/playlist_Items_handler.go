package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/services"
)

type PlayListItemsHandler struct {
	pl_item services.PlayListItemsServices
}

func NewPlayListItemsHandler(pl_item services.PlayListItemsServices) *PlayListItemsHandler {
	return &PlayListItemsHandler{pl_item: pl_item}
}

func (h *PlayListItemsHandler) RegisterRoutes(r *gin.Engine) {
	pl_itemGroup := r.Group("/playlist_items")
	{
		pl_itemGroup.POST("/", h.CreateTrack)
		pl_itemGroup.GET("/:playlist_id/song/:song_id", h.GetByID)
		pl_itemGroup.DELETE("/:playlist_id/song/:song_id", h.DeletePlayListItem)
	}
}

func (h *PlayListItemsHandler) CreateTrack(c *gin.Context) {
	var inputTrackInfos models.CreatePlayListItemsRequest

	if err := c.ShouldBindJSON(&inputTrackInfos); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	track, err := h.pl_item.CreatePlayListItems(inputTrackInfos)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, track)
}

func (h *PlayListItemsHandler) GetByID(c *gin.Context) {
	pidStr := c.Param("playlist_id")
	pid, err := strconv.ParseUint(pidStr, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sidStr := c.Param("song_id")
	sid, err := strconv.ParseUint(sidStr, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trackInfo, err := h.pl_item.GetByID(uint(pid), uint(sid))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, trackInfo)
}

func (h *PlayListItemsHandler) DeletePlayListItem(c *gin.Context) {
	pidParam := c.Param("playlist_id")
	sidParam := c.Param("song_id")

	pid, err1 := strconv.ParseUint(pidParam, 10, 64)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid playlist id"})
		return
	}

	sid, err2 := strconv.ParseUint(sidParam, 10, 64)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid song id"})
		return
	}

	if err3 := h.pl_item.Delete(uint(pid), uint(sid)); err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err3.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "song removed from playlist"})
}
