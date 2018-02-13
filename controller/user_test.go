package controller

import (
	"encoding/json"
	"github.com/BooookStore/echovue/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// test data
var (
	mockdb = []*model.User{
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
	if assert.NoError(t, h.getUsers(c)) {
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
	if assert.NoError(t, h.getUser(c)) {
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
	if assert.NoError(t, h.getUser(c)) {
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

func TestRemoveUser(t *testing.T) {

}
