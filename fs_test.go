package main

import (
	"testing"
)

func TestFileExists(t *testing.T) {
    got := FileExists("./test.json")
    if got == true{
        t.Errorf("fileExists('./test.json')")
    }
}

