package controllers

import (
	"fmt"
	"net/http"
	"wanderwallet/internal/dto"
	"wanderwallet/internal/models"
	"wanderwallet/internal/services"

	"time"

	"github.com/gin-gonic/gin"
)

type TravelController struct {
	travelService *services.TravelService
}

func NewTravelController(travelService *services.TravelService) *TravelController {
	return &TravelController{
		travelService: travelService,
	}
}

// CreateTravel godoc
// @Summary Создать новое путешествие
// @Description Создает запись о путешествии для авторизованного пользователя
// @Tags travel
// @Accept json
// @Produce json
// @Param travel body dto.CreateTravelRequest true "Данные путешествия"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/travel [post]
func (ctrl *TravelController) CreateTravel(c *gin.Context) {
	var req dto.CreateTravelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
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

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date"})
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date"})
		return
	}

	travel, err := ctrl.travelService.CreateTravel(user.ID, req.Title, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Travel %s created", travel.Title),
	})
}
