package controllers

import (
	"context"

	"github.com/open-cloud-initiative/marketplace/internal/models"
	"github.com/open-cloud-initiative/marketplace/internal/ports"
	pb "github.com/open-cloud-initiative/marketplace/proto/catalog/v1"

	"github.com/google/uuid"
	"github.com/katallaxie/pkg/dbx"
)

var _ pb.CatalogServiceServer = (*CatalogController)(nil)

// CatalogController is the controller for managing catalog items.
type CatalogController struct {
	store dbx.Database[ports.ReadTx, ports.ReadWriteTx]
	pb.UnimplementedCatalogServiceServer
}

// NewCatalogController creates a new CatalogController.
func NewCatalogController(store dbx.Database[ports.ReadTx, ports.ReadWriteTx]) *CatalogController {
	return &CatalogController{
		store: store,
	}
}

// Create implements pb.CatalogServiceServer
func (c *CatalogController) Create(ctx context.Context, req *pb.CreateCatalogRequest) (*pb.Catalog, error) {
	catalog := &models.Catalog{
		Name: req.Catalog.Name,
	}

	if err := c.store.ReadWriteTx(ctx, func(ctx context.Context, rw ports.ReadWriteTx) error {
		return rw.CreateCatalog(ctx, catalog)
	}); err != nil {
		return nil, err
	}

	return &pb.Catalog{
		Id:   catalog.ID.String(),
		Name: catalog.Name,
	}, nil
}

// Get implements pb.CatalogServiceServer
func (c *CatalogController) Get(ctx context.Context, req *pb.GetCatalogRequest) (*pb.Catalog, error) {
	id, err := uuid.Parse(req.GetCatalogId())
	if err != nil {
		return nil, err
	}

	catalog := &models.Catalog{
		ID: id,
	}

	if err := c.store.ReadTx(ctx, func(ctx context.Context, r ports.ReadTx) error {
		return r.GetCatalog(ctx, catalog)
	}); err != nil {
		return nil, err
	}

	return &pb.Catalog{
		Id:   catalog.ID.String(),
		Name: catalog.Name,
	}, nil
}

// Delete implements pb.CatalogServiceServer
func (c *CatalogController) Delete(ctx context.Context, req *pb.DeleteCatalogRequest) (*pb.DeleteCatalogRequest, error) {
	id, err := uuid.Parse(req.GetCatalog().GetId())
	if err != nil {
		return nil, err
	}

	catalog := &models.Catalog{
		ID: id,
	}

	if err := c.store.ReadWriteTx(ctx, func(ctx context.Context, rw ports.ReadWriteTx) error {
		return rw.DeleteCatalog(ctx, catalog)
	}); err != nil {
		return nil, err
	}

	return &pb.DeleteCatalogRequest{
		Catalog: catalog.ToProto(),
	}, nil
}
