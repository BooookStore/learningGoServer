package controller

import (
	"github.com/BooookStore/learningGoServer/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// handler is center type
type Handler struct {
	UserData model.UserList
}

// ErrorJSON is error type caused json.
type ErrorJSON struct {
	Message string `json:"message"`
}

// GetUser return json data for user.
func (h *Handler) GetAllUser(c echo.Context) error {
	return c.JSON(http.StatusOK, h.UserData)
}

// GetOneUser return json data for one user by user id.
func (h *Handler) GetUser(c echo.Context) error {
	idstr := c.Param("id")
	idint, err := strconv.Atoi(idstr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorJSON{Message: "Bad Request user id " + idstr})
	}

	// retrieve user by id
	for _, user := range h.UserData {
		if user.ID == idint {
			return c.JSON(http.StatusOK, user)
		}
	}

	return c.JSON(http.StatusNotFound, ErrorJSON{Message: "Not Found user id " + strconv.Itoa(idint)})
}

// CreateUser create new user.
func (h *Handler) CreateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	h.UserData = append(h.UserData, *u)
	return c.JSON(http.StatusCreated, u)
}

// DeleteUser delete user.
func (h *Handler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorJSON{Message: "Bad Request user id " + c.Param("id")})
	}

	err = h.UserData.DeleteById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorJSON{Message: "Not found user by id " + c.Param("id")})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	h.UserData.Update(id, u)
	return c.JSON(http.StatusOK, u)
}
