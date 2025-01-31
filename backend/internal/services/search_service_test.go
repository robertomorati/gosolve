package services

import (
	"testing"

	"gosolve/backend/internal/errors"
	"gosolve/backend/internal/mocks"
	"gosolve/backend/internal/models"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// TestFindClosest ...
func TestFindClosest(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	mockRepo := new(mocks.MockRepository)
	service := NewSearchService(mockRepo, logger)

	tests := []struct {
		name          string
		target        int
		margin        float64
		mockIndex     int
		mockValue     int
		mockFindErr   error
		mockGetErr    error
		wantErr       error
		wantResp      *models.SearchResponse
		expectGetData bool
	}{
		{
			name:          "Exact match found",
			target:        10,
			margin:        0.1,
			mockIndex:     1,
			mockValue:     10,
			mockFindErr:   nil,
			mockGetErr:    nil,
			wantErr:       nil,
			wantResp:      &models.SearchResponse{Index: 1, Value: 10, Message: "Closest match found"},
			expectGetData: true,
		},
		{
			name:          "Closest match found within margin",
			target:        15,
			margin:        0.1,
			mockIndex:     1,
			mockValue:     10,
			mockFindErr:   nil,
			mockGetErr:    nil,
			wantErr:       nil,
			wantResp:      &models.SearchResponse{Index: 1, Value: 10, Message: "Closest match found"},
			expectGetData: true,
		},
		{
			name:          "No closest match found (-1 index)",
			target:        999,
			margin:        0.1,
			mockIndex:     -1,
			mockValue:     -1,
			mockFindErr:   errors.ErrNoClosestMatchFound,
			mockGetErr:    nil,
			wantErr:       errors.ErrNoClosestMatchFound,
			wantResp:      nil,
			expectGetData: false,
		},
		{
			name:          "Invalid index from repository",
			target:        200,
			margin:        0.1,
			mockIndex:     5,
			mockValue:     -1,
			mockFindErr:   nil,
			mockGetErr:    errors.ErrInvalidIndexFromRepo,
			wantErr:       errors.ErrInvalidIndexFromRepo,
			wantResp:      nil,
			expectGetData: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("FindClosestMatchIndex", tt.target, tt.margin).Return(tt.mockIndex, tt.mockFindErr)

			if tt.expectGetData {
				mockRepo.On("GetData", tt.mockIndex).Return(tt.mockValue)
			}

			resp, err := service.FindClosest(tt.target, tt.margin)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantResp, resp)

			mockRepo.AssertExpectations(t)
		})
	}
}
