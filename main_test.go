package main

import (
	"reflect"
	"testing"
)

type Dot_product_Test struct {
	arg1, arg2, arg3 Coordinate
	expected         float64
}

var Dot_product_Tests = []Dot_product_Test{
	{Coordinate{1, 0}, Coordinate{0, 0}, Coordinate{1, 1}, 0.0},
	{Coordinate{1, 3}, Coordinate{2, 2}, Coordinate{3, 5}, 0.0},
	{Coordinate{1, 1}, Coordinate{3, 2}, Coordinate{3, 1}, 4.0},
	{Coordinate{1, 4}, Coordinate{1, 1}, Coordinate{3, 1}, 9.0},
}

func Test_Dot_product(t *testing.T) {
	for _, test := range Dot_product_Tests {
		if output := Dot_product(test.arg1, test.arg2, test.arg3); output != test.expected {
			t.Errorf("Output %v not equal to expected %v ", output, test.expected)
		}
	}
}

type Unique_items_Test struct {
	arg1, expected []Coordinate
}

var Unique_items_Tests = []Unique_items_Test{
	{[]Coordinate{{0, 0}, {1, 0}, {0, 0}}, []Coordinate{{0, 1}, {0, 0}}},
	{[]Coordinate{{0, 0}, {0, 12}, {0, 12}}, []Coordinate{{0, 12}, {0, 0}}},
	{[]Coordinate{{0, 0}, {1, 2}, {1, 0}}, []Coordinate{{0, 0}, {1, 0}, {1, 2}}},
}

func Test_Unique_items(t *testing.T) {
	for _, test := range Unique_items_Tests {
		if output := Unique_items(test.arg1); reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %v not equal to expected %v ", output, test.expected)
		}
	}
}

type Count_Rectangle_Test struct {
	arg1     []Coordinate
	expected int
}

var Count_Rectangle_Tests = []Count_Rectangle_Test{
	{[]Coordinate{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {4, 0}, {0, 2}, {1, 2}, {6, 2}, {4, 4}, {3, 5}, {1, 3}, {2, 2}, {0, 0}}, 8},
	{[]Coordinate{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {4, 4}, {3, 5}, {1, 3}, {2, 2}, {0, 0}}, 2},
	{[]Coordinate{{0, 0}, {4, 0}, {0, 2}, {1, 2}, {6, 2}, {4, 4}, {3, 5}, {1, 3}, {2, 2}, {0, 0}}, 3},
}

func Test_Count_Rectangle(t *testing.T) {
	for _, test := range Count_Rectangle_Tests {
		if output := Count_Rectangle(test.arg1); output != test.expected {
			t.Errorf("Output %v not equal to expected %v ", output, test.expected)
		}
	}
}
