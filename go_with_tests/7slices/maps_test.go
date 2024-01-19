package maps

import (
	"testing"
)

func TestSearch(t *testing.T){
	t.Run("Known word", func(t *testing.T){
		dict := Dictionary{"Test": "This is just a test"}

		got,_ := dict.Search("Test")
		want := "This is just a test"
		assertString(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T){
		dict := Dictionary{"Test": "This is just a test"}
		_, err := dict.Search("Unknown")
		assertError(t, err, NotFoundError)
	})
}

func assertString(t testing.TB, got, want string){
	t.Helper()
	if got != want{
		t.Errorf("got %q, want %q given %q", got, want, "Test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}