package db

import (
	"context"

	"github.com/open-cloud-initiative/marketplace/internal/models"
	"github.com/open-cloud-initiative/marketplace/internal/ports"

	"github.com/katallaxie/pkg/dbx"
	"gorm.io/gorm"
)

var _ ports.ReadTx = (*readTxImpl)(nil)

type readTxImpl struct {
	conn *gorm.DB
}

// NewReadTx ...
func NewReadTx() dbx.ReadTxFactory[ports.ReadTx] {
	return func(db *gorm.DB) (ports.ReadTx, error) {
		return &readTxImpl{conn: db}, nil
	}
}

type writeTxImpl struct {
	conn *gorm.DB
	readTxImpl
}

// NewWriteTx ...
func NewWriteTx() dbx.ReadWriteTxFactory[ports.ReadWriteTx] {
	return func(db *gorm.DB) (ports.ReadWriteTx, error) {
		return &writeTxImpl{conn: db}, nil
	}
}

// CreateCatalog ...
func (r *writeTxImpl) CreateCatalog(ctx context.Context, catalog *models.Catalog) error {
	return r.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(catalog).Error
	})
}

// GetCatalog ...
func (r *readTxImpl) GetCatalog(ctx context.Context, catalog *models.Catalog) error {
	return r.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.First(catalog, catalog.ID).Error
	})
}

// UpdateCatalog ...
func (r *writeTxImpl) UpdateCatalog(ctx context.Context, catalog *models.Catalog) error {
	return r.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.First(catalog, catalog.ID)
		if err.Error != nil {
			return err.Error
		}

		return tx.Save(catalog).Error
	})
}

// DeleteCatalog ...
func (r *writeTxImpl) DeleteCatalog(ctx context.Context, catalog *models.Catalog) error {
	return r.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Delete(catalog).Error
	})
}
