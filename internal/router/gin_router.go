package router

import (
	"github.com/Josuehmz/user-api/internal/handler"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter configura y retorna el enrutador principal de la aplicación.
// Define las rutas del CRUD de usuarios y la documentación Swagger.
func NewRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	// Ruta para acceder a la documentación Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Grupo de rutas relacionadas con usuarios
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userHandler.CreateUser)       // Crear usuario
		userRoutes.GET("", userHandler.GetAllUsers)       // Listar usuarios
		userRoutes.GET("/:id", userHandler.GetUserByID)   // Obtener usuario por ID
		userRoutes.PATCH("/:id", userHandler.PatchUser)   // Actualizar parcialmente
		userRoutes.DELETE("/:id", userHandler.DeleteUser) // Eliminar usuario
	}

	return r
}
