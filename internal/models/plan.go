package models

import (
	"time"

	pb "github.com/open-cloud-initiative/marketplace/proto/catalog/plans/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"
	"github.com/katallaxie/pkg/conv"
	"github.com/katallaxie/pkg/protox"
	"gorm.io/gorm"
)

// Plan represents a plan.
type Plan struct {
	// ID ...
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	// Name is the tag name.
	Name string `json:"name" gorm:"uniqueIndex;not null"`
	// Description is the tag description.
	Description string `json:"description"`
	// CreatedAt ...
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updatedAt"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deletedAt"`

	protox.ProtoX[pb.Plan] `gorm:"-"`
}

// ToProto converts the Plan to its protobuf representation.
func (c *Plan) ToProto() *pb.Plan {
	return &pb.Plan{
		Id:          conv.String(c.ID),
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   timestamppb.New(c.CreatedAt),
		UpdatedAt:   timestamppb.New(c.UpdatedAt),
		DeletedAt:   timestamppb.New(c.DeletedAt.Time),
	}
}

// FromProto populates the Plan from its protobuf representation.
func (c *Plan) FromProto(pb *pb.Plan) error {
	uuid, err := uuid.Parse(pb.GetId())
	if err != nil {
		return err
	}
	c.ID = uuid
	c.Name = pb.GetName()
	c.Description = pb.GetDescription()

	return nil
}
