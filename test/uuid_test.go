package test

import (
	"encoding/json"
	uuidGen "github.com/elastic/go-json-schema-generate/test/uuid_gen"
	"github.com/google/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	data := []byte(`{
	"id": "f1c85b8a-3872-40d1-84db-6c0cb9c3ca23"
    }`)
	uuidStruct := &uuidGen.UUID{}
	if err := json.Unmarshal(data, &uuidStruct); err != nil {
		t.Fatal(err)
	}
	expectedUUID, err := uuid.Parse("f1c85b8a-3872-40d1-84db-6c0cb9c3ca23")
	if err != nil {
		t.Fatal(err)
	}
	if uuidStruct.Id != expectedUUID {
		t.Errorf("expected uuid to be %s got %s", expectedUUID.String(), uuidStruct.Id.String())
	}
}
