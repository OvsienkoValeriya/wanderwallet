package controllers

import (
	"log"
	"net/http"
	"wanderwallet/internal/dto"
	"wanderwallet/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Register godoc
// @Summary Регистрация пользователя
// @Description Регистрирует нового пользователя и устанавливает куку авторизации
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.UserRequest true "Данные пользователя"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/register [post]
func (ctrl *UserController) Register(c *gin.Context) {
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request format",
		})
		return
	}
	response, err := ctrl.userService.Register(req.Login, req.Password)
	if err != nil {
		switch err {
		case services.ErrUserAlreadyExists:
			c.JSON(http.StatusConflict, gin.H{
				"error": "login already exists",
			})
		default:
			log.Printf("Registration error: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
		}
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", response.Token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "user registered and authenticated successfully",
	})
}

// Login godoc
// @Summary Вход пользователя
// @Description Аутентифицирует пользователя и устанавливает куку авторизации
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.UserRequest true "Данные пользователя для входа"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/login [post]
func (ctrl *UserController) Login(c *gin.Context) {
	var req dto.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request format",
		})
		return
	}

	response, err := ctrl.userService.Login(req.Login, req.Password)
	if err != nil {
		switch err {
		case services.ErrUserNotFound, services.ErrInvalidPassword:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid login or password",
			})
		default:
			log.Printf("Login error: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
		}
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", response.Token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "user authenticated successfully",
	})
}
