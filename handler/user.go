package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	// perlu service interface
	userService user.Service
}

func NewuserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserIput
	// struct diatas kita passing sebagai parameter service

	var input user.RegisteruserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	// panggil formatter
	formatter := user.FotmatUser(newUser, "initokenrahasia")
	// panggil helper untuk response
	response := helper.APIResponse("Account has been created", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
