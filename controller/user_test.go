package controller

import (
	"encoding/json"
	"github.com/BooookStore/learningGoServer/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// test data
var (
	mockdb = model.UserList{
		{ID: 1, Name: "bookstore", Age: 24},
		{ID: 2, Name: "ryosuke", Age: 25},
		{ID: 3, Name: "yuki", Age: 26},
	}
)

func TestRetrieveUsers(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := Handler{mockdb}

	// Verify
	if assert.NoError(t, h.GetAllUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// JSON Response check.
		var actual []model.User
		json.Unmarshal(rec.Body.Bytes(), &actual)

		expected := []model.User{
			{ID: 1, Name: "bookstore", Age: 24},
			{ID: 2, Name: "ryosuke", Age: 25},
			{ID: 3, Name: "yuki", Age: 26},
		}

		assert.Equal(t, expected, actual)
	}
}

func TestRetrieveOneUserById1(t *testing.T) {
	// Setup Server
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	h := Handler{mockdb}

	// Verify
	if assert.NoError(t, h.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// JSON Response check.
		var actual model.User
		json.Unmarshal(rec.Body.Bytes(), &actual)

		expected := model.User{ID: 1, Name: "bookstore", Age: 24}

		assert.Equal(t, expected, actual)
	}
}

func TestRetrieveOneUserById2(t *testing.T) {
	// Setup Server
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	h := Handler{mockdb}

	// Verify
	if assert.NoError(t, h.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// JSON Response check.
		var actual model.User
		json.Unmarshal(rec.Body.Bytes(), &actual)

		expected := model.User{ID: 2, Name: "ryosuke", Age: 25}

		assert.Equal(t, expected, actual)
	}
}

func TestCreateUser(t *testing.T) {
	// Post JSON Data
	expected := model.User{ID: 4, Name: "John", Age: 45}
	postJSON, _ := json.Marshal(expected)

	// Setup Server
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(postJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := Handler{mockdb}

	// Verify
	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		// JSON Responce check
		var actual model.User
		json.Unmarshal(rec.Body.Bytes(), &actual)
		expected := expected

		assert.Equal(t, expected, actual)

		// mockdb check
		assert.Equal(t, 4, len(h.UserData))
	}

}

func TestDeleteUser(t *testing.T) {
	// Setup Server
	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	h := Handler{mockdb}

	if assert.NoError(t, h.DeleteUser(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Equal(t, 2, len(h.UserData))
		assert.Equal(t, model.UserList{{1, "bookstore", 24}, {3, "yuki", 26}}, h.UserData)
	}
}

func TestUpdateUser(t *testing.T) {
	// Setup update target user
	deleteTargetUser := model.User{Name: "new bookstore", Age: 35}
	postJSONBytes, _ := json.Marshal(deleteTargetUser)

	// Setup Server
	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/", strings.NewReader(string(postJSONBytes)))
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	h := Handler{mockdb}

	if assert.NoError(t, h.UpdateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, 3, len(h.UserData))
		assert.Equal(t, model.UserList{{1, "new bookstore", 35}, {2, "ryosuke", 25}, {3, "yuki", 26}}, h.UserData)

	}
}
