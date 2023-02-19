package main

import (
	"fmt"
	"math"
	"sort"
)

type Coordinate struct {
	x float64
	y float64
}

// realization of sort.Inteface
type XthenY []Coordinate

func (a XthenY) Len() int { return len(a) }
func (a XthenY) Less(i, j int) bool {
	if a[i].x == a[j].x {
		return a[i].y < a[j].y
	} else {
		return a[i].x < a[j].x
	}
}
func (a XthenY) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// checks if two floating point numbers are equal within a given error to avoid rounding issues
func Is_equal(a, b, e float64) bool {
	return math.Abs(a-b) < e
}

// computes the dot product of the vectors ab and ac (if result equal 0 -> angle between vectors is 90)
func Dot_product(a, b, c Coordinate) float64 {
	return (b.x-a.x)*(c.x-a.x) + (b.y-a.y)*(c.y-a.y)
}

// checks input array and forming array with unique items
func Unique_items(points []Coordinate) []Coordinate {
	var unique []Coordinate
	m := map[Coordinate]bool{}
	for _, v := range points {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}

// Counting rectangles in a given set of coordinate points
func Count_Rectangle(points []Coordinate) int {
	//unique check
	points = Unique_items(points)
	//sorting array ascending first by X, then by Y
	sort.Sort(XthenY(points))

	rectangles_count := 0
	for a := 0; a < len(points); a++ {
		for b := a + 1; b < len(points); b++ {
			for c := b + 1; c < len(points); c++ {
				for d := c + 1; d < len(points); d++ {
					if Is_equal(Dot_product(points[a], points[b], points[c]), 0.0, 1e-7) &&
						Is_equal(Dot_product(points[b], points[a], points[d]), 0.0, 1e-7) &&
						Is_equal(Dot_product(points[d], points[b], points[c]), 0.0, 1e-7) &&
						Is_equal(Dot_product(points[a], points[b], points[c]), 0.0, 1e-7) {
						rectangles_count += 1
					}
				}
			}
		}
	}
	return rectangles_count
}

// algorithm checks all possible rectangles(not only rectangles that parallel to x and y axis) in O(n^4) time.
func main() {
	var points = []Coordinate{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {4, 0}, {0, 2}, {1, 2}, {6, 2}, {4, 4}, {3, 5}, {1, 3}, {2, 2}, {0, 0}}
	fmt.Println(Count_Rectangle(points))
}
