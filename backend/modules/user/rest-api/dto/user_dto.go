package dto

import "encoding/json"

type User struct {
	Name         string `form:"name" json:"name"`
	Username     string `form:"username" json:"username"`
	Email        string `form:"email" json:"email"`
	Password     string `form:"password" json:"password"`
	ProfilePhoto string `form:"profilePhoto" json:"profilePhoto"`
	Level        int    `form:"level" json:"level"`
}

type PostUser struct {
	Name     string `form:"name" json:"name"`
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type LoginUser struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

// data, _ := json.Marshal(user)
func (u User) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Email        string
		Name         string
		Password     string
		ProfilePhoto string
		Username     string
		Level        int
	}{
		Email:        u.Email,
		Name:         u.Name,
		Password:     u.Password,
		ProfilePhoto: u.ProfilePhoto,
		Username:     u.Username,
		Level:        u.Level,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

//-------------------------------------------------
// Getter Setter User
func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) SetUsername(username string) {
	u.Username = username
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) GetProfilePhoto() string {
	return u.ProfilePhoto
}

func (u *User) SetProfilePhoto(profilePhoto string) {
	u.ProfilePhoto = profilePhoto
}

func (u *User) GetLevel() int {
	return u.Level
}

func (u *User) SetLevel(level int) {
	u.Level = level
}

//-------------------------------------------------
// Getter Setter PostUser
func (u *PostUser) GetName() string {
	return u.Name
}

func (u *PostUser) SetName(name string) {
	u.Name = name
}

func (u *PostUser) GetUsername() string {
	return u.Username
}

func (u *PostUser) SetUsername(username string) {
	u.Username = username
}

func (u *PostUser) GetEmail() string {
	return u.Email
}

func (u *PostUser) SetEmail(email string) {
	u.Email = email
}

func (u *PostUser) GetPassword() string {
	return u.Password
}

func (u *PostUser) SetPassword(password string) {
	u.Password = password
}
