package tribo

import "testing"

func TestTribonacci(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"n=0", 0, 0},
		{"n=1", 1, 1},
		{"n=2", 2, 1},
		{"n=3", 3, 2},
		{"n=4", 4, 4},
		{"n=25", 25, 1389537},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := tribonacci(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
func TestMemo(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"n=0", 0, 0},
		{"n=1", 1, 1},
		{"n=2", 2, 1},
		{"n=3", 3, 2},
		{"n=4", 4, 4},
		{"n=25", 25, 1389537},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := memorization(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
