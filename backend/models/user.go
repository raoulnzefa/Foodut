package models

type User struct {
	email        string
	name         string
	password     string
	profilePhoto string
	username     string
	level        int
}

// Getter Setter
func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) GetProfilePhoto() string {
	return u.profilePhoto
}

func (u *User) SetProfilePhoto(profilePhoto string) {
	u.profilePhoto = profilePhoto
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) GetLevel() int {
	return u.level
}

func (u *User) SetLevel(level int) {
	u.level = level
}
