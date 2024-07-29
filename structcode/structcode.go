package structcode

type Rectangle struct {
	Height, Width float32
}

type Circle struct {
	Radius float32
}

func perimeter(rect Rectangle) float32 {
	return 2 * (rect.Height + rect.Width)
}
func area(rect Rectangle) float32 {
	return rect.Height * rect.Width
}
func (c Circle) area() float32 {
	return 2 * 3.14 * c.Radius
}
