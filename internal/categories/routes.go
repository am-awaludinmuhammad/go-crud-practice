package categories

import "github.com/gin-gonic/gin"

func RegisterCategoryRoutes(router *gin.Engine, handler *CategoryHandler) {
	categoryRoutes := router.Group("/categories")
	{
		categoryRoutes.GET("/", handler.GetAllCategories)
		categoryRoutes.POST("/", handler.CreateCategory)
		categoryRoutes.GET("/:id", handler.GetCategoryByID)
		categoryRoutes.DELETE("/:id", handler.DestroyCategory)
	}
}
