package field

import (
	"mediadex/config"
	"mediadex/database/dbmodel"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type FieldConfig struct {
	*config.Config
}

func New(config *config.Config) *FieldConfig {
	return &FieldConfig{config}
}

// PostHandler godoc
// @Summary      Create a new Field
// @Description  Creates a new Field entry in the database
// @Tags         Field
// @Accept       json
// @Produce      json
// @Param        field  body      FieldRequest  true  "Field creation payload"
// @Security     BearerAuth
// @Success      200    {object}  FieldResponse
// @Failure      400    {object}  map[string]string  "Invalid Field POST request payload !"
// @Failure      500    {object}  map[string]string  "Failed to create Field !"
// @Router       /fields [post]
func (config *FieldConfig) PostHandler(w http.ResponseWriter, r *http.Request) {

	// Get the request.
	req := &FieldRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"Error": "Invalid Field POST request payload !" + err.Error()})
		return
	}

	// Convert the requested data into dbmodel.Field type for the "Create" function.
	field := &dbmodel.Field{
		UserId: req.UserId,
		Name:   req.Name}

	// Request the DB to Create the Field.
	savedField, err := config.FieldRepository.Create(field)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to create Field !" + err.Error()})
		return
	}

	// Set up to a dedicated type for the response.
	res := &FieldResponse{
		UserId: savedField.UserId,
		Name:   savedField.Name}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// GetByIdHandler godoc
// @Summary      Get field by ID
// @Description  Retrieves a specific field from the database by its ID
// @Tags         Field
// @Produce      json
// @Param        id   path      string  true  "field ID"
// @Security     BearerAuth
// @Success      200  {object}  FieldResponse
// @Failure      404  {object}  map[string]string  "Field not found"
// @Failure      500  {object}  map[string]string  "Failed to find specific Field !"
// @Router       /fields/{id} [get]
func (config *FieldConfig) GetByIdHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to retrieve ID !"})
		return
	}

	// Request the DB to Get the needed informations
	field, err := config.FieldRepository.FindById(uint(id))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to Find specific Field !" + err.Error()})
		return
	}

	// Set up to a dedicated type for the response
	res := &FieldResponse{
		UserId: field.UserId,
		Name:   field.Name}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// GetAllHandler godoc
// @Summary      Get all field
// @Description  Retrieve all field
// @Tags         Field
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   FieldResponse
// @Failure      500  {object}  map[string]string
// @Router       /fields [get]
func (config *FieldConfig) GetAllHandler(w http.ResponseWriter, r *http.Request) {

	fields, err := config.FieldRepository.FindAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "failed to fetch Field !" + err.Error()})
		return
	}

	var result []FieldResponse

	for _, field := range fields {
		res := FieldResponse{
			UserId: field.UserId,
			Name:   field.Name,
		}
		result = append(result, res)
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}

// UpdateHandler godoc
// @Summary      Update a field
// @Description  Update an existing field
// @Tags         Field
// @Accept       json
// @Produce      json
// @Param        id     path     string        true  "Field ID"
// @Param        field  body     FieldRequest  true  "Updated field payload"
// @Security     BearerAuth
// @Success      200   {object}  FieldResponse
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /fields/{id} [patch]
func (config *FieldConfig) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to retrieve ID !"})
		return
	}

	// Get the request
	req := &FieldRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request payload !" + err.Error()})
		return
	}

	// Request the DB to get the Field data
	existing, err := config.FieldRepository.FindById(uint(id))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Field not found !" + err.Error()})
		return
	}

	if req.UserId > 0 {
		existing.UserId = req.UserId
	}
	if req.Name == "" {
		existing.Name = req.Name
	}

	updatedField, err := config.FieldRepository.Update(existing)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to update Field !" + err.Error()})
		return
	}

	res := FieldResponse{
		UserId: updatedField.UserId,
		Name:   updatedField.Name,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// DeleteHandler godoc
// @Summary      Delete a field
// @Description  Deletes a field from the database by its ID
// @Tags         Field
// @Produce      json
// @Param        id   path      string  true  "Field ID"
// @Security     BearerAuth
// @Success      200  {object}  map[string]string  "Field deleted successfully"
// @Failure      404  {object}  map[string]string  "Field not found !"
// @Failure      500  {object}  map[string]string  "Failed to delete Field !"
// @Router       /fields/{id} [delete]
func (config *FieldConfig) DeleteHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to retrieve ID !"})
		return
	}

	// Request the DB to Delete the informations
	err = config.FieldRepository.Delete(uint(id))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to Delete Field !" + err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"message": "Field deleted successfully."})
}
