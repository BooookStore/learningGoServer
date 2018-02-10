package controller

import (
	"github.com/BooookStore/echovue/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// handler is center type
type Handler struct {
	UserData []*model.User
}

type ErrorJSON struct {
	Message string `json:"message"`
}

// GetUser return json data for user.
func (h *Handler) RetrieveUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, h.UserData)
}

// GetOneUser return json data for one user by user id.
func (h *Handler) RetrieveUser(c echo.Context) error {
	idstr := c.Param("id")
	idint, err := strconv.Atoi(idstr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorJSON{Message: "Bad Request user id " + idstr})
	}

	// retrive user by id
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
	h.UserData = append(h.UserData, u)
	return c.JSON(http.StatusCreated, u)
}
