package main

import (
"testing"
)

func TestHisFunc(t *testing.T) {
	if !myFunc() {
		t.Errorf("expected true, but got false")
	} else {
		print("PASS1\n")
	}

}

func TestHisFunc2(t *testing.T) {
	if myFunc() {
		t.Errorf("expected true, but got false")
	} else {
		print("PASS2\n")
	}
}
