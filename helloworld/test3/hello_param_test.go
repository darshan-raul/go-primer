package test3

import "testing"

/*
As part of TDD, you should first write the test
 and then you need to update the code to pass the test
*/
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris","")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("","")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T){

		got := Hello("Chris","Spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)

	})
	t.Run("in french", func(t *testing.T){

		got := Hello("Chris","French")
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	
// t.Helper() is needed to tell the test suite that this method is a helper. 
// By doing this when it fails the line number reported will be in our function call 
// rather than inside our test helper. This will help other developers track down problems easier. 
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}