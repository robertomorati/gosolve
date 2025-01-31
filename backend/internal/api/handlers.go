package api

import (
	"context"
	"encoding/json"
	"gosolve/backend/internal/errors"
	"net/http"

	"strconv"

	"gosolve/backend/internal/controller"

	"gosolve/backend/internal/models"

	"github.com/gorilla/mux"
)

// Handler ...
type Handler struct {
	controller controller.Controller
}

// NewHandler ...
func NewHandler(ctrl controller.Controller) *Handler {
	return &Handler{controller: ctrl}
}

// @Summary Search for the closest value
// @Description Search for the closest value
// @Tags Search
// @Accept json
// @Produce json
// @Param value path int true "Value to search"
// @Success 200 {object} models.SearchResponse
// @Failure 400 {object} errors.APIError
// @Failure 404 {object} errors.APIError
// @Failure 500 {object} errors.APIError
// @Router /search/{value} [get]
// @
func (h *Handler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value, err := strconv.Atoi(vars["value"])
	if err != nil {
		w.WriteHeader(errors.ErrInvalidRequestValue.StatusCode)
		return
	}

	ctx := context.Background()
	response, err := h.controller.SearchValue(
		ctx,
		&models.SearchRequest{
			Target: value,
		},
	)

	if err != nil {
		if apiErr, ok := err.(*errors.APIError); ok {
			w.WriteHeader(apiErr.StatusCode)
			response, _ = handleAPIError(w, r, apiErr)
			json.NewEncoder(w).Encode(response)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			response, _ = handleAPIError(w, r, apiErr)
			json.NewEncoder(w).Encode(response)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// handleAPIError writes the error response to the client
func handleAPIError(w http.ResponseWriter, r *http.Request, apiErr *errors.APIError) (*models.SearchResponse, error) {
	if apiErr == nil {
		apiErr = errors.ErrInternalServer
	}
	return &models.SearchResponse{
		Index:   -1,
		Message: apiErr.Message,
		Value:   -1,
	}, nil
}
