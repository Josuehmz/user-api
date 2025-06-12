package handler

import (
	"context"
	"net/http"

	"github.com/Josuehmz/user-api/internal/dto"
	"github.com/Josuehmz/user-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

// UserHandler define los controladores HTTP para las rutas de usuario.
type UserHandler struct {
	userUC usecase.UserUseCaseInterface
}

// NewUserHandler crea una nueva instancia de UserHandler.
func NewUserHandler(userUC usecase.UserUseCaseInterface) *UserHandler {
	return &UserHandler{userUC: userUC}
}

// CreateUser godoc
// @Summary      Crear usuario
// @Description  Crea un nuevo usuario
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.UserDTO  true  "Usuario a crear"
// @Success      201   {object}  dto.UserDTO
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input dto.UserDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user := dto.ToDomain(&input)

	if err := h.userUC.CreateUser(context.Background(), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.FromDomain(user))
}

// GetAllUsers godoc
// @Summary      Listar usuarios
// @Description  Obtiene todos los usuarios
// @Tags         users
// @Produce      json
// @Success      200   {array}   dto.UserDTO
// @Failure      500   {object}  map[string]string
// @Router       /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUC.GetAllUsers(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var dtos []dto.UserDTO
	for _, u := range users {
		dtos = append(dtos, *dto.FromDomain(&u))
	}

	c.JSON(http.StatusOK, dtos)
}

// GetUserByID godoc
// @Summary      Obtener usuario por ID
// @Description  Retorna un usuario dado su ID
// @Tags         users
// @Produce      json
// @Param        id    path      string  true  "ID del usuario"
// @Success      200   {object}  dto.UserDTO
// @Failure      404   {object}  map[string]string
// @Router       /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userUC.GetUserByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, dto.FromDomain(user))
}

// PatchUser godoc
// @Summary      Actualizar parcialmente un usuario
// @Description  Actualiza solo los campos enviados
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      string           true  "ID del usuario"
// @Param        user  body      dto.UserPatchDTO true  "Campos a actualizar"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users/{id} [patch]
func (h *UserHandler) PatchUser(c *gin.Context) {
	id := c.Param("id")
	var patch dto.UserPatchDTO

	if err := c.ShouldBindJSON(&patch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	fields := patch.ToUpdateMap()
	if len(fields) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no fields provided"})
		return
	}

	if err := h.userUC.PatchUser(context.Background(), id, fields); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user partially updated"})
}

// DeleteUser godoc
// @Summary      Eliminar usuario
// @Description  Elimina un usuario por ID
// @Tags         users
// @Produce      json
// @Param        id    path      string  true  "ID del usuario"
// @Success      200   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := h.userUC.DeleteUser(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
