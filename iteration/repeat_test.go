package iteration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleRepeat() {
	repeatedThing := Repeat("t", 5)
	fmt.Println(repeatedThing)
	// Output: ttttt
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {
	assert := assert.New(t)
	want := "aaaa"
	got := Repeat("a", 4)

	assert.Equal(want, got)

	t.Run("Negative count returns empty string", func(t *testing.T) {
		want := ""
		got := Repeat("a", -1)

		assert.Equal(want, got)
	})

	t.Run("Zero count returns empty string", func(t *testing.T) {
		want := ""
		got := Repeat("a", 0)

		assert.Equal(want, got)
	})
}
