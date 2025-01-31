package controller

import (
	"context"
	"gosolve/backend/internal/models"
	"gosolve/backend/internal/services"
)

// Params ...
type Params struct {
	SearchService services.SearchService
}

type controller struct {
	p Params
}

// New returns an instance of Controller
func New(p Params) (Controller, error) {
	return newController(p)
}

func newController(p Params) (*controller, error) {
	return &controller{
		p: p,
	}, nil
}

// Controller interface for the component controller
type Controller interface {
	SearchValue(ctx context.Context, request *models.SearchRequest) (*models.SearchResponse, error)
}

func (c *controller) SearchValue(ctx context.Context, request *models.SearchRequest) (*models.SearchResponse, error) {
	//
	return c.p.SearchService.FindClosest(request.Target, 0.1)
}
