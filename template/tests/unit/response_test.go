package domain_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	handler "github.com/example/ddd-go-project/src/interface"
)

func TestWriteJSONResponse(t *testing.T) {
	writer := httptest.NewRecorder()
	if err := handler.WriteJSON(writer, 201, map[string]string{"id": "one"}); err != nil {
		t.Fatal(err)
	}
	if writer.Code != 201 || writer.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("unexpected response metadata: %d, %s", writer.Code, writer.Header())
	}
	var payload map[string]string
	if err := json.Unmarshal(writer.Body.Bytes(), &payload); err != nil || payload["id"] != "one" {
		t.Fatalf("unexpected JSON response: %v, %v", payload, err)
	}
}

func TestWriteResponseHelpers(t *testing.T) {
	success := httptest.NewRecorder()
	if err := handler.WriteSuccess(success, "ok"); err != nil || success.Code != 200 {
		t.Fatalf("success helper failed: %v", err)
	}
	var successPayload handler.ResponseDTO
	if err := json.Unmarshal(success.Body.Bytes(), &successPayload); err != nil {
		t.Fatal(err)
	}
	if !successPayload.Success || successPayload.Data != "ok" {
		t.Fatalf("unexpected success payload: %+v", successPayload)
	}

	failure := httptest.NewRecorder()
	if err := handler.WriteError(failure, 400, "bad"); err != nil || failure.Code != 400 {
		t.Fatalf("error helper failed: %v", err)
	}
	var failurePayload handler.ResponseDTO
	if err := json.Unmarshal(failure.Body.Bytes(), &failurePayload); err != nil {
		t.Fatal(err)
	}
	if failurePayload.Success || failurePayload.Error != "bad" {
		t.Fatalf("unexpected error payload: %+v", failurePayload)
	}
}

func TestWriteJSONUnencodablePayload(t *testing.T) {
	invalid := httptest.NewRecorder()
	if err := handler.WriteJSON(invalid, 200, make(chan int)); err == nil {
		t.Fatal("unencodable response should return an error")
	}
}
