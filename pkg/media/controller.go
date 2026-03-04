package media

import (
	"MediaDex/database/dbmodel"
	"mediadex/mediadex-API/config"
	"mediadex/mediadex-API/database/dbmodel"
	"mediadex/mediadex-API/pkg/model"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type MediaConfig struct {
	*config.Config
}

func New(config *config.Config) *MediaConfig {
	return &MediaConfig{config}
}

// PostHandler godoc
// @Summary      Create a new Media
// @Description  Creates a new Media entry in the database
// @Tags         Media
// @Accept       json
// @Produce      json
// @Param        media  body      model.MediaRequest  true  "Media creation payload"
// @Security     BearerAuth
// @Success      200    {object}  model.MediaResponse
// @Failure      400    {object}  map[string]string  "Invalid Media POST request payload !"
// @Failure      500    {object}  map[string]string  "Failed to create Media !"
// @Router       /api/v1/media [post]
func (config *MediaConfig) PostHandler(w http.ResponseWriter, r *http.Request) {

	// Get the request.
	req := &model.MediaRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"Error": "Invalid Media POST request payload !" + err.Error()})
		return
	}

	// Convert the requested data into dbmodel.Media type for the "Create" function.
	media := &dbmodel.Media{
		userId:    *req.userId,
		name:      *req.name,
		status:    *req.status,
		mediaType: *req.mediaType,
		imgURL:    *req.imgURL,
		rating:    *req.rating,
		notes:     *req.notes}

	// Request the DB to Create the Media.
	savedMedia, err := config.MediaRepository.Create(media)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to create Media !" + err.Error()})
		return
	}

	// Set up to a dedicated type for the response.
	res := &model.MediaResponse{
		userId:    savedMedia.userId,
		name:      savedMedia.name,
		status:    savedMedia.status,
		mediaType: savedMedia.mediaType,
		imgURL:    savedMedia.imgURL,
		rating:    savedMedia.rating,
		notes:     savedMedia.notes}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// GetByIdHandler godoc
// @Summary      Get media by ID
// @Description  Retrieves a specific media from the database by its ID
// @Tags         Media
// @Produce      json
// @Param        id   path      string  true  "media ID"
// @Security     BearerAuth
// @Success      200  {object}  model.mediaResponse
// @Failure      404  {object}  map[string]string  "Media not found"
// @Failure      500  {object}  map[string]string  "Failed to find specific Media !"
// @Router       /api/v1/media/{id} [get]
func (config *MediaConfig) GetByIdHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Missing Id !"})
		return
	}

	uuid, err := uuid.Parse(idStr)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Invalid Id !" + err.Error()})
		return
	}

	// Request the DB to Get the needed informations
	media, err := config.MediaRepository.FindById(uuid)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Failed to Find specific Media !" + err.Error()})
		return
	}

	// Set up to a dedicated type for the response
	res := &model.MediaResponse{
		userId:    media.userId,
		name:      media.name,
		status:    media.status,
		mediaType: media.mediaType,
		imgURL:    media.imgURL,
		rating:    media.rating,
		notes:     media.notes}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// GetAllHandler godoc
// @Summary      Get all media
// @Description  Retrieve all media
// @Tags         Media
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   model.MediaResponse
// @Failure      500  {object}  map[string]string
// @Router       /api/v1/media [get]
func (config *MediaConfig) GetAllHandler(w http.ResponseWriter, r *http.Request) {

	media, err := config.MediaRepository.Find()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "failed to fetch Media !" + err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, media)
}

// UpdateHandler godoc
// @Summary      Update a media
// @Description  Update an existing media
// @Tags         Media
// @Accept       json
// @Produce      json
// @Param        id    path      string               true  "Media ID"
// @Param        media  body      model.mediaRequest   true  "Updated media payload"
// @Security     BearerAuth
// @Success      200   {object}  model.mediaResponse
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /api/v1/media/{id} [patch]
func (config *MediaConfig) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Missing Id !"})
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Invalid Id !" + err.Error()})
		return
	}

	// Get the request
	req := &model.MediaRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request payload !" + err.Error()})
		return
	}

	// Request the DB to get the Media data
	existing, err := config.MediaRepository.FindById(id)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Media not found !" + err.Error()})
		return
	}

	// TODO Check if the value is null if not put in the request
	existing.userId = req.userId
	existing.name = req.name
	existing.status = req.status
	existing.mediaType = req.mediaType
	existing.imgURL = req.imgURL
	existing.rating = req.rating
	existing.notes = req.notes

	updatedMedia, err := config.MediaRepository.Update(existing)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to update Media !" + err.Error()})
		return
	}

	res := model.MediaResponse{
		userId:    updatedMedia.userId,
		name:      updatedMedia.name,
		status:    updatedMedia.status,
		mediaType: updatedMedia.mediaType,
		imgURL:    updatedMedia.imgURL,
		rating:    updatedMedia.rating,
		notes:     updatedMedia.notes,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// DeleteHandler godoc
// @Summary      Delete a media
// @Description  Deletes a media from the database by its ID
// @Tags         Media
// @Produce      json
// @Param        id   path      string  true  "Media ID"
// @Security     BearerAuth
// @Success      200  {object}  map[string]string  "Media deleted successfully"
// @Failure      404  {object}  map[string]string  "Media not found !"
// @Failure      500  {object}  map[string]string  "Failed to delete Media !"
// @Router       /api/v1/media/{id} [delete]
func (config *MediaConfig) DeleteHandler(w http.ResponseWriter, r *http.Request) {

	// Get the id in the URL
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"Error": "Missing Id !"})
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Invalid Id !" + err.Error()})
		return
	}

	// Request the DB to Delete the informations
	errDelete := config.MediaRepository.Delete(id)
	if errDelete != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"Error": "Failed to Delete Media !" + err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"message": "Media deleted successfully."})
}
