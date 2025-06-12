package router

import (
	"github.com/Josuehmz/user-api/internal/handler"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userHandler.CreateUser)
		userRoutes.GET("", userHandler.GetAllUsers)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.PATCH("/:id", userHandler.PatchUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	return r
}
