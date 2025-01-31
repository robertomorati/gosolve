package models

// SearchRequest ...
type SearchRequest struct {
	Target int `json:"target"`
}

type SearchResponse struct {
	Index   int    `json:"index"`
	Value   int    `json:"value"`
	Message string `json:"message"`
}
