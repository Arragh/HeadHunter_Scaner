package serializer

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDeserialize_Valid(t *testing.T) {
	serializedBody, err := json.Marshal(TestStruct{Field: "test"})
	if err != nil {
		t.Errorf("ошибка сериализации тестового body %v", err)
	}

	deserialized, err := Deserialize[TestStruct](serializedBody)
	if err != nil {
		t.Fatalf("ошибка десериализации: %v", err)
	}

	if deserialized == nil {
		t.Fatal("deserialized = nil")
	}

	want := TestStruct{Field: "test"}

	got := *deserialized

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Deserialize() = %v, want %v", got, want)
	}
}

func TestDeserialize_InvalidType(t *testing.T) {
	_, err := Deserialize[TestStruct]([]byte("12345"))
	if err == nil {
		t.Error("ожидалась ошибка десериализации из-за неверного типа")
	}
}

func TestDeserialize_InvalidBody(t *testing.T) {
	_, err := Deserialize[TestStruct]([]byte("{Field: \"test\""))
	if err == nil {
		t.Fatal("ожидалась ошибка десериализации из-за битого body")
	}
}

func TestDeserialize_NilBody(t *testing.T) {
	_, err := Deserialize[TestStruct](nil)
	if err == nil {
		t.Error("ожидалась ошибка десериализации из-за nil-body")
	}
}

type TestStruct struct {
	Field string
}
