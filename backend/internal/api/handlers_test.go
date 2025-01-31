package api

import (
	"bytes"

	"gosolve/backend/internal/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"gosolve/backend/internal/errors"
	"gosolve/backend/internal/mocks"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestSearchHandler ...
func TestSearchHandler(t *testing.T) {
	mockController := new(mocks.MockController)
	handler := NewHandler(mockController)

	tests := []struct {
		name           string
		requestValue   string
		mockResponse   *models.SearchResponse
		mockError      error
		expectedStatus int
	}{
		{
			name:           "Valid search request",
			requestValue:   "100",
			mockResponse:   &models.SearchResponse{Index: 2, Value: 100, Message: "Closest match found"},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid search request",
			requestValue:   "invalid",
			mockResponse:   nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Internal server error",
			requestValue:   "500",
			mockResponse:   nil,
			mockError:      errors.ErrInputFileNotSet,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/search/"+tt.requestValue, bytes.NewBuffer(nil))
			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/search/{value}", handler.SearchHandler)

			if value, err := strconv.Atoi(tt.requestValue); err == nil {
				mockController.On("SearchValue", mock.Anything, mock.MatchedBy(func(req *models.SearchRequest) bool {
					return req.Target == value
				})).
					Return(tt.mockResponse, tt.mockError).
					Once()
			}

			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			mockController.AssertExpectations(t)
		})
	}
}
