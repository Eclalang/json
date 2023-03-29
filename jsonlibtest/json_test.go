package jsonlibtest

import (
	"encoding/json"
	PackageJson "github.com/Eclalang/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {
	// Marshal map
	maps := map[string]interface{}{"name": "Ecla", "members": 7, "language": "Go"}
	expected, _ := json.Marshal(maps)
	actual := PackageJson.Marshal(maps)
	if string(expected) != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Marshal map array
	maps2 := []map[string]interface{}{{"name": "Ecla", "members": 7, "language": "Go"}, {"name": "NewProject", "members": 3, "language": "Python"}}
	expected, _ = json.Marshal(maps2)
	actual = PackageJson.Marshal(maps2)
	if string(expected) != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestUnmarshal(t *testing.T) {
	// Unmarshal Json file
	var expected map[string]interface{}
	file, _ := os.Open("test.json")
	defer file.Close()
	fileContent, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal(fileContent, &expected)
	actual := PackageJson.Unmarshal(string(fileContent))
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Unmarshal Json Array File
	var expect []map[string]interface{}
	file2, _ := os.Open("test2.json")
	defer file2.Close()
	fileContent2, _ := ioutil.ReadAll(file2)
	_ = json.Unmarshal(fileContent2, &expect)
	actual = PackageJson.Unmarshal(string(fileContent2))
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("Expected %v, got %v", expect, actual)
	}
}
