package models

import (
	"time"

	pb "github.com/open-cloud-initiative/marketplace/proto/catalog/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"
	"github.com/katallaxie/pkg/conv"
	"github.com/katallaxie/pkg/protox"
	"gorm.io/gorm"
)

// Catalog represents a catalog
type Catalog struct {
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

	protox.ProtoX[pb.Catalog] `gorm:"-"`
}

// ToProto converts the Catalog to its protobuf representation.
func (c *Catalog) ToProto() *pb.Catalog {
	return &pb.Catalog{
		Id:        conv.String(c.ID),
		Name:      c.Name,
		CreatedAt: timestamppb.New(c.CreatedAt),
		UpdatedAt: timestamppb.New(c.UpdatedAt),
		DeletedAt: timestamppb.New(c.DeletedAt.Time),
	}
}

// FromProto populates the Catalog from its protobuf representation.
func (c *Catalog) FromProto(pb *pb.Catalog) error {
	uuid, err := uuid.Parse(pb.GetId())
	if err != nil {
		return err
	}
	c.ID = uuid

	c.Name = pb.GetName()

	return nil
}
