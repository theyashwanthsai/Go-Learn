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

func TestAdd(t *testing.T){
	t.Run("New word", func(t *testing.T){
		dict := Dictionary{}
		err := dict.Add("Test", "This is just a test")
		assertError(t, err, nil)
		assertDefinition(t, dict, "Test", "This is just a test")
	})

	t.Run("Already exists", func(t *testing.T){
		dict := Dictionary{"Test": "This is just a test"}
		err := dict.Add("Test", "This is just a test")
		assertError(t, err, nil)
		assertDefinition(t, dict, "Test", "This is just a test")
	})
}

func assertDefinition(t testing.TB, dict Dictionary, key, value string){
	t.Helper()
	got, err := dict.Search(key)
	want := value

	if err != nil{
		t.Fatal("Should find added elements only:", err)
	}
	assertString(t, got, want)
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

