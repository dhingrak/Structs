package main

import "testing"

func TestIsWorking(t *testing.T) {

	result := true

	if result != true {
		t.Errorf("Assert failed, Expected true, got False, hey testing")
	}

}
