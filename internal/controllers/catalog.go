package controllers

import (
	"context"

	"github.com/open-cloud-initiative/marketplace/internal/models"
	"github.com/open-cloud-initiative/marketplace/internal/ports"

	"github.com/katallaxie/pkg/dbx"
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
func (c *CatalogController) Create(ctx context.Context, req *pb.CreateCatalogRequest) (*pb.Catalog, error) {
	catalog := models.Catalog{
		Name: req.Catalog.Name,
	}

	if err := c.store.ReadWriteTx(ctx, func(ctx context.Context, rw ports.ReadWriteTx) error {
		return rw.CreateCatalog(ctx, &catalog)
	}); err != nil {
		return nil, err
	}

	return &pb.Catalog{
		Id:   catalog.ID.String(),
		Name: catalog.Name,
	}, nil
}
