package domain

import (
	"errors"
	"time"
)

// BaseEntity represents common fields for all domain entities
type BaseEntity struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewBaseEntity creates a domain entity with initialized timestamps.
func NewBaseEntity(id string) BaseEntity {
	now := time.Now().UTC()
	return BaseEntity{ID: id, CreatedAt: now, UpdatedAt: now}
}

func (e BaseEntity) GetID() string           { return e.ID }
func (e BaseEntity) GetCreatedAt() time.Time { return e.CreatedAt }
func (e BaseEntity) GetUpdatedAt() time.Time { return e.UpdatedAt }

func (e BaseEntity) Validate() error {
	if e.ID == "" {
		return errors.New("id is required")
	}
	return nil
}

// Entity defines the interface for all domain entities
type Entity interface {
	GetID() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	Validate() error
}

// ValueObject defines the interface for value objects in DDD
type ValueObject interface {
	Equals(other ValueObject) bool
	Validate() error
}
