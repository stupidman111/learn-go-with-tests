package main

import "testing"

func TestArea(t *testing.T) {

	checkArea := func (t *testing.T, shape Shape, want float64)  {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, &rectangle, 72.0)
	}) 

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, &circle, 314.1592653589793)
	})
}

func TestAreaB(t *testing.T) {

    areaTests := []struct {
        shape Shape
        want  float64
    }{
        {Rectangle{12, 6}, 72.0},
        {Circle{10}, 314.1592653589793},

    }

	a := struct{
		r Rectangle
		c Circle
	} {
		r: Rectangle{},
		c: Circle{}, 
	}

    for _, tt := range areaTests {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %.2f want %.2f", got, tt.want)
        }
    }

}