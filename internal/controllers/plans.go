package controllers

import (
	"github.com/open-cloud-initiative/marketplace/internal/ports"
	pb "github.com/open-cloud-initiative/marketplace/proto/catalog/plans/v1"

	"github.com/katallaxie/pkg/dbx"
)

var _ pb.PlansServiceServer = (*PlansController)(nil)

// PlansController is the controller for managing catalog items.
type PlansController struct {
	store dbx.Database[ports.ReadTx, ports.ReadWriteTx]
	pb.UnimplementedPlansServiceServer
}

// NewPlansController creates a new PlansController.
func NewPlansController(store dbx.Database[ports.ReadTx, ports.ReadWriteTx]) *PlansController {
	return &PlansController{
		store: store,
	}
}
