package structs

import "testing"

func TestArea(t *testing.T){

	areaTests := []struct{
		shape Shape
		want float64
	}{
		{Rectangle{12,6}, 72.0},
		{Circle{10},314.1592653589793},
		{Triangle{base: 12, height: 6}, 36.0},
	}

	for _, tt := range  areaTests{
		got := tt.shape.Area()
		if got != tt.want{
			t.Errorf("%#v got %g want %g", tt, got, tt.want)
		}
	}

}

func TestPerimeter(t *testing.T){

	rect := Rectangle{10.0, 20.0}
	got := rect.Perimeter()
	want := 60.0

	if got != want{
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
