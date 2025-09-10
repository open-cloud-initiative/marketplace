package ports

import (
	"context"

	"github.com/open-cloud-initiative/marketplace/internal/models"

	"github.com/katallaxie/pkg/dbx"
)

// Datastore provides methods for transactional operations.
type Datastore interface {
	// ReadTx starts a read only transaction.
	ReadTx(context.Context, func(context.Context, ReadTx) error) error
	// ReadWriteTx starts a read write transaction.
	ReadWriteTx(context.Context, func(context.Context, ReadWriteTx) error) error

	dbx.Migrator
}

// ReadTx provides methods for transactional read operations.
type ReadTx interface {
	GetCatalog(ctx context.Context, catalog *models.Catalog) error
}

// ReadWriteTx provides methods for transactional read and write operations.
type ReadWriteTx interface {
	// CreateCatalog creates a new catalog.
	CreateCatalog(ctx context.Context, catalog *models.Catalog) error
	// UpdateCatalog updates an existing catalog.
	UpdateCatalog(ctx context.Context, catalog *models.Catalog) error
	// DeleteCatalog deletes a catalog.
	DeleteCatalog(ctx context.Context, catalog *models.Catalog) error

	ReadTx
}
