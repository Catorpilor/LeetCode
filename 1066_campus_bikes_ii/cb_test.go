package cb

import "testing"

func TestAssignBikes(t *testing.T) {
	st := []struct {
		name           string
		workers, bikes [][]int
		exp            int
	}{
		{"empty workers", [][]int{}, [][]int{[]int{1, 2}}, 0},
		{"empty bikes", [][]int{[]int{0, 0}}, [][]int{}, 0},
		{"testcase1", [][]int{[]int{0, 0}, []int{2, 1}}, [][]int{[]int{1, 2}, []int{3, 3}}, 6},
		{"testcase2", [][]int{[]int{0, 0}, []int{1, 1}, []int{2, 0}}, [][]int{[]int{1, 0}, []int{2, 2}, []int{2, 1}}, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := assignBikes(tt.workers, tt.bikes)
			if out != tt.exp {
				t.Fatalf("with workers: %v and bikes %v wanted %d but got %d", tt.workers, tt.bikes, tt.exp, out)
			}
			t.Log("pass")
		})
	}

}

func TestUseDp(t *testing.T) {
	st := []struct {
		name           string
		workers, bikes [][]int
		exp            int
	}{
		{"empty workers", [][]int{}, [][]int{[]int{1, 2}}, 0},
		{"empty bikes", [][]int{[]int{0, 0}}, [][]int{}, 0},
		{"testcase1", [][]int{[]int{0, 0}, []int{2, 1}}, [][]int{[]int{1, 2}, []int{3, 3}}, 6},
		{"testcase2", [][]int{[]int{0, 0}, []int{1, 1}, []int{2, 0}}, [][]int{[]int{1, 0}, []int{2, 2}, []int{2, 1}}, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useDp(tt.workers, tt.bikes)
			if out != tt.exp {
				t.Fatalf("with workers: %v and bikes %v wanted %d but got %d", tt.workers, tt.bikes, tt.exp, out)
			}
			t.Log("pass")
		})
	}

}

func TestNumOfOnes(t *testing.T) {
	st := []struct {
		name string
		i    int
		exp  int
	}{
		{"i eq 0", 0, 0},
		{"i eq 1", 1, 1},
		{"i eq 3", 3, 2},
		{"i eq 7", 7, 3},
		{"i eq 8", 8, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := numOfOnes(tt.i)
			if out != tt.exp {
				t.Fatalf("with i:%d waned %d but got %d", tt.i, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
