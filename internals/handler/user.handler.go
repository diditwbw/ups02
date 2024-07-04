package handler

import (
	"net/http"
	"ups02/internals/services"
	"ups02/x/interfacesx"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Userhandler struct {
	userService services.UserService
	validate    *validator.Validate
}

func NewUserhandler(userService services.UserService) *Userhandler {
	return &Userhandler{
		userService: userService,
		validate:    validator.New(),
	}
}

func (h *Userhandler) CreateUser(c *gin.Context) {
	var userRequest interfacesx.UserRegistrationRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}

	if err := h.validate.Struct(userRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}

	userData, err := h.userService.CreateUserAccount(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusInternalServerError,
		})

		return
	}

	c.JSON(http.StatusBadRequest, interfacesx.UserResponse{
		Message: "user created succesfully",
		Status:  interfacesx.StatusError,
		Code:    http.StatusBadRequest,
		Data:    *userData,
	})
}
