package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("can greet a name", func(t *testing.T) {
		got := Hello("Alex", "")
		want := "Hello, Alex"
		assertCorrectMessage(t, got, want)
	})

	t.Run("greets world by default", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", SPANISH)
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello(
			"Jean Luc",
			FRENCH,
		)
		want := "Bonjour, Jean Luc"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Hawaiian", func(t *testing.T) {
		got := Hello(
			"Jean Luc",
			HAWAIIAN,
		)
		want := "Aloha, Jean Luc"
		assertCorrectMessage(t, got, want)
	})
}
