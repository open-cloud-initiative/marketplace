package db

import (
	"context"

	"github.com/open-cloud-initiative/marketplace/internal/models"
	"gorm.io/gorm"
)

// GetPlan ...
func (r *readTxImpl) GetPlan(ctx context.Context, plan *models.Plan) error {
	return r.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.First(plan, plan.ID).Error
	})
}

// CreatePlan ...
func (r *writeTxImpl) CreatePlan(ctx context.Context, plan *models.Plan) error {
	return r.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(plan).Error
	})
}

// UpdatePlan ...
func (r *writeTxImpl) UpdatePlan(ctx context.Context, plan *models.Plan) error {
	return r.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.First(plan, plan.ID)
		if err.Error != nil {
			return err.Error
		}

		return tx.Save(plan).Error
	})
}

// DeletePlan ...
func (r *writeTxImpl) DeletePlan(ctx context.Context, plan *models.Plan) error {
	return r.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.First(plan, plan.ID)
		if err.Error != nil {
			return err.Error
		}

		return tx.Delete(plan).Error
	})
}
