package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strconv"
	"errors"
)

func TestDeleteUser(t *testing.T) {
	// Setup
	users := UserList{{1, "A", 20,}, {2, "B", 20,}, {3, "C", 20,}, {4, "D", 20,}}

	// Action & Verify
	users.DeleteById(2)
	assert.Equal(t, UserList{{1, "A", 20}, {3, "C", 20}, {4, "D", 20,}}, users)
	users.DeleteById(1)
	assert.Equal(t, UserList{{3, "C", 20}, {4, "D", 20,}}, users)

	// Verify error and error message
	if assert.Error(t, users.DeleteById(2)) {
		assert.Equal(t, users.DeleteById(2), errors.New("Not found user id "+strconv.Itoa(2)))
	}
}

func TestUpdateUser(t *testing.T) {
	// Setup
	users := UserList{{1, "A", 20,}, {2, "B", 20,}, {3, "C", 20,}, {4, "D", 20,}}

	// Action & Verify
	users.Update(1, &User{1, "AA", 200})
	assert.Equal(t, UserList{{1, "AA", 200,}, {2, "B", 20,}, {3, "C", 20,}, {4, "D", 20,}}, users)
	users.Update(2, &User{2, "BBB", 300})
	assert.Equal(t, UserList{{1, "AA", 200,}, {2, "BBB", 300,}, {3, "C", 20,}, {4, "D", 20,}}, users)
}
