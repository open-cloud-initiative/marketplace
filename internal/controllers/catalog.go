package controllers

import (
	"context"

	"github.com/open-cloud-initiative/marketplace/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/katallaxie/pkg/dbx"
	ops "github.com/open-cloud-initiative/marketplace/proto"
	pb "github.com/open-cloud-initiative/marketplace/proto/catalog/v1"
)

var _ pb.CatalogServiceServer = (*CatalogController)(nil)

// CatalogController is the controller for managing catalog items.
type CatalogController struct {
	store dbx.Database[ports.ReadTx, ports.ReadWriteTx]
}

// NewCatalogController creates a new CatalogController.
func NewCatalogController(store dbx.Database[ports.ReadTx, ports.ReadWriteTx]) *CatalogController {
	return &CatalogController{
		store: store,
	}
}

// Create implements pb.CatalogServiceServer
func (c *CatalogController) Create(ctx context.Context, req *pb.CreateCatalogRequest) (*ops.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
