package user

import (
	"errors"
	"net/http"
)

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *UserRequest) Bind(r *http.Request) error {
	if &a.Username == nil || a.Username == "" {
		return errors.New("No valid username has been submited !")
	}
	if &a.Email == nil || a.Email == "" {
		return errors.New("No valid email has been submited !")
	}
	if &a.Password == nil || a.Password == "" {
		return errors.New("No valid password type has benn submited !")
	}

	return nil
}

// Custom response type
type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
