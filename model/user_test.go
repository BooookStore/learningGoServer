package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	users := UserList{
		User{
			ID:   1,
			Name: "A",
			Age:  20,
		},
		User{
			ID:   2,
			Name: "B",
			Age:  20,
		},
		User{
			ID:   3,
			Name: "C",
			Age:  20,
		},
	}

	users.DeleteById(1)

	assert.Contains(t, UserList{
		{2, "B", 20},
		{3, "C", 20},
	}, users)
}
