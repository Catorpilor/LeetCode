package kth

import "testing"

func constructWithTesting(t *testing.T) *TreeAncestor {
	ta := Construct(6, []int{-1, 0, 0, 1, 1, 2, 2})
	t.Cleanup(func() {
		ta = nil
	})
	return ta
}

func TestKthAncestor(t *testing.T) {
	ta := constructWithTesting(t)
	st := []struct {
		name         string
		node, k, exp int
	}{}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := ta.GetKthAncestor(tt.node, tt.k)
			if out != tt.exp {
				t.Fatalf("with input node:%d and k:%d wanted %d but got %d", tt.node, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
