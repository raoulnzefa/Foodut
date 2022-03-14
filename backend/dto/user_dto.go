package entities

import "encoding/json"

type User struct {
	Name         string `form:"name" json:"name"`
	Username     string `form:"username" json:"username"`
	Email        string `form:"email" json:"email"`
	Password     string `form:"password" json:"password"`
	ProfilePhoto string `form:"profilePhoto" json:"profilePhoto"`
	Level        int    `form:"level" json:"level"`
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

// Getter Setter
func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
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

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) SetUsername(username string) {
	u.Username = username
}

func (u *User) GetLevel() int {
	return u.Level
}

func (u *User) SetLevel(level int) {
	u.Level = level
}
