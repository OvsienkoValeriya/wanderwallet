package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"wanderwallet/internal/dto"
	"wanderwallet/internal/models"
	"wanderwallet/internal/repository"
	"wanderwallet/internal/services"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService *services.CategoryService
	expenseService  *services.ExpenseService
}

func NewCategoryController(categoryService *services.CategoryService, expenseService *services.ExpenseService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
		expenseService:  expenseService,
	}
}

// GetCategoriesByUserID godoc
// @Summary Получить категории пользователя
// @Description Возвращает список категорий расходов для текущего пользователя
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {array} dto.CategoryResponse
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/categories [get]
func (ctrl *CategoryController) GetCategoriesByUserID(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	ctx := c.Request.Context()
	categories, err := ctrl.categoryService.GetAllCategories(ctx, user.ID)
	if err != nil {
		log.Printf("Failed to get categories for user %d: %v\n", user.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	categoryResponses := make([]dto.CategoryResponse, 0, len(categories))
	for _, cat := range categories {
		categoryResponses = append(categoryResponses, dto.CategoryResponse{
			ID:      fmt.Sprintf("%v", cat.ID),
			Name:    cat.Name,
			Builtin: cat.Builtin,
		})
	}
	c.JSON(http.StatusOK, categoryResponses)
}

// CreateCategory godoc
// @Summary Создать категорию
// @Description Добавляет новую категорию расходов для авторизованного пользователя
// @Tags categories
// @Accept json
// @Produce json
// @Param category body dto.CreateCategoryRequest true "Данные категории"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/categories [post]
func (ctrl *CategoryController) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user := c.MustGet("user").(models.User)
	ctx := c.Request.Context()

	category := &models.Category{
		UserID:  &user.ID,
		Name:    req.Name,
		Builtin: false,
	}

	if err := ctrl.categoryService.CreateCategory(ctx, category); err != nil {
		if errors.Is(err, repository.ErrCategoryExists) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "category already exists"})
			return
		}
		log.Printf("Failed to create category for user %d: %v\n", user.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Category %s created", category.Name),
	})
}

// DeleteCategoryByID godoc
// @Summary Удалить категорию
// @Description Удаляет пользовательскую категорию по ID, если она не системная и не используется
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "ID категории"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/categories/{id} [delete]
func (ctrl *CategoryController) DeleteCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	categoryID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	user := c.MustGet("user").(models.User)
	ctx := c.Request.Context()
	category, err := ctrl.categoryService.GetCategoryByID(ctx, uint(categoryID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	if category.Builtin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete builtin category"})
		return
	}

	if category.UserID == nil || *category.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "cannot delete another user's category"})
		return
	}

	inUse, err := ctrl.expenseService.ExistsByCategoryID(ctx, uint(categoryID))
	if err != nil {
		log.Printf("Failed to check category usage for category %d: %v\n", categoryID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	if inUse {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category is used in expenses"})
		return
	}

	if err := ctrl.categoryService.DeleteCategory(ctx, uint(categoryID)); err != nil {
		log.Printf("Failed to delete category %d: %v\n", categoryID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Category %s deleted successfully", category.Name),
	})
}
