//https://www.codewars.com/kata/5886e082a836a691340000c3/train/go

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(RectangleRotation(6, 4))
	fmt.Println(RectangleRotation(30, 2))
	fmt.Println(RectangleRotation(8, 6))
	fmt.Println(RectangleRotation(16, 20))
	// fmt.Println(RectangleRotation(10000, 10000))
}

type Point struct {
	X, Y float64
}

type SlopeInterceptLine struct {
	Slope, Intercept float64
}

func (si SlopeInterceptLine) LineIntercept(line SlopeInterceptLine) Point {
	x := (line.Intercept - si.Intercept) / (si.Slope - line.Slope)
	return Point{
		X: x,
		Y: (si.Slope * x) + si.Intercept,
	}
}

type Rectangle struct {
	P1, P2, P3, P4 Point
	L1, L2, L3, L4 SlopeInterceptLine
}

func NewRectangleFromPoints(a, b int) *Rectangle {
	return &Rectangle{
		P1: Point{
			X: float64(a / 2 * -1),
			Y: float64(b / 2),
		},
		P2: Point{
			X: float64(a / 2),
			Y: float64(b / 2),
		},
		P3: Point{
			X: float64(a / 2),
			Y: float64(b / 2 * -1),
		},
		P4: Point{
			X: float64(a / 2 * -1),
			Y: float64(b / 2 * -1),
		},
	}
}

// Rotate - Will return a new Rectangle object that has been rotated
// by the parameter degrees
func (r *Rectangle) Rotate(degrees float64) (r2 *Rectangle) {
	radians := degrees * (math.Pi / 180)
	return &Rectangle{
		P1: Point{
			X: (r.P1.X * math.Cos(radians)) - (r.P1.Y * math.Sin(radians)),
			Y: (r.P1.X * math.Sin(radians)) + (r.P1.Y * math.Cos(radians)),
		},
		P2: Point{
			X: (r.P2.X * math.Cos(radians)) - (r.P2.Y * math.Sin(radians)),
			Y: (r.P2.X * math.Sin(radians)) + (r.P2.Y * math.Cos(radians)),
		},
		P3: Point{
			X: (r.P3.X * math.Cos(radians)) - (r.P3.Y * math.Sin(radians)),
			Y: (r.P3.X * math.Sin(radians)) + (r.P3.Y * math.Cos(radians)),
		},
		P4: Point{
			X: (r.P4.X * math.Cos(radians)) - (r.P4.Y * math.Sin(radians)),
			Y: (r.P4.X * math.Sin(radians)) + (r.P4.Y * math.Cos(radians)),
		},
	}
}

func (r *Rectangle) CalculateLines() *Rectangle {
	m := (r.P1.Y - r.P2.Y) / (r.P1.X - r.P2.X)
	r.L1 = SlopeInterceptLine{
		Slope:     m,
		Intercept: r.P1.Y - (r.P1.X * m),
	}
	m = (r.P2.Y - r.P3.Y) / (r.P2.X - r.P3.X)
	r.L2 = SlopeInterceptLine{
		Slope:     m,
		Intercept: r.P2.Y - (r.P2.X * m),
	}
	m = (r.P3.Y - r.P4.Y) / (r.P3.X - r.P4.X)
	r.L3 = SlopeInterceptLine{
		Slope:     m,
		Intercept: r.P3.Y - (r.P3.X * m),
	}
	m = (r.P4.Y - r.P1.Y) / (r.P4.X - r.P1.X)
	r.L4 = SlopeInterceptLine{
		Slope:     m,
		Intercept: r.P4.Y - (r.P4.X * m),
	}
	return r
}

func RectangleRotation(a, b int) int {
	rec := NewRectangleFromPoints(a, b).Rotate(45).CalculateLines()

	var lBound, rBound Point
	var pointCount float64
	l := SlopeInterceptLine{Slope: rec.L1.Slope}
	for i := int(rec.L1.Intercept); i >= int(rec.L3.Intercept); i-- {
		l.Intercept = float64(i)

		lBound = l.LineIntercept(rec.L4)
		rBound = l.LineIntercept(rec.L2)
		pointCount += math.Floor(rBound.X) - math.Floor(lBound.X)
	}

	return int(pointCount)
}
