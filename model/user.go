// Define model type.
package model

import (
	"errors"
	"strconv"
)

type (
	User struct {
		ID   int    `json:"id"`
		Name string `json:"Name"`
		Age  int    `json:"Age"`
	}

	UserList []User
)

// DeleteById delete user by user id.
func (u *UserList) DeleteById(id int) error {
	var (
		dti     = 0
		isFound = false
	)

	for _, user := range *u {
		if user.ID == id {
			isFound = true
			break
		}
		dti++
	}

	if !isFound {
		return errors.New("Not found user id " + strconv.Itoa(id))
	} else {
		*u = append((*u)[:dti], (*u)[dti+1:]...)
		return nil
	}
}
