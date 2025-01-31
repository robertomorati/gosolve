package errors

import (
	"net/http"
)

// APIERrror ...
type APIError struct {
	Message    string
	StatusCode int
}

func (e *APIError) Error() string {
	return e.Message
}

// New ...
func New(message string, statusCode int) *APIError {
	return &APIError{
		Message:    message,
		StatusCode: statusCode,
	}
}

var (
	ErrInvalidRequest       = New("Invalid request format", http.StatusBadRequest)
	ErrInvalidRequestValue  = New("Invalid request value", http.StatusBadRequest)
	ErrNoClosestMatchFound  = New("No closest match found", http.StatusNotFound)
	ErrDataNotLoaded        = New("Data is empty", http.StatusNotFound)
	ErrInvalidIndexFromRepo = New("Invalid index from repository", http.StatusNotFound)
	ErrInputFileNotSet      = New("Input file not set", http.StatusInternalServerError)
	ErrFailedToInitiateRepo = New("Failed to initiate repository", http.StatusInternalServerError)
	ErrProcessingData       = New("Error processing data", http.StatusInternalServerError)
	ErrInvalidValueInFile   = New("Invalid value in file", http.StatusInternalServerError)
	ErrInternalServer       = New("Internal server error", http.StatusInternalServerError)
)
