package integration_test

import (
	"context"
	"testing"

	"github.com/example/ddd-go-project/src/domain"
	"github.com/example/ddd-go-project/src/infrastructure"
)

func TestRepositoryPersistsDomainEntityAcrossOperations(t *testing.T) {
	ctx := context.Background()
	repository := infrastructure.NewInMemoryRepository[domain.BaseEntity]()
	entity := domain.NewBaseEntity("account-1")
	if err := repository.Save(ctx, &entity); err != nil {
		t.Fatal(err)
	}

	found, err := repository.FindWhere(ctx, map[string]interface{}{"ID": "account-1"})
	if err != nil || len(found) != 1 || found[0].ID != "account-1" {
		t.Fatalf("domain entity was not persisted: %#v, %v", found, err)
	}
}
