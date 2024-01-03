package handler

import (
	"fakhry/clean-arch/features/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) GetAllUsers(c echo.Context) error {
	// panggil func di service layer
	results, errSelect := handler.userService.GetAll()
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error read data. " + errSelect.Error(),
		})
	}
	// proses mapping dari core ke response
	var usersResult []UserResponse
	for _, value := range results {
		usersResult = append(usersResult, UserResponse{
			ID:    value.ID,
			Name:  value.Name,
			Email: value.Email,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    usersResult,
	})
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data. data not valid",
		})
	}

	userCore := user.Core{
		Name:        newUser.Name,
		Email:       newUser.Email,
		Password:    newUser.Password,
		Address:     newUser.Address,
		PhoneNumber: newUser.PhoneNumber,
		Role:        newUser.Role,
	}
	errInsert := handler.userService.Create(userCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data. insert failed",
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "insert success",
	})
}
