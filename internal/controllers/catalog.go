package controllers

import (
	"github.com/open-cloud-initiative/marketplace/internal/ports"

	"github.com/katallaxie/pkg/dbx"
	pb "github.com/open-cloud-initiative/marketplace/proto/catalog/v1"
)

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
