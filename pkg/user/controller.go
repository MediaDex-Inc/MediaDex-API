package user

import (
	"mediadex/config"
	"mediadex/database/dbmodel"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type UserConfig struct {
	*config.Config
}

func New(config *config.Config) *UserConfig {
	return &UserConfig{config}
}

// PostHandler godoc
// @Summary      Create a new User
// @Description  Creates a new User entry in the database
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body      UserRequest  true  "User creation payload"
// @Security     BearerAuth
// @Success      200    {object}  UserResponse
// @Failure      400    {object}  map[string]string  "Invalid User POST request payload !"
// @Failure      500    {object}  map[string]string  "Failed to create User !"
// @Router       /api/v1/user [post]
func (config *UserConfig) PostHandler(w http.ResponseWriter, r *http.Request) {

	// Get the request.
	req := &UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"Error": "Invalid User POST request payload !" + err.Error()})
		return
	}

	// Convert the requested data into dbmodel.User type for the "Create" function.
	user := &dbmodel.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	// Request the DB to Create the User.
	savedUser, err := config.UserRepository.Create(user)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to create User !" + err.Error()})
		return
	}

	// Set up to a dedicated type for the response.
	res := &UserResponse{
		Username: savedUser.Username,
		Email:    savedUser.Email,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// GetByIdHandler godoc
// @Summary      Get User by ID
// @Description  Retrieves a specific User from the database by its ID
// @Tags         User
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Security     BearerAuth
// @Success      200  {object}  UserResponse
// @Failure      404  {object}  map[string]string  "User not found"
// @Failure      500  {object}  map[string]string  "Failed to find specific User !"
// @Router       /api/v1/User/{id} [get]
func (config *UserConfig) GetByIdHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to retrieve ID !"})
		return
	}

	// Request the DB to Get the needed informations
	user, err := config.UserRepository.FindById(uint(id))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to Find specific User !" + err.Error()})
		return
	}

	// Set up to a dedicated type for the response
	res := &UserResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// GetAllHandler godoc
// @Summary      Get all user
// @Description  Retrieve all user
// @Tags         User
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   UserResponse
// @Failure      500  {object}  map[string]string
// @Router       /api/v1/user [get]
func (config *UserConfig) GetAllHandler(w http.ResponseWriter, r *http.Request) {

	user, err := config.UserRepository.Find()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "failed to fetch User !" + err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

// UpdateHandler godoc
// @Summary      Update a user
// @Description  Update an existing user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id     path     string        true  "User ID"
// @Param        user  body     UserRequest  true  "Updated user payload"
// @Security     BearerAuth
// @Success      200   {object}  UserResponse
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /api/v1/user/{id} [patch]
func (config *UserConfig) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to retrieve ID !"})
		return
	}

	// Get the request
	req := &UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request payload !" + err.Error()})
		return
	}

	// Request the DB to get the User data
	existing, err := config.UserRepository.FindById(uint(id))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "User not found !" + err.Error()})
		return
	}

	// TODO Check if the value is null if not put in the request
	// existing.Id = req.Id
	if &req.Username != nil {
		existing.Username = req.Username
	}
	if &req.Email != nil {
		existing.Email = req.Email
	}
	if &req.Password != nil {
		existing.Password = req.Password
	}

	updatedUser, err := config.UserRepository.Update(existing)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to update User !" + err.Error()})
		return
	}

	res := UserResponse{
		Email:    updatedUser.Email,
		Username: updatedUser.Username,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// DeleteHandler godoc
// @Summary      Delete a user
// @Description  Deletes a user from the database by its ID
// @Tags         User
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Security     BearerAuth
// @Success      200  {object}  map[string]string  "User deleted successfully"
// @Failure      404  {object}  map[string]string  "User not found !"
// @Failure      500  {object}  map[string]string  "Failed to delete User !"
// @Router       /api/v1/user/{id} [delete]
func (config *UserConfig) DeleteHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to retrieve ID !"})
		return
	}

	// Request the DB to Delete the informations
	err = config.UserRepository.Delete(uint(id))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to Delete User !" + err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"message": "User deleted successfully."})
}
