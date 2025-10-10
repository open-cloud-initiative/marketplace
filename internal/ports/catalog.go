package ports

import (
	"context"

	"github.com/open-cloud-initiative/marketplace/internal/models"
)

// CatalogReadTx provides methods for transactional read operations on catalogs.
type CatalogReadTx interface {
	// GetCatalog retrieves a catalog by its ID.
	GetCatalog(ctx context.Context, catalog *models.Catalog) error
}

// CatalogReadWriteTx provides methods for transactional write operations on catalogs.
type CatalogWriteTx interface {
	// CreateCatalog creates a new catalog.
	CreateCatalog(ctx context.Context, catalog *models.Catalog) error
	// UpdateCatalog updates an existing catalog.
	UpdateCatalog(ctx context.Context, catalog *models.Catalog) error
	// DeleteCatalog deletes a catalog.
	DeleteCatalog(ctx context.Context, catalog *models.Catalog) error
}
