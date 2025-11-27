package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/services"
)

type CategoryHandler struct {
	category services.CategoryService
}

func NewCategoryHanlder(category services.CategoryService) *CategoryHandler {
	return &CategoryHandler{category: category}
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.category.CreateCategory(inputCategory)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetAllCategoryList(c *gin.Context) {

	category, err := h.category.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, category)
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	category, err := h.category.GetByID(uint(id))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, category)
}

func (h *CategoryHandler) UpdateCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var updateCategory models.UpdateCategoryRequest

	if err := c.ShouldBindJSON(&updateCategory); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	category, err := h.category.UpdateCategory(uint(id), updateCategory)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, category)
}
