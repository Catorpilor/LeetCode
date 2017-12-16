package stock

import "testing"

func TestStock(t *testing.T) {
	st := []struct {
		name   string
		prices []int
		k, exp int
	}{
		{"empty", []int{}, 1, 0},
		{"single", []int{}, 1, 0},
		{"two decreasing", []int{2, 1}, 1, 0},
		{"test 1", []int{7, 1, 5, 3, 6, 4}, 2, 7},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Stock(c.prices, c.k)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v and %d",
					c.exp, ret, c.prices, c.k)
			}
		})
	}
}

func TestStock2(t *testing.T) {
	st := []struct {
		name   string
		prices []int
		k, exp int
	}{
		{"empty", []int{}, 1, 0},
		{"single", []int{}, 1, 0},
		{"two decreasing", []int{2, 1}, 1, 0},
		{"test 1", []int{7, 1, 5, 3, 6, 4}, 2, 7},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Stock2(c.prices, c.k)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v and %d",
					c.exp, ret, c.prices, c.k)
			}
		})
	}
}
