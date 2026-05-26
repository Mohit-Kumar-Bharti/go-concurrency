package main

type Area interface {
	CalculateArea() float64
}

type Square struct {
	Side float64
}

func (s Square) CalculateArea() float64 {
	return s.Side * s.Side
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) CalculateArea() float64 {
	return r.Length * r.Width
}

func PrintArea(a Area) {
	area := a.CalculateArea()
	println("Area:", area)
}

func main() {
	square := Square{Side: 5}
	rectangle := Rectangle{Length: 4, Width: 6}
	PrintArea(square)
	PrintArea(rectangle)
}
