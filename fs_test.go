package main

import (
	"testing"
)

func TestFileExists(t *testing.T) {
	got := FileExists("./test.json")
	if got == true {
		t.Errorf("FileExists('./test.json')")
	}
}

func TestGetMimeTypeFromFileType(t *testing.T) {
	got := GetMimeTypeFromFileType("hgt")
	if got != "application/octet-stream" {
		t.Errorf("GetMimeTypeFromFileType('hgt')")
	}

	got = GetMimeTypeFromFileType("json")
	if got != "application/json" {
		t.Errorf("FileExists('json')")
	}
}

func TestGetTileUrl(t *testing.T) {
	var result Config
    result.Offset = ""
      
	got := GetTileUrl(result, "test", "3", "3", "4", "json")
    if got != "/test/default.json" {
		t.Errorf("GetTileUrl()")
	}
}
