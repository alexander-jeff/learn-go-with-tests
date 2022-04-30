package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleSum() {
	numbers := []int{1, 2, 3, 5}
	fmt.Println(Sum(numbers))
	// Output: 11
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func TestSum(t *testing.T) {
	t.Run("empty returns zero", func(t *testing.T) {
		empty := []int{}

		want := 0
		got := Sum(empty)

		assert.Equal(t, want, got)
	})

	t.Run("sums integers in a slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		want := 1 + 2 + 3 + 4 + 5
		got := Sum(numbers)

		assert.Equal(t, want, got)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("returns a slice of sums given a slice of int slices", func(t *testing.T) {
		want := []int{2, 11}
		got := SumAll([]int{0, 2}, []int{1, 2,  3, 5})

		assert.Equal(t, want, got)
	})
}
