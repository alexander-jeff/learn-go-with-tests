package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShapesArea(t *testing.T) {
	assertCorrectArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		assert.Equal(t, want, got)
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{6.0, 4.0}
		want := 24.0
		assertCorrectArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793
		assertCorrectArea(t, circle, want)
	})
}

func TestShapesPerimeters(t *testing.T) {
	assert := assert.New(t)

	t.Run("rectangle calculates its perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 40.0
		got := rectangle.Perimeter()

		assert.Equal(want, got)
	})

	t.Run("circle calculates its circumference", func(t *testing.T) {
		circle := Circle{1.0}
		want := 2 * 3.141592653589793
		got := circle.Circumference()

		assert.Equal(want, got)
	})
}
