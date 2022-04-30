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

func ExampleSumAll() {
	slicesToSum := [][]int{
		{1, 1, 1},
		{1, 2, 3, 5},
	}
	fmt.Println(SumAll(slicesToSum...))
	// Output: [3 11]
}

func ExampleSumAllTails() {
	slicesToSum := [][]int{
		{1, 1, 1},
		{1, 2, 3, 5},
	}
	fmt.Println(SumAllTails(slicesToSum...))
	// Output: [2 10]
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
		got := SumAll([]int{0, 2}, []int{1, 2, 3, 5})

		assert.Equal(t, want, got)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("returns a slice of tail sums given a slice of int slices", func(t *testing.T) {
		slicesToSum := [][]int{
			{1, 1, 1},
			{1, 2, 3, 5},
		}
		want := []int{2, 10}
		got := SumAllTails(slicesToSum...)

		assert.Equal(t, want, got)
	})

	t.Run("safely handles empty slices", func(t *testing.T) {
		slicesToSum := [][]int{
			{},
			{1, 2},
			{1},
		}
		want := []int{0, 2, 0}
		got := SumAllTails(slicesToSum...)

		assert.Equal(t, want, got)
	})
}
