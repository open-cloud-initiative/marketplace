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
func (r *readTxImpl) CreateCatalog(ctx context.Context, design *models.Catalog) error {
	return r.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(design).Error
	})
}
