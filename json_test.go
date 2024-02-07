package json

import (
	"encoding/json"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {
	// Marshal map
	maps := map[string]string{}
	file, err := os.Open("unit_test_files/test.json")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer file.Close()
	fileContent, _ := io.ReadAll(file)
	err = json.Unmarshal(fileContent, &maps)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	expected, _ := json.Marshal(maps)
	actual, err := Marshal(maps)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(expected) != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Marshal map array
	var maps2 []map[string]string
	file2, err := os.Open("unit_test_files/test2.json")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer file2.Close()
	fileContent, err = io.ReadAll(file2)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	err = json.Unmarshal(fileContent, &maps2)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	expected, _ = json.Marshal(maps2)
	actual, err = Marshal(maps2)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(expected) != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestUnmarshal(t *testing.T) {
	// Unmarshal Json file
	var expected map[string]string
	file, _ := os.Open("unit_test_files/test.json")
	defer file.Close()
	fileContent, _ := io.ReadAll(file)
	_ = json.Unmarshal(fileContent, &expected)
	actual, err := Unmarshal(string(fileContent))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Unmarshal Json Array File
	var expect []map[string]string
	file2, _ := os.Open("unit_test_files/test2.json")
	defer file2.Close()
	fileContent2, _ := io.ReadAll(file2)
	_ = json.Unmarshal(fileContent2, &expect)
	actual, err = Unmarshal(string(fileContent2))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("Expected %v, got %v", expect, actual)
	}
}
