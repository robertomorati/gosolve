package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"gosolve/backend/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAbs ...
func TestAbs(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Positive number", 5, 5},
		{"Negative number", -5, 5},
		{"Zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Abs(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestEncodeResponse ...
func TestEncodeResponse(t *testing.T) {
	tests := []struct {
		name         string
		acceptHeader string
		response     *models.SearchResponse
		wantErr      error
		wantStatus   int
	}{
		{
			name:         "Valid JSON response",
			acceptHeader: "application/json",
			response:     &models.SearchResponse{},
			wantErr:      nil,
			wantStatus:   http.StatusOK,
		},
		{
			name:         "Unsupported content type",
			acceptHeader: "text/html",
			response:     &models.SearchResponse{},
			wantErr:      errors.New("unsupported content type: text/html"),
			wantStatus:   http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", bytes.NewBuffer(nil))
			req.Header.Set("Accept", tt.acceptHeader)

			rr := httptest.NewRecorder()
			err := EncodeResponse(rr, req, tt.response)

			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)

				var got models.SearchResponse
				err := json.NewDecoder(rr.Body).Decode(&got)
				assert.NoError(t, err)
				assert.Equal(t, *tt.response, got)
			}
		})
	}
}
