package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/services"
)

type UserHandler struct {
	user   services.UserServices
	logger *logrus.Logger
}

func NewUserHandler(user services.UserServices, logger *logrus.Logger) *UserHandler {
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
		h.logger.WithError(err).Warn("error json format or type information")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.user.CreateUser(inputUser)
	if err != nil {
		h.logger.WithError(err).Error("regist user failed")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.WithField("user:", user).Info("user register succces")
	c.IndentedJSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.WithError(err).Warn("error parse id or invalid id")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.user.GetUserByID(uint(id))
	if err != nil {
		logrus.WithField("user_id", id).Error("User not found")
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	logrus.WithField("user_id", id).Info("User launch succesf")
	c.IndentedJSON(http.StatusOK, user)
}

func (h *UserHandler) GetListUsers(c *gin.Context) {
	users, err := h.user.GetAllUsers()
	if err != nil {
		h.logger.WithError(err).Error("error launch list users")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("user list launch succes")
	c.IndentedJSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.WithError(err).Warn("error parse id or invalid id")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var upUser models.UpdateUserRequest
	if err := c.ShouldBindJSON(&upUser); err != nil {
		h.logger.WithError(err).Warn("error json formate or type information")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := h.user.UpdatsUser(uint(id), upUser)
	if err != nil {
		h.logger.WithError(err).Error("not found user or failed update")
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	h.logger.WithField("user_id: ", users.ID).Info("user updated succes")
	c.IndentedJSON(http.StatusOK, users)
}

func (h *UserHandler) DeleteUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.WithError(err).Warn("error parse id or invalid id")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.user.DeleteUser(uint(id)); err != nil {
		h.logger.WithError(err).Error("not found user or failed delete")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("User deleted succes")
	c.IndentedJSON(http.StatusOK, gin.H{"status": true})
}
