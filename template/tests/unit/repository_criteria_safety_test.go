package domain_test

import (
	"context"
	"testing"

	"github.com/example/ddd-go-project/src/infrastructure"
)

type hiddenFieldRecord struct {
	ID   string
	name string
}

func (r *hiddenFieldRecord) GetID() string { return r.ID }

type scalarRecord string

func (r *scalarRecord) GetID() string { return string(*r) }

func TestFindWhereRejectsUnexportedField(t *testing.T) {
	repository := infrastructure.NewInMemoryRepository[hiddenFieldRecord]()
	if err := repository.Save(context.Background(), &hiddenFieldRecord{ID: "one", name: "sample"}); err != nil {
		t.Fatal(err)
	}
	if _, err := repository.FindWhere(context.Background(), map[string]interface{}{"name": "sample"}); err == nil {
		t.Fatal("filtering by an unexported field should return an error")
	}
}

func TestFindWhereRejectsNonStructEntity(t *testing.T) {
	repository := infrastructure.NewInMemoryRepository[scalarRecord]()
	entity := scalarRecord("one")
	if err := repository.Save(context.Background(), &entity); err != nil {
		t.Fatal(err)
	}
	if _, err := repository.FindWhere(context.Background(), map[string]interface{}{"Value": "one"}); err == nil {
		t.Fatal("filtering a non-struct entity should return an error")
	}
}
