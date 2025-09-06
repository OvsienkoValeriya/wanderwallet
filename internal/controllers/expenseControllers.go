package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"wanderwallet/internal/dto"
	"wanderwallet/internal/models"
	"wanderwallet/internal/services"

	"time"

	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
	expenseService  *services.ExpenseService
	categoryService *services.CategoryService
	travelService   *services.TravelService
}

func NewExpenseController(expenseService *services.ExpenseService, categoryService *services.CategoryService, travelService *services.TravelService) *ExpenseController {
	return &ExpenseController{
		expenseService:  expenseService,
		categoryService: categoryService,
		travelService:   travelService,
	}
}

// CreateExpense godoc
// @Summary Создать расход
// @Description Добавляет новый расход для авторизованного пользователя
// @Tags expenses
// @Accept json
// @Produce json
// @Param expense body dto.CreateExpenseRequest true "Данные расхода"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/expenses [post]
func (ctrl *ExpenseController) CreateExpense(c *gin.Context) {
	var req dto.CreateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userVal, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, ok := userVal.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user in context"})
		return
	}

	category, err := ctrl.categoryService.GetCategoryByName(req.Category)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found or unavailable"})
		return
	}

	travel, err := ctrl.travelService.GetTravelByID(req.TravelID)
	if err != nil || travel.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "invalid travel"})
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date"})
		return
	}

	expense := &models.Expense{
		UserID:      user.ID,
		CategoryID:  category.ID,
		TravelID:    req.TravelID,
		Amount:      req.Amount,
		CreatedAt:   date,
		Description: req.Comment,
	}

	if err := ctrl.expenseService.CreateExpense(expense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Expense with amount %.2f created", expense.Amount),
	})

}

// GetExpensesByUserID godoc
// @Summary Получить расходы пользователя
// @Description Возвращает список расходов текущего пользователя по категории и дате (опционально)
// @Tags expenses
// @Accept json
// @Produce json
// @Param category query string false "Категория"
// @Param from query string false "Дата начала, формат YYYY-MM-DD"
// @Param to query string false "Дата окончания, формат YYYY-MM-DD"
// @Success 200 {array} dto.ExpenseResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/expenses [get]
func (ctrl *ExpenseController) GetExpensesByUserID(c *gin.Context) {
	var req dto.GetUsersExpenseRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params"})
		return
	}

	userVal, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, _ := userVal.(models.User)

	var fromDate, toDate *time.Time
	if req.From != "" {
		t, err := time.Parse("2006-01-02", req.From)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from date"})
			return
		}
		fromDate = &t
	}
	if req.To != "" {
		t, err := time.Parse("2006-01-02", req.To)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid to date"})
			return
		}
		toDate = &t
	}

	var categoryID *uint
	if req.Category != "" {
		cat, err := ctrl.categoryService.GetCategoryByName(req.Category)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
			return
		}
		categoryID = &cat.ID
	}

	expenses, err := ctrl.expenseService.GetExpensesByUserTimeAndCategory(user.ID, fromDate, toDate, categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch expenses"})
		return
	}

	expenseResponses := make([]dto.ExpenseResponse, 0, len(expenses))
	for _, e := range expenses {
		category, _ := ctrl.categoryService.GetCategoryByID(e.CategoryID)
		expenseResponses = append(expenseResponses, dto.ExpenseResponse{
			ID:       fmt.Sprintf("%v", e.ID),
			Category: category.Name,
			Amount:   e.Amount,
			Date:     e.CreatedAt.Format("2006-01-02"),
			Comment:  e.Description,
		})
	}
	c.JSON(http.StatusOK, expenseResponses)

}

// UpdateExpenseByUserID godoc
// @Summary Обновить расход
// @Description Обновляет данные расхода по его ID для текущего пользователя
// @Tags expenses
// @Accept json
// @Produce json
// @Param id path int true "ID расхода"
// @Param expense body dto.UpdateExpenseRequest true "Новые данные расхода"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/expenses/{id} [put]
func (ctrl *ExpenseController) UpdateExpenseByUserID(c *gin.Context) {

	idStr := c.Param("id")
	expenseID, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expense ID"})
		return
	}

	var req dto.UpdateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userVal, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, ok := userVal.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user in context"})
		return
	}

	expense, err := ctrl.expenseService.GetExpenseByID(uint(expenseID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
		return
	}

	if expense.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "cannot edit another user's expense"})
		return
	}

	category, err := ctrl.categoryService.GetCategoryByName(req.Category)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	expenseDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date"})
		return
	}

	expense.CategoryID = category.ID
	expense.Amount = req.Amount
	expense.CreatedAt = expenseDate
	expense.Description = req.Comment

	if err := ctrl.expenseService.UpdateExpense(expense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Expense updated successfully",
	})
}

// DeleteExpenseByID godoc
// @Summary Удалить расход
// @Description Удаляет расход текущего пользователя по ID
// @Tags expenses
// @Accept json
// @Produce json
// @Param id path int true "ID расхода"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/expenses/{id} [delete]
func (ctrl *ExpenseController) DeleteExpenseByID(c *gin.Context) {
	idStr := c.Param("id")
	expenseID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expense ID"})
		return
	}

	userVal, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	user, ok := userVal.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user in context"})
		return
	}

	expense, err := ctrl.expenseService.GetExpenseByID(uint(expenseID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
		return
	}

	if expense.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "cannot delete another user's expense"})
		return
	}

	if err := ctrl.expenseService.DeleteExpense(uint(expenseID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Expense deleted successfully",
	})

}
