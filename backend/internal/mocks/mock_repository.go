package mocks

import (
	"context"
	"gosolve/backend/internal/models"
	"gosolve/backend/internal/repository"

	"github.com/stretchr/testify/mock"
)

var _ repository.DataRepository = (*MockRepository)(nil)

// MockRepository ...
type MockRepository struct {
	mock.Mock
}

// GetData ...
func (m *MockRepository) GetData(index int) int {
	args := m.Called(index)
	return args.Int(0)
}

// FindClosestMatchIndex ...
func (m *MockRepository) FindClosestMatchIndex(target int, margin float64) (int, error) {
	args := m.Called(target, margin)
	return args.Int(0), args.Error(1)
}

// LoadData ...
func (m *MockRepository) LoadData(filePath string) error {
	args := m.Called(filePath)
	return args.Error(0)
}

// MockController ...
type MockController struct {
	mock.Mock
}

// SearchValue ...
func (m *MockController) SearchValue(ctx context.Context, req *models.SearchRequest) (*models.SearchResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*models.SearchResponse), args.Error(1)
}
