package controllers

import (
	"context"

	"github.com/open-cloud-initiative/marketplace/internal/models"
	"github.com/open-cloud-initiative/marketplace/internal/ports"
	pb "github.com/open-cloud-initiative/marketplace/proto/catalog/plans/v1"

	"github.com/katallaxie/pkg/dbx"
)

var _ pb.PlansServiceServer = (*PlansController)(nil)

// PlansController is the controller for managing catalog items.
type PlansController struct {
	store dbx.Database[ports.ReadTx, ports.ReadWriteTx]
	pb.UnimplementedPlansServiceServer
}

// NewPlansController creates a new PlansController.
func NewPlansController(store dbx.Database[ports.ReadTx, ports.ReadWriteTx]) *PlansController {
	return &PlansController{
		store: store,
	}
}

// Create implements pb.PlansServiceServer.
func (c *PlansController) Create(ctx context.Context, req *pb.CreatePlanRequest) (*pb.Plan, error) {
	plan := &models.Plan{}

	if err := plan.FromProto(req.Plan); err != nil {
		return nil, err
	}

	if err := c.store.ReadWriteTx(ctx, func(ctx context.Context, rw ports.ReadWriteTx) error {
		return rw.CreatePlan(ctx, plan)
	}); err != nil {
		return nil, err
	}

	return plan.ToProto(), nil
}

// Get implements pb.PlansServiceServer.
func (c *PlansController) Get(ctx context.Context, req *pb.GetPlanRequest) (*pb.Plan, error) {
	plan := &models.Plan{}

	if err := plan.FromProto(req.GetPlan()); err != nil {
		return nil, err
	}

	if err := c.store.ReadTx(ctx, func(ctx context.Context, r ports.ReadTx) error {
		return r.GetPlan(ctx, plan)
	}); err != nil {
		return nil, err
	}

	return plan.ToProto(), nil
}

// Update implements pb.PlansServiceServer.
func (c *PlansController) Update(ctx context.Context, req *pb.UpdatePlanRequest) (*pb.Plan, error) {
	plan := &models.Plan{}

	if err := plan.FromProto(req.GetPlan()); err != nil {
		return nil, err
	}

	if err := c.store.ReadWriteTx(ctx, func(ctx context.Context, rw ports.ReadWriteTx) error {
		return rw.UpdatePlan(ctx, plan)
	}); err != nil {
		return nil, err
	}

	return plan.ToProto(), nil
}

// Delete implements pb.PlansServiceServer.
func (c *PlansController) Delete(ctx context.Context, req *pb.DeletePlanRequest) (*pb.Plan, error) {
	plan := &models.Plan{}

	if err := plan.FromProto(req.GetPlan()); err != nil {
		return nil, err
	}

	if err := c.store.ReadWriteTx(ctx, func(ctx context.Context, rw ports.ReadWriteTx) error {
		return rw.DeletePlan(ctx, plan)
	}); err != nil {
		return nil, err
	}

	return plan.ToProto(), nil
}
