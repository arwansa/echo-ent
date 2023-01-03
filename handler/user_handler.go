package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/arwansa/echo-ent/ent"
	"github.com/arwansa/echo-ent/usecase"
	"github.com/arwansa/echo-ent/utils"
	"github.com/labstack/echo"
)

type userHandler struct {
	e      *echo.Echo
	userUc usecase.UserUsecase
}

func NewUserHandler(e *echo.Echo, userUc usecase.UserUsecase) {
	handler := &userHandler{e: e, userUc: userUc}
	e.POST("/users", handler.Create)
	e.GET("/users/:id", handler.GetById)
	e.PUT("/users/:id", handler.UpdateById)
	e.DELETE("/users/:id", handler.DeleteById)
}

func (h *userHandler) Create(c echo.Context) error {
	defer c.Request().Body.Close()

	user := ent.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return utils.ReturnResponse(c, http.StatusBadRequest, err, nil)
	}

	result, err := h.userUc.Create(user.Name, user.Email, user.Role.String())
	if err != nil {
		return utils.ReturnResponse(c, http.StatusBadRequest, err, nil)
	}

	return utils.ReturnResponse(c, http.StatusCreated, nil, result)
}

func (h *userHandler) GetById(c echo.Context) error {
	result, err := h.userUc.GetById(c.Param("id"))
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidUserId) {
			return utils.ReturnResponse(c, http.StatusBadRequest, err, nil)
		}
		return utils.ReturnResponse(c, http.StatusNotFound, err, nil)
	}

	return utils.ReturnResponse(c, http.StatusOK, nil, result)
}

func (h *userHandler) UpdateById(c echo.Context) error {
	defer c.Request().Body.Close()

	user := ent.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return utils.ReturnResponse(c, http.StatusBadRequest, err, nil)
	}

	result, err := h.userUc.UpdateById(c.Param("id"), user.Name, user.Email, user.Role.String())
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidUserId) {
			return utils.ReturnResponse(c, http.StatusBadRequest, err, nil)
		}
		return utils.ReturnResponse(c, http.StatusNotFound, err, nil)
	}

	return utils.ReturnResponse(c, http.StatusOK, nil, result)
}

func (h *userHandler) DeleteById(c echo.Context) error {
	err := h.userUc.DeleteById(c.Param("id"))
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidUserId) {
			return utils.ReturnResponse(c, http.StatusBadRequest, err, nil)
		}
		return utils.ReturnResponse(c, http.StatusNotFound, err, nil)
	}

	return utils.ReturnResponse(c, http.StatusOK, nil, nil)
}
