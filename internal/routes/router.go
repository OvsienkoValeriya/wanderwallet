package routes

import (
	"wanderwallet/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	r *gin.Engine,
	userController *controllers.UserController,
	travelController *controllers.TravelController,
	expenseController *controllers.ExpenseController,
	categoryController *controllers.CategoryController,
	analyticsController *controllers.AnalyticsController,
) {

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

		analyticsRoutes := api.Group("/analytics")
		{
			analyticsRoutes.GET("", analyticsController.GetAnalytics)
		}
	}
}
