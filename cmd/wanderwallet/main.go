package main

import (
	"wanderwallet/initializers"
	"wanderwallet/internal/controllers"
	"wanderwallet/internal/middleware"
	"wanderwallet/internal/repository"
	"wanderwallet/internal/services"

	_ "wanderwallet/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(middleware.DebugMiddleware())
	r.Use(middleware.AuthMiddleware)

	userRepo := repository.NewUserRepository(initializers.DB)
	travelRepo := repository.NewTravelRepository(initializers.DB)
	categoryRepo := repository.NewCategoryRepository(initializers.DB)
	expenseRepo := repository.NewExpenseRepository(initializers.DB)

	userService := services.NewUserService(userRepo)
	travelService := services.NewTravelService(travelRepo)
	categoryService := services.NewCategoryService(categoryRepo, expenseRepo)
	expenseService := services.NewExpenseService(expenseRepo)

	userController := controllers.NewUserController(userService)
	travelController := controllers.NewTravelController(travelService)
	expenseController := controllers.NewExpenseController(expenseService, categoryService, travelService)
	categoryController := controllers.NewCategoryController(categoryService, expenseService)

	api := r.Group("/api")
	{
		userRoutes := api.Group("/auth")
		{
			userRoutes.POST("/register", userController.Register)
			userRoutes.POST("/login", userController.Login)
		}

		travelRoutes := api.Group("/travel")
		{
			travelRoutes.POST("", travelController.CreateTravel)
		}

		expenseRoutes := api.Group("/expenses")
		{
			expenseRoutes.GET("", expenseController.GetExpensesByUserID)
			expenseRoutes.POST("", expenseController.CreateExpense)
			expenseRoutes.PUT("/:id", expenseController.UpdateExpenseByUserID)
			expenseRoutes.DELETE("/:id", expenseController.DeleteExpenseByID)
		}

		categoryRoutes := api.Group("/categories")
		{
			categoryRoutes.GET("", categoryController.GetCategoriesByUserID)
			categoryRoutes.POST("", categoryController.CreateCategory)
			categoryRoutes.DELETE("/:id", categoryController.DeleteCategoryByID)
		}
	}
	r.Run()
}
