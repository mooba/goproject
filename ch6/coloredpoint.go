// Copyright 2021 Shopee, Inc.
// author pengchengbai
// date 2021/1/24

package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(p1 Point) float64 {
	return math.Sqrt((p1.X - p.X) * (p1.X - p.X) + (p1.Y - p.Y) * (p1.Y - p.Y))
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"

	var red = color.RGBA{255, 0, 0, 255}
	var cp1 = ColoredPoint{Point{1, 1}, red}

	fmt.Println(cp1.Distance(cp.Point))
}
