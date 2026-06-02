package user

import (
	"errors"
	"net/http"
)

type UserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *UserRequest) Bind(r *http.Request) error {
	if a.Email == "" {
		return errors.New("No valid email has been submited !")
	}
	if a.Username == "" {
		return errors.New("No valid username has been submited !")
	}
	if a.Password == "" {
		return errors.New("No valid password type has benn submited !")
	}

	return nil
}

// Custom response type
type UserResponse struct {
	ID       uint   `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
