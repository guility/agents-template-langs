package domain_test

import (
	"testing"
	"time"

	"github.com/example/ddd-go-project/src/domain"
)

func TestBaseEntity_Creation(t *testing.T) {
	before := time.Now().UTC()
	entity := domain.NewBaseEntity("test-id")
	after := time.Now().UTC()

	if entity.ID != "test-id" {
		t.Errorf("Expected ID 'test-id', got '%s'", entity.ID)
	}

	if entity.CreatedAt.Before(before) || entity.CreatedAt.After(after) {
		t.Error("CreatedAt should be initialized at construction")
	}

	if !entity.UpdatedAt.Equal(entity.CreatedAt) {
		t.Error("UpdatedAt should match CreatedAt for a new entity")
	}
	if entity.GetID() != entity.ID || entity.GetCreatedAt() != entity.CreatedAt || entity.GetUpdatedAt() != entity.UpdatedAt {
		t.Error("entity accessors should expose the stored values")
	}
	if err := entity.Validate(); err != nil {
		t.Errorf("valid entity returned an error: %v", err)
	}
}

func TestBaseEntity_Timestamps(t *testing.T) {
	before := time.Now()
	entity := domain.NewBaseEntity("test")
	after := time.Now()

	if entity.CreatedAt.Before(before) || entity.CreatedAt.After(after) {
		t.Error("CreatedAt should be within expected range")
	}

	if entity.UpdatedAt.Before(before) || entity.UpdatedAt.After(after) {
		t.Error("UpdatedAt should be within expected range")
	}
}

func TestBaseEntity_ValidateRejectsEmptyID(t *testing.T) {
	if err := (domain.BaseEntity{}).Validate(); err == nil {
		t.Fatal("empty ID should fail validation")
	}
}
