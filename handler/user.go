package handler

import (
	resultdto "dumbmerch/dto/result"
	userdto "dumbmerch/dto/user"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) FindUser(c echo.Context) error {
	user, err := h.UserRepository.FindUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: user})

}

func (h *handler) FindUserId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := h.UserRepository.FindUserId(id)

	if user.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Data Gaada Bos"})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: user})
}
func (h *handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := h.UserRepository.FindUserId(id)

	if user.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "KOSONG OM"})
	}

	data, err := h.UserRepository.DeleteUser(id, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}
func (h *handler) CreateUser(c echo.Context) error {
	request := new(userdto.CreateUser)
	// Id, _ := strconv.Atoi(c.Param("id"))
	// user, _ := h.UserRepository.FindUserId(Id)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	user := models.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	data, err := h.UserRepository.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func convertResponse(user models.User) userdto.UserResponse {
	return userdto.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}