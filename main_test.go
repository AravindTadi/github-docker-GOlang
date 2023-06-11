package main

import "testing"

// in the pipeline we need to test
func Test_sayHello(t *testing.T) {
	name := "Aravind"
	want := "Hello Aravind"

	if got := sayHello(name); got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}
