package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/services"
)

type CategoryHandler struct {
	category services.CategoryService
	logger   *logrus.Logger
}

func NewCategoryHanlder(category services.CategoryService, logger *logrus.Logger) *CategoryHandler {
	return &CategoryHandler{
		category: category,
		logger:   logger,
	}

}

func (h *CategoryHandler) RegisterRoutes(r *gin.Engine) {
	categoryGroup := r.Group("/category")

	{
		categoryGroup.GET("/", h.GetAllCategoryList)
		categoryGroup.GET("/:id", h.GetCategoryByID)
		categoryGroup.POST("/", h.CreateCategory)
		categoryGroup.PATCH("/:id", h.UpdateCategoryByID)
	}

}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var inputCategory models.CreateCategoryRequest

	if err := c.ShouldBindJSON(&inputCategory); err != nil {
		h.logger.WithError(err).Warn("error is not json form")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.category.CreateCategory(inputCategory)
	if err != nil {
		h.logger.WithError(err).Error("create category error")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.WithField("create category:", category).Info("Succes created")
	c.IndentedJSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetAllCategoryList(c *gin.Context) {

	category, err := h.category.GetAll()
	if err != nil {
		h.logger.WithError(err).Error("category list launch fail")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("category list launch succes")
	c.IndentedJSON(http.StatusOK, category)
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.WithError(err).Warn("error parse id")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	category, err := h.category.GetByID(uint(id))
	if err != nil {
		h.logger.WithError(err).Error("category id not found")
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	h.logger.WithField("category_id:", category.ID).Info("category found")
	c.IndentedJSON(http.StatusOK, category)
}

func (h *CategoryHandler) UpdateCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.WithError(err).Warn("error parse id")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var updateCategory models.UpdateCategoryRequest

	if err := c.ShouldBindJSON(&updateCategory); err != nil {
		h.logger.WithError(err).Warn("error: not json format")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	category, err := h.category.UpdateCategory(uint(id), updateCategory)
	if err != nil {
		h.logger.WithError(err).Error("error: update fail")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.WithField("category_id: ", category.ID).Info("category updated succes")
	c.IndentedJSON(http.StatusOK, category)
}
