package repository

import (
	"bufio"
	"gosolve/backend/internal/errors"
	"gosolve/backend/internal/utils"
	"math"
	"os"
	"strconv"
	"sync"

	"go.uber.org/zap"
)

// DataRepository ...
type DataRepository interface {
	GetData(index int) int
	FindClosestMatchIndex(target int, margin float64) (int, error)
}

// dataRepository ...
type dataRepository struct {
	data   []int
	logger *zap.Logger
	mu     sync.RWMutex
}

// NewDataRepository ...
func NewDataRepository(logger *zap.Logger) (DataRepository, error) {
	filePath := os.Getenv("INPUT_FILE")

	if filePath == "" {
		return nil, errors.ErrInputFileNotSet
	}

	repo := &dataRepository{
		logger: logger,
	}

	logger.Info("Initializing data repository", zap.String("file_path", filePath))
	// Load data from file
	if err := repo.loadData(filePath); err != nil {
		logger.Error("Error failed to initialize repository", zap.Error(err))
		return nil, errors.ErrFailedToInitiateRepo
	}
	return repo, nil
}

// GetData ...
func (r *dataRepository) GetData(index int) int {

	if index < 0 || index >= len(r.data) {
		return -1
	}
	return r.data[index]
}

func (r *dataRepository) loadData(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.ErrInputFileNotSet
	}
	defer file.Close()

	var data []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			r.logger.Error("Error invalid value in file", zap.String("value", scanner.Text()), zap.Error(err))
			return errors.ErrInvalidValueInFile
		}
		data = append(data, value)
	}

	if err := scanner.Err(); err != nil {
		r.logger.Error("Error reading file", zap.Error(err))
		return errors.ErrProcessingData
	}

	// Safely assign data to the repository
	r.mu.RLock()
	r.data = data
	r.mu.RUnlock()

	r.logger.Info("Data loaded successfully", zap.Int("data_size", len(data)))
	return nil
}

// FindClosestMatchIndex executes a BS (Binary Seaarch) to find the target value
func (r *dataRepository) FindClosestMatchIndex(target int, margin float64) (int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.data) == 0 {
		return -1, errors.ErrDataNotLoaded
	}

	low, high := 0, len(r.data)-1
	closestIndex := -1
	minDiff := math.MaxInt
	acceptableMargin := int(math.Round(float64(target) * margin))

	// Binary Search
	for low <= high {
		mid := (low + high) / 2

		if r.data[mid] == target {
			return mid, nil
		}

		diff := utils.Abs(r.data[mid] - target)

		// verifies if the current value is within the acceptable margin
		if diff <= acceptableMargin && diff < minDiff {
			minDiff = diff
			closestIndex = mid
		}

		if r.data[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if closestIndex != -1 && utils.Abs(r.data[closestIndex]-target) <= acceptableMargin {
		return closestIndex, nil
	}

	return closestIndex, errors.ErrNoClosestMatchFound
}
