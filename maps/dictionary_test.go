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
		want := errorWordNotFound

		assertError(t, got, want)

	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)

	})

	t.Run("existing word", func(t *testing.T) {
		existingWord := "test"
		existingDefinition := "existing"
		dictionary := Dictionary{existingWord: existingDefinition}
		err := dictionary.Add(existingWord, "new definition")

		assertError(t, err, errWordExists)
		assertDefinition(t, dictionary, existingWord, existingDefinition)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("update word existing", func(t *testing.T) {
		word := "test"
		definition := "testdefinition"
		dictionary := Dictionary{word: definition}
		newDefinition := "newDefinition"

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("error if word does not exist", func(t *testing.T) {
		word := "test"
		definition := "testdefinition"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)

		assertError(t, err, errWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete existing word", func (t *testing.T) {
		word := "test"
		definition := "testDefinition"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)
		_, err:=	dictionary.Search(word)
		assertError(t, err, errorWordNotFound)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("Could not find word", err)
	}
	if got != definition {
		t.Errorf("got %q, want %q", got, definition)
	}
}

func assertError(t *testing.T, got, want error) {
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
