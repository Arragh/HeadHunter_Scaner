package serializer

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDeserialize(t *testing.T) {
	testBody := TestStruct{
		Field: "test",
	}

	serializedBody, err := json.Marshal(testBody)
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

	want := TestStruct{
		Field: "test",
	}

	got := *deserialized

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Deserialize() = %v, want %v", got, want)
	}
}

type TestStruct struct {
	Field string
}
