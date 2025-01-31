package utils

import (
	"encoding/json"
	"fmt"
	"gosolve/backend/internal/models"
	"net/http"
)

// Abs ...
func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// EncodeResponse encodes the API response based on the Accept header
func EncodeResponse(w http.ResponseWriter, r *http.Request, response *models.SearchResponse) error {
	accept := r.Header.Get("Accept")

	switch accept {
	case "application/json":
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(response)
	default:
		return fmt.Errorf("unsupported content type: %s", accept)
	}
}
