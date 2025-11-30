package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/services"
)

type UserHandler struct {
	user services.UserServices
}

func NewUserHandler(user services.UserServices) *UserHandler {
	return &UserHandler{user: user}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/", h.CreateUser)
		userGroup.GET("/", h.GetListUsers)
		userGroup.GET("/:id", h.GetUserByID)
		userGroup.PATCH("/:id", h.UpdateUserByID)
		userGroup.DELETE("/:id", h.DeleteUserByID)
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var inputUser models.CreateUserRequest

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.user.CreateUser(inputUser)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.user.GetUserByID(uint(id))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (h *UserHandler) GetListUsers(c *gin.Context) {
	users, err := h.user.GetAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var upUser models.UpdateUserRequest
	if err := c.ShouldBindJSON(&upUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := h.user.UpdatsUser(uint(id), upUser)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (h *UserHandler) DeleteUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.user.DeleteUser(uint(id)); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": true})
}
