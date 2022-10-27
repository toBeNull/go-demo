package main

import (
	"testing"
)

func TestAFunc(t *testing.T) {
	if !aFunc() {
		t.Errorf("expected true, but got false")
	} else {
		print("PASS1\n")
	}

}

func TestAFunc2(t *testing.T) {
	if !aFunc() {
		t.Errorf("expected true, but got false")
	} else {
		print("PASS2\n")
	}
}
