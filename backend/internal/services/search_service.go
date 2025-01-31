package services

import (
	"fmt"

	"gosolve/backend/internal/errors"
	"gosolve/backend/internal/models"
	"gosolve/backend/internal/repository"

	"go.uber.org/zap"
)

type searchService struct {
	repository repository.DataRepository
	logger     *zap.Logger
}

// NewSearchService ...
func NewSearchService(repo repository.DataRepository, logger *zap.Logger) *searchService {
	return &searchService{
		repository: repo,
		logger:     logger,
	}
}

// SearchService ...
type SearchService interface {
	FindClosest(target int, margin float64) (*models.SearchResponse, error)
}

// FindClosest
func (s *searchService) FindClosest(target int, margin float64) (*models.SearchResponse, error) {

	s.logger.Info("Searching for target value", zap.Int("target", target))
	index, err := s.repository.FindClosestMatchIndex(target, margin)
	if err != nil {
		fmt.Println("Error searching for target value: ", err)
		return nil, err
	}

	fmt.Println("Closest index found: ", index)
	value := s.repository.GetData(index)
	fmt.Println("Closest value found: ", value)
	if value == -1 {
		fmt.Println("Invalid index from repository")
		return nil, errors.ErrInvalidIndexFromRepo
	}

	// Return the response
	fmt.Println("Closest match found")
	fmt.Println("Value: ", value)
	fmt.Println("Index: ", index)

	return &models.SearchResponse{
		Index:   index,
		Value:   value,
		Message: "Closest match found",
	}, nil
}
