package dto

// Login represent fields needed to log in to the app
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User represents an app user
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Points   int    `json:"points"`
}

func (u *User) IsValid() bool {
	switch {
	case u.Name == "":
		return false
	case u.Email == "":
		return false
	case u.Password == "":
		return false
	default:
		return true
	}
}
