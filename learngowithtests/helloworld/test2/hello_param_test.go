package test2

import "testing"

/*
As part of TDD, you should first write the test
 and then you need to update the code to pass the test
*/
func TestHello(t *testing.T){

	got := Hello("Chris")
	want := "Hello Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}


}