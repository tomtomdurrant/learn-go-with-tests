package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is a test"}
	t.Run("Dictionary contains word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is a test"

		assertString(t, got, want)
	})
	t.Run("Dictionary doesn't contain word", func(t *testing.T) {
		_, got := dictionary.Search("Not there")
		want := errorMesssage 

		assertError(t, got, want)

	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is a test")

	want :=  "this is a test"
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("Could not find word", err)
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}








func assertError(t *testing.T, got, want error)  {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
