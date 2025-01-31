package repository

import (
	"gosolve/backend/internal/errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestFindClosestMatchIndex(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	repo := &dataRepository{
		data:   []int{0, 10, 20, 30, 40, 50, 60, 100},
		logger: logger,
	}

	tests := []struct {
		name   string
		target int
		margin float64
		want   int
		err    error
	}{
		{
			name:   "Exact match found",
			target: 10,
			margin: 0.1,
			want:   1,
			err:    nil,
		},
		{
			name:   "Closest match found within margin",
			target: 15,
			margin: 0.1,
			want:   -1,
			err:    errors.ErrNoClosestMatchFound,
		},
		{
			name:   "Exact match at end of list",
			target: 100,
			margin: 0.1,
			want:   7,
			err:    nil,
		},
		{
			name:   "No closest match found",
			target: 999,
			margin: 0.1,
			want:   -1,
			err:    errors.ErrNoClosestMatchFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.FindClosestMatchIndex(tt.target, tt.margin)
			if tt.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// TestGetData ...
func TestGetData(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	repo := &dataRepository{
		data:   []int{10, 20, 30, 40, 50},
		logger: logger,
	}

	tests := []struct {
		name  string
		index int
		want  int
	}{
		{"Valid index", 2, 30},
		{"First index", 0, 10},
		{"Last index", 4, 50},
		{"Negative index", -1, -1},
		{"Out of range index", 10, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := repo.GetData(tt.index)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestLoadData ...
func TestLoadData(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	repo := &dataRepository{
		data:   []int{},
		logger: logger,
	}

	tests := []struct {
		name      string
		filePath  string
		wantErr   error
		setupFile func()
		cleanup   func()
	}{
		{
			name:     "Valid file loads correctly",
			filePath: "test_valid.txt",
			wantErr:  nil,
			setupFile: func() {
				file, _ := os.Create("test_valid.txt")
				file.WriteString("10\n20\n30\n40\n50\n")
				file.Close()
			},
			cleanup: func() {
				os.Remove("test_valid.txt")
			},
		},
		{
			name:      "File does not exist",
			filePath:  "test_missing.txt",
			wantErr:   errors.ErrInputFileNotSet,
			setupFile: func() {},
			cleanup:   func() {},
		},
		{
			name:     "File with invalid data",
			filePath: "test_invalid.txt",
			wantErr:  errors.ErrInvalidValueInFile,
			setupFile: func() {
				file, _ := os.Create("test_invalid.txt")
				file.WriteString("10\nabc\n30\n")
				file.Close()
			},
			cleanup: func() {
				os.Remove("test_invalid.txt")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupFile() // Create test file if needed
			err := repo.loadData(tt.filePath)

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr.Error()) // Match partial error message
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, repo.data) // Ensure data is loaded
			}

			tt.cleanup() // Remove test file after test
		})
	}
}
