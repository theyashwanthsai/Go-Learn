package structs



import (
	"testing"
)
func TestPerimeter(t *testing.T){
	rectangle := Rectangle{10.0, 10.0}

	got := rectangle.Perimeter()
	want := 40.0

	if got != want{
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T){

	// checkArea := func(t testing.TB, shape Shape, want float64){
	// 	t.Helper()

	// 	got := shape.Area()
	// 	if got != want{
	// 		t.Errorf("got %g, want %g", got, want)
	// 	}
	// }

	// t.Run("Rectangles", func(t *testing.T){
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	want := 72.0

	// 	checkArea(t, rectangle, want)
	// })

	// t.Run("Circles", func(t *testing.T){
	// 	circle := Circle{10}
	// 	want := 314.1592653589793
	// 	checkArea(t, circle, want)
	// })

	areaTests := []struct{
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle" , shape: Rectangle{12.0, 6.0}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12.0, 6.0}, want: 36.0},
	}
	
	for _, tt := range areaTests{

		t.Run(tt.name, func(t *testing.T){
			got := tt.shape.Area()
			want := tt.want

			if got != want{
				t.Errorf("%v got %g want %g", tt.shape, got, want)
			}
		})
		
	}
}