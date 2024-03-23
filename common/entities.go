package common;

import "math"

type Point struct {
	X float64
	Y float64
}

func (p1 Point) EuclideanDistance(p2 Point) float64 {
	xDiff := p1.X - p2.X
	yDiff := p1.Y - p2.Y

	sum := (xDiff * xDiff) + (yDiff * yDiff)

	return math.Sqrt(sum)
}

type Cluster struct {
	Center Point
	Points []Point
}

