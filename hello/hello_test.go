package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	assert := assert.New(t)

	t.Run("can greet a name", func(t *testing.T) {
		got := Hello("Alex", "")
		want := "Hello, Alex"
		assert.Equal(got, want)
	})

	t.Run("greets world by default", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assert.Equal(got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", SPANISH)
		want := "Hola, Elodie"
		assert.Equal(got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello(
			"Jean Luc",
			FRENCH,
		)
		want := "Bonjour, Jean Luc"
		assert.Equal(got, want)
	})

	t.Run("in Hawaiian", func(t *testing.T) {
		got := Hello(
			"Jean Luc",
			HAWAIIAN,
		)
		want := "Aloha, Jean Luc"
		assert.Equal(got, want)
	})
}
