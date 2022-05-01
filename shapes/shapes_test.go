package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	assert := assert.New(t)

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{6.0, 4.0}, 24.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		assert.Equal(tt.want, got)
	}
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
