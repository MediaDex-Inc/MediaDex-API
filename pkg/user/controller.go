package user

import (
	"encoding/json"
	"mediadex/config"
	"mediadex/pkg/authentication"
	"mediadex/pkg/model"
	"net/http"

	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type UserConfig struct {
	*config.Config
}

func New(config *config.Config) *UserConfig {
	return &UserConfig{config}
}

// GetMeHandler godoc
// @Summary      Get current user
// @Description  Retrieves the authenticated user's profile
// @Tags         User
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  model.UserResponse
// @Failure      404  {object}  map[string]string
// @Router       /users/me [get]
func (config *UserConfig) GetMeHandler(w http.ResponseWriter, r *http.Request) {
	userID := authentication.GetUserFromContext(r.Context())

	user, err := config.UserRepository.FindById(userID)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": "user not found"})
		return
	}

	res := &model.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// UpdateHandler godoc
// @Summary      Update current user
// @Description  Update the authenticated user's profile
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body     model.UserUpdateRequest  true  "Updated user payload"
// @Security     BearerAuth
// @Success      200   {object}  model.UserResponse
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users/me [patch]
func (config *UserConfig) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	userID := authentication.GetUserFromContext(r.Context())

	req := &model.UserUpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "invalid user payload: " + err.Error()})
		return
	}

	existing, err := config.UserRepository.FindById(userID)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": "user not found"})
		return
	}

	if req.Username != nil && *req.Username != "" {
		existing.Username = *req.Username
	}
	if req.Email != nil && *req.Email != "" {
		existing.Email = *req.Email
	}
	if req.Password != nil && *req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": "failed to hash password"})
			return
		}
		existing.Password = string(hashedPassword)
	}

	updatedUser, err := config.UserRepository.Update(existing)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "failed to update user: " + err.Error()})
		return
	}

	res := model.UserResponse{
		ID:       updatedUser.ID,
		Email:    updatedUser.Email,
		Username: updatedUser.Username,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// DeleteHandler godoc
// @Summary      Delete current user
// @Description  Deletes the authenticated user's account
// @Tags         User
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]string  "User deleted successfully"
// @Failure      500  {object}  map[string]string
// @Router       /users/me [delete]
func (config *UserConfig) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	userID := authentication.GetUserFromContext(r.Context())

	err := config.UserRepository.Delete(userID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "failed to delete user: " + err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"message": "User deleted successfully."})
}
