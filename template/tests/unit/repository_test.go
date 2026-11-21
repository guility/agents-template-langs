package domain_test

import (
	"context"
	"testing"

	"github.com/example/ddd-go-project/src/infrastructure"
)

type record struct {
	ID   string
	Name string
}

func (r *record) GetID() string { return r.ID }

func TestRepositoryMissingEntity(t *testing.T) {
	ctx := context.Background()
	repository := infrastructure.NewInMemoryRepository[record]()
	if _, err := repository.GetByID(ctx, "missing"); err == nil {
		t.Fatal("missing entity should return an error")
	}
	if err := repository.Delete(ctx, "missing"); err == nil {
		t.Fatal("deleting a missing entity should return an error")
	}
}

func TestRepositorySaveAndGetByID(t *testing.T) {
	ctx := context.Background()
	repository := infrastructure.NewInMemoryRepository[record]()
	entity := &record{ID: "one", Name: "sample"}
	if err := repository.Save(ctx, entity); err != nil {
		t.Fatal(err)
	}
	found, err := repository.GetByID(ctx, "one")
	if err != nil || found != entity {
		t.Fatalf("GetByID returned %v, %v", found, err)
	}
}

func TestRepositoryGetAll(t *testing.T) {
	ctx := context.Background()
	repository := infrastructure.NewInMemoryRepository[record]()
	if err := repository.Save(ctx, &record{ID: "one", Name: "sample"}); err != nil {
		t.Fatal(err)
	}
	all, err := repository.GetAll(ctx)
	if err != nil || len(all) != 1 {
		t.Fatalf("GetAll returned %v, %v", all, err)
	}
}

func TestRepositoryFindWhere(t *testing.T) {
	ctx := context.Background()
	repository := infrastructure.NewInMemoryRepository[record]()
	if err := repository.Save(ctx, &record{ID: "one", Name: "sample"}); err != nil {
		t.Fatal(err)
	}
	matched, err := repository.FindWhere(ctx, map[string]interface{}{"Name": "sample"})
	if err != nil || len(matched) != 1 {
		t.Fatalf("FindWhere should match one record: %v, %v", matched, err)
	}
	unmatched, err := repository.FindWhere(ctx, map[string]interface{}{"missing": "sample"})
	if err != nil || len(unmatched) != 0 {
		t.Fatalf("FindWhere should skip unknown fields: %v, %v", unmatched, err)
	}
}

func TestRepositoryDelete(t *testing.T) {
	ctx := context.Background()
	repository := infrastructure.NewInMemoryRepository[record]()
	if err := repository.Save(ctx, &record{ID: "one", Name: "sample"}); err != nil {
		t.Fatal(err)
	}
	if err := repository.Delete(ctx, "one"); err != nil {
		t.Fatal(err)
	}
	if err := repository.Delete(ctx, "one"); err == nil {
		t.Fatal("deleting a missing entity should return an error")
	}
}

func TestRepositoryRejectsInvalidEntities(t *testing.T) {
	ctx := context.Background()
	repository := infrastructure.NewInMemoryRepository[record]()
	if err := repository.Save(ctx, nil); err == nil {
		t.Fatal("nil entity should fail")
	}
	withoutID := infrastructure.NewInMemoryRepository[struct{ Name string }]()
	if err := withoutID.Save(ctx, &struct{ Name string }{Name: "sample"}); err == nil {
		t.Fatal("entity without GetID should fail")
	}
}
