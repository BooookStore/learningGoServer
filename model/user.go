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

func (u UserList) DeleteById(id int) {
	var index int
	//for ; index < len(u); index++ {
	//	if *u.ID == id {
	//		break
	//	}
	//}

	for _, user := range u {
		if user.ID == id {
			break
		}
		index++
	}

	h.UserData = append(h.UserData[:index], h.UserData[index+1:]...)
}
