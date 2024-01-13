package main

import(
	"testing"
)

func TestHello(t *testing.T){
	
	t.Run("Saying hello to ppl", func(t *testing.T){
		got := hello("Sai", "english")
		want := "Hello, Sai"
		assert(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := hello("Sai", "spanish")
		want := "Hola, Sai"
		assert(t, got, want)
	})

	t.Run("in french", func(t *testing.T){
		got := hello("Sai", "french")
		want := "Bonjour, Sai"
		assert(t, got, want)
	})

	t.Run("", func(t *testing.T){
		got := hello("", "")
		want := "Hello, World"
		assert(t, got, want)
	})
}


func assert(t testing.TB, got string, want string){
	t.Helper()
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}