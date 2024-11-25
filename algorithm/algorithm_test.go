package algorithm_test

import (
	"testing"

	"github.com/gox-studio/gona/algorithm"
)

func TestRandPR(t *testing.T) {
	c := 0
	for i := 1; i < 1000; i++ {
		value := algorithm.RandPR(0.01)
		if value {
			c++
		}
	}

	t.Log("RandPR:", c)
}

func TestFastFindIndex(t *testing.T) {
	var ns = []float64{0, 2.0, 2.1, 3.8, 4.9, 6.0, 6.222222}
	var s = []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}
	for index, n := range ns {
		result := algorithm.FastFindIndex(n, s)
		if result != index {
			t.Error("fast find index error", index, result)
		}
	}

	ns = []float64{0, 2.0, 2.1, 3.8, 4.9, 6.0}
	s = []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}
	for index, n := range ns {
		result := algorithm.FastFindIndex(n, s)
		if result != index {
			t.Error("fast find index error", index, result)
		}
	}
}
