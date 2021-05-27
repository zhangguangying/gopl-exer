package geometry

import "math"

type Point struct {
	X, Y float64
}

func Distince(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distince(q Point) float64 {
	return math.Hypot(q.X-p.X,q.Y-q.Y)
}

type Path []Point

func (p Path) Distince() float64 {
	sum := 0.0
	for i := range p {
		sum += p[i-1].Distince(p[i])
	}
	return sum
}