package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"wanderwallet/internal/models"
	"wanderwallet/internal/services"

	"github.com/gin-gonic/gin"
)

type AnalyticsController struct {
	expenseService   *services.ExpenseService
	analyticsService *services.AnalyticsService
}

func NewAnalyticsController(expenseService *services.ExpenseService, analyticsService *services.AnalyticsService) *AnalyticsController {
	return &AnalyticsController{
		expenseService:   expenseService,
		analyticsService: analyticsService,
	}
}

// GetAnalytics godoc
// @Summary Получение агрегированной аналитики
// @Description Возвращает суммы по категориям, динамику по датам и общую сумму расходов за период
// @Tags analytics
// @Accept json
// @Produce json
// @Param travel_id query int true "ID путешествия"
// @Param from query string false "Дата начала, формат YYYY-MM-DD"
// @Param to query string false "Дата окончания, формат YYYY-MM-DD"
// @Success 200 {object} dto.AnalyticsResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/analytics [get]
// @Security ApiKeyAuth
// @Security BearerAuth
func (ctrl *AnalyticsController) GetAnalytics(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	travelIDStr := c.Query("travel_id")
	travelIDUint64, err := strconv.ParseUint(travelIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid travel_id"})
		return
	}
	travelID := uint(travelIDUint64)

	fromStr := c.Query("from")
	toStr := c.Query("to")

	var from, to time.Time
	if fromStr != "" {
		from, err = time.Parse("2006-01-02", fromStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from date"})
			return
		}
	}
	if toStr != "" {
		to, err = time.Parse("2006-01-02", toStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid to date"})
			return
		}
	}
	ctx := c.Request.Context()
	resp, err := ctrl.analyticsService.Aggregate(ctx, user.ID, travelID, from, to)
	if err != nil {
		log.Printf("Analytics aggregation error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, resp)
}
