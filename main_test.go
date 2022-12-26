package main

import "testing"

func TestHello(t *testing.T) {
	hello := Hello("Adhithia")

	if hello != "Hi, Adhithia" {
		t.Errorf("hasil: %q, diharapkan: %q", hello, "Hi, Adhithia")
	}
}
