package mediaField

import (
	"mediadex/config"
	"mediadex/database/dbmodel"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type MediaFieldConfig struct {
	*config.Config
}

func New(config *config.Config) *MediaFieldConfig {
	return &MediaFieldConfig{config}
}

// PostHandler godoc
// @Summary      Create a new MediaField
// @Description  Creates a new MediaField entry in the database
// @Tags         MediaField
// @Accept       json
// @Produce      json
// @Param        mediaField  body      MediaFieldRequest  true  "MediaField creation payload"
// @Security     BearerAuth
// @Success      200    {object}  MediaFieldResponse
// @Failure      400    {object}  map[string]string  "Invalid MediaField POST request payload !"
// @Failure      500    {object}  map[string]string  "Failed to create MediaField !"
// @Router       /mediaField [post]
func (config *MediaFieldConfig) PostHandler(w http.ResponseWriter, r *http.Request) {

	// Get the request.
	req := &MediaFieldRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"Error": "Invalid MediaField POST request payload !" + err.Error()})
		return
	}

	// Convert the requested data into dbmodel.MediaField type for the "Create" function.
	mediaField := &dbmodel.MediaField{
		FieldID: req.FieldID,
		MediaID: req.MediaID,
		Value:   req.Value}

	// Request the DB to Create the MediaField.
	savedMediaField, err := config.MediaFieldRepository.Create(mediaField)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to create MediaField !" + err.Error()})
		return
	}

	// Set up to a dedicated type for the response.
	res := &MediaFieldResponse{
		FieldID: savedMediaField.FieldID,
		MediaID: savedMediaField.MediaID,
		Value:   savedMediaField.Value}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// GetByIdHandler godoc
// @Summary      Get mediaField by ID
// @Description  Retrieves a specific mediaField from the database by its ID
// @Tags         MediaField
// @Produce      json
// @Param        id   path      string  true  "mediaField ID"
// @Security     BearerAuth
// @Success      200  {object}  MediaFieldResponse
// @Failure      404  {object}  map[string]string  "MediaField not found"
// @Failure      500  {object}  map[string]string  "Failed to find specific MediaField !"
// @Router       /mediaField/{id} [get]
func (config *MediaFieldConfig) GetByIdHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	fieldID, err := strconv.Atoi(chi.URLParam(r, "fieldID"))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to retrieve fieldID !"})
		return
	}
	mediaID, err := strconv.Atoi(chi.URLParam(r, "mediaID"))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to retrieve mediaID !"})
		return
	}

	// Request the DB to Get the needed informations
	mediaField, err := config.MediaFieldRepository.FindById(uint(fieldID), uint(mediaID))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to Find specific MediaField !" + err.Error()})
		return
	}

	// Set up to a dedicated type for the response
	res := &MediaFieldResponse{
		FieldID: mediaField.FieldID,
		MediaID: mediaField.MediaID,
		Value:   mediaField.Value}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// GetAllHandler godoc
// @Summary      Get all mediaField
// @Description  Retrieve all mediaField
// @Tags         MediaField
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   MediaFieldResponse
// @Failure      500  {object}  map[string]string
// @Router       /mediaField [get]
func (config *MediaFieldConfig) GetAllHandler(w http.ResponseWriter, r *http.Request) {

	mediaField, err := config.MediaFieldRepository.FindAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "failed to fetch all MediaField !" + err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, mediaField)
}

// UpdateHandler godoc
// @Summary      Update a mediaField
// @Description  Update an existing mediaField
// @Tags         MediaField
// @Accept       json
// @Produce      json
// @Param        id     path     string        true  "MediaField ID"
// @Param        mediaField  body     MediaFieldRequest  true  "Updated mediaField payload"
// @Security     BearerAuth
// @Success      200   {object}  MediaFieldResponse
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /mediaField/{id} [patch]
func (config *MediaFieldConfig) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	fieldID, err := strconv.Atoi(chi.URLParam(r, "fieldID"))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to retrieve fieldID !"})
		return
	}
	mediaID, err := strconv.Atoi(chi.URLParam(r, "mediaID"))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to retrieve mediaID !"})
		return
	}

	// Get the request
	req := &MediaFieldRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request payload !" + err.Error()})
		return
	}

	// Request the DB to get the MediaField data
	existing, err := config.MediaFieldRepository.FindById(uint(fieldID), uint(mediaID))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "MediaField not found !" + err.Error()})
		return
	}

	if req.FieldID <= 0 {
		existing.FieldID = req.FieldID
	}
	if req.MediaID <= 0 {
		existing.MediaID = req.MediaID
	}
	if req.Value == "" {
		existing.Value = req.Value
	}

	updatedMediaField, err := config.MediaFieldRepository.Update(existing)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to update MediaField !" + err.Error()})
		return
	}

	res := MediaFieldResponse{
		FieldID: updatedMediaField.FieldID,
		MediaID: updatedMediaField.MediaID,
		Value:   updatedMediaField.Value,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// DeleteHandler godoc
// @Summary      Delete a mediaField
// @Description  Deletes a mediaField from the database by its ID
// @Tags         MediaField
// @Produce      json
// @Param        id   path      string  true  "MediaField ID"
// @Security     BearerAuth
// @Success      200  {object}  map[string]string  "MediaField deleted successfully"
// @Failure      404  {object}  map[string]string  "MediaField not found !"
// @Failure      500  {object}  map[string]string  "Failed to delete MediaField !"
// @Router       /mediaField/{id} [delete]
func (config *MediaFieldConfig) DeleteHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to retrieve ID !"})
		return
	}

	// Request the DB to Delete the informations
	err = config.MediaFieldRepository.Delete(uint(id))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to Delete MediaField !" + err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"message": "MediaField deleted successfully."})
}
