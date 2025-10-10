package ports

import (
	"context"

	"github.com/open-cloud-initiative/marketplace/internal/models"
)

// PlanReadTx provides methods for transactional read operations on plans.
type PlanReadTx interface {
	// GetPlan retrieves a plan by its ID.
	GetPlan(ctx context.Context, plan *models.Plan) error
}

// PlanReadWriteTx provides methods for transactional write operations on plans.
type PlanWriteTx interface {
	// CreatePlan creates a new plan.
	CreatePlan(ctx context.Context, plan *models.Plan) error
	// UpdatePlan updates an existing plan.
	UpdatePlan(ctx context.Context, plan *models.Plan) error
	// DeletePlan deletes a plan.
	DeletePlan(ctx context.Context, plan *models.Plan) error
}
