package models

import (
	"time"

	"github.com/google/uuid"
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
}
