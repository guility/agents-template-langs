package e2e_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/example/ddd-go-project/src/domain"
	"github.com/example/ddd-go-project/src/infrastructure"
	handler "github.com/example/ddd-go-project/src/interface"
)

func TestStoredEntityIsReturnedAsSuccessResponse(t *testing.T) {
	repository := infrastructure.NewInMemoryRepository[domain.BaseEntity]()
	entity := domain.NewBaseEntity("account-1")
	if err := repository.Save(context.Background(), &entity); err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/accounts/account-1", nil)
	account, err := repository.GetByID(request.Context(), "account-1")
	if err != nil {
		t.Fatal(err)
	}
	if err := handler.WriteSuccess(recorder, account); err != nil {
		t.Fatal(err)
	}

	var response handler.ResponseDTO
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}
	if recorder.Code != http.StatusOK || !response.Success {
		t.Fatalf("unexpected response: status=%d body=%+v", recorder.Code, response)
	}
}
