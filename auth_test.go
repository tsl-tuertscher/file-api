package main

import (
	"testing"
)

func TestCheckKey(t *testing.T) {
	keys := []string{"set","feoin","2134n"}
    got := CheckKey(keys, "feoin")
    if got == false {
        t.Errorf("CheckKey()")
    }
}
