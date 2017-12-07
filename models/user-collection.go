package models

type UserCollection []User

func GetAllUsers() (uc UserCollection, err error) {
	uc = UserCollection{}

	if err != nil {
		return
	}

	uc = append(uc, User{ID: "1", Firstname: "Mirjana"})
	uc = append(uc, User{ID: "2", Firstname: "Marko"})

	return
}
