package storage

import (
	"fmt"
	"os"
	"slices"
	"testing"
)

func TestReadData_Valid(t *testing.T) {
	tempfile, _ := os.CreateTemp("./", "data")
	fmt.Println(tempfile.Name())
	defer os.Remove(tempfile.Name())

	tempfile.WriteString("1\n2\n3\n")
	tempfile.Close()

	got, err := ReadData(tempfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	want := []int64{1, 2, 3}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestReadData_InvalidData(t *testing.T) {
	tempfile, _ := os.CreateTemp("./", "data")
	defer os.Remove(tempfile.Name())

	tempfile.WriteString("1\nABC\n3\n")
	tempfile.Close()

	_, err := ReadData(tempfile.Name())
	if err == nil {
		t.Fatalf("ожидалась ошибка")
	}
}

func TestSaveData_Valid(t *testing.T) {
	tempfile, _ := os.CreateTemp("./", "data")
	defer os.Remove(tempfile.Name())

	data := []int64{1, 2}
	err := SaveData(data, tempfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	got, _ := os.ReadFile(tempfile.Name())
	want := "1\n2\n"

	if string(got) != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
