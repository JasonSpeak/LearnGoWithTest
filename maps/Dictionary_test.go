package maps

import "testing"

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' but want '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got error '%s' want '%s'", got, want)
    }
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }

    if definition != got {
        t.Errorf("got '%s' want '%s'", got, definition)
    }
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known words", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("Unknown words", func(t *testing.T) {
		_, err := dictionary.Search("Unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("except to get a error")
		}
		assertStrings(t, err.Error(), want)
	})

	t.Run("Add key-value", func(t *testing.T) {
		d := Dictionary{}
		d.Add("New", "New key-value")

		want := "New key-value"
		got, err := d.Search("New")

		if err != nil {
			t.Fatal("should find added word:", err)
		}

		if want != got {
			t.Errorf("got '%s' but want '%s'", got, want)
		}
	})

	t.Run("Add exist key", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}
