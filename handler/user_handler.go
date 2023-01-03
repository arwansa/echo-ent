package handler

import (
	"encoding/json"
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

	newUser := ent.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&newUser)
	if err != nil {
		return utils.ErrorResponse(c, false, http.StatusBadRequest, err)
	}

	result, err := h.userUc.Create(newUser.Name, newUser.Email, newUser.Role.String())
	if err != nil {
		return utils.ErrorResponse(c, false, http.StatusBadRequest, err)
	}

	return utils.SuccessResponse(c, true, http.StatusCreated, result)
}

func (h *userHandler) GetById(c echo.Context) error {
	result, err := h.userUc.GetById(c.Param("id"))
	if err != nil {
		return utils.ErrorResponse(c, false, http.StatusNotFound, err)
	}

	return utils.SuccessResponse(c, true, http.StatusOK, result)
}

func (h *userHandler) UpdateById(c echo.Context) error {
	defer c.Request().Body.Close()

	newUser := ent.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&newUser)
	if err != nil {
		return utils.ErrorResponse(c, false, http.StatusBadRequest, err)
	}

	result, err := h.userUc.UpdateById(c.Param("id"), newUser.Name, newUser.Email, newUser.Role.String())
	if err != nil {
		return utils.ErrorResponse(c, false, http.StatusBadRequest, err)
	}

	return utils.SuccessResponse(c, true, http.StatusOK, result)
}

func (h *userHandler) DeleteById(c echo.Context) error {
	err := h.userUc.DeleteById(c.Param("id"))
	if err != nil {
		return utils.ErrorResponse(c, false, http.StatusBadRequest, err)
	}

	return utils.SuccessResponse(c, true, http.StatusOK, nil)
}
