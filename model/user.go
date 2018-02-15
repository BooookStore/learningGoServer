// Define model type.
package model

type (
	User struct {
		ID   int    `json:"id"`
		Name string `json:"Name"`
		Age  int    `json:"Age"`
	}

	Company struct {
		Name          string
		AddressNumber string
		Address       string
	}

	UserList []User
)
