package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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
		return utils.ReturnResponse(c, getStatusCode(err), err, result)
	}

	return utils.ReturnResponse(c, http.StatusCreated, nil, result)
}

func (h *userHandler) GetById(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ReturnResponse(c, http.StatusNotFound, usecase.ErrNotFound, nil)

	}

	result, err := h.userUc.GetById(userId)
	if err != nil {
		return utils.ReturnResponse(c, getStatusCode(err), err, nil)
	}

	return utils.ReturnResponse(c, http.StatusOK, nil, result)
}

func (h *userHandler) UpdateById(c echo.Context) error {
	defer c.Request().Body.Close()
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ReturnResponse(c, http.StatusNotFound, usecase.ErrNotFound, nil)

	}

	user := ent.User{}
	err = json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return utils.ReturnResponse(c, http.StatusBadRequest, err, nil)
	}

	result, err := h.userUc.UpdateById(userId, user.Name, user.Email, user.Role.String())
	if err != nil {
		return utils.ReturnResponse(c, getStatusCode(err), err, nil)
	}

	return utils.ReturnResponse(c, http.StatusOK, nil, result)
}

func (h *userHandler) DeleteById(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ReturnResponse(c, http.StatusNotFound, usecase.ErrNotFound, nil)

	}

	err = h.userUc.DeleteById(userId)
	if err != nil {
		return utils.ReturnResponse(c, getStatusCode(err), err, nil)
	}

	return utils.ReturnResponse(c, http.StatusOK, nil, nil)
}

func getStatusCode(err error) int {
	switch {
	case err == nil:
		return http.StatusOK
	case ent.IsValidationError(err):
		return http.StatusUnprocessableEntity
	case ent.IsNotFound(err):
		return http.StatusNotFound
	case ent.IsConstraintError(err):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
