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

// Update is update exist user data. Update target user found by user id.
func (u *UserList) Update(id int, newUser *User) error {
	updateIndex, isFound := u.searchById(id)
	if !isFound {
		return errors.New("Not found user id " + strconv.Itoa(id))
	}

	// delete old data and insert new data
	head := append((*u)[:updateIndex], *newUser)
	*u = append(head, (*u)[updateIndex+1:]...)
	return nil
}

// DeleteById delete user by user id.
func (u *UserList) DeleteById(id int) error {
	deleteIndex, isFound := u.searchById(id)
	if !isFound {
		return errors.New("Not found user id " + strconv.Itoa(id))
	}

	*u = append((*u)[:deleteIndex], (*u)[deleteIndex+1:]...)
	return nil
}

// SearchById is search user by user id.
// If user not found, return index -1 and found false.
func (u *UserList) searchById(id int) (index int, found bool) {
	index = 0

	for _, user := range *u {
		if user.ID == id {
			return index, true
		}
		index++
	}

	return -1, false
}
