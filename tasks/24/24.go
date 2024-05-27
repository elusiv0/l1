package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) dist(to *Point) float64 {
	dx := p.x - to.x
	dy := p.y - to.y
	dist := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))

	return dist
}

func main() {
	firstPoint := &Point{
		x: -4,
		y: -2,
	}
	secondPoint := &Point{
		x: -2,
		y: -1,
	}

	fmt.Printf("distance equal %.3f", firstPoint.dist(secondPoint))
}
