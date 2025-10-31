package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/example/user-api/model"
	"github.com/example/user-api/service"
)

// UserController handles HTTP requests for user operations
type UserController struct {
	userService service.UserService
}

// NewUserController creates a new instance of UserController
func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// CreateUser handles POST /users
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	if err := c.userService.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, Response{
		Success: true,
		Data:    user,
	})
}

// GetUserByID handles GET /users/:id
func (c *UserController) GetUserByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "invalid user ID",
		})
		return
	}

	user, err := c.userService.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Data:    user,
	})
}

// GetAllUsers handles GET /users
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Data:    users,
	})
}

// UpdateUser handles PUT /users/:id
func (c *UserController) UpdateUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "invalid user ID",
		})
		return
	}

	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	user.ID = id
	if err := c.userService.UpdateUser(&user); err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, Response{
				Success: false,
				Error:   err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Data:    user,
	})
}

// DeleteUser handles DELETE /users/:id
func (c *UserController) DeleteUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "invalid user ID",
		})
		return
	}

	if err := c.userService.DeleteUser(id); err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, Response{
				Success: false,
				Error:   err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Data:    "user deleted successfully",
	})
}
