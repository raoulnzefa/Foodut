package dto

type Admin struct {
	User User `form:"user" json:"user"`
}

// Getter Setter
func (a *Admin) GetUser() User {
	return a.User
}

func (a *Admin) SetUser(user User) {
	a.User = user
}
