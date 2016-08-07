package main

import (
	"testing"
)

func Example_main() {
	main()
	// Output:
	// "{1 2 3}"
}

// newIntSets resturns IntSet pair.
func newIntSets() []IntSet {
	return []IntSet{&MyIntSet{}, NewMapIntSet()}
}

// validateIntSets compares contents of IntSet.
func validateIntSets(t *testing.T, intsets []IntSet) {
	felems, selems := intsets[0].Elems(), intsets[1].Elems()
	for i := 0; i < intsets[0].Len(); i++ {
		if felems[i] != selems[i] {
			t.Errorf("[%d]: felem %d, selem %d", i, felems[i], selems[i])
		}
	}
}

func TestLen(t *testing.T) {
	var tests = []struct {
		words []int
	}{
		{[]int{}},
		{[]int{1, 2, 90}},
		{[]int{1, 2, 3, 4, 200, 6, 7, 8}},
	}
	for _, test := range tests {
		intsets := newIntSets()
		for _, w := range test.words {
			intsets[0].Add(w)
			intsets[1].Add(w)
		}
		if intsets[0].Len() != intsets[1].Len() {
			t.Errorf("Result = %v, Expected %v", intsets[0].Len(), intsets[1].Len())
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		words   []int
		removes []int
	}{
		{[]int{}, []int{1}},
		{[]int{1, 2, 32}, []int{1, 2, 90, 32}},
		{[]int{1, 2, 90}, []int{2}},
		{[]int{1, 2, 3, 4, 200, 6, 7, 8}, []int{200, 4}},
	}

	for _, test := range tests {
		intsets := newIntSets()
		for _, w := range test.words {
			intsets[0].Add(w)
			intsets[1].Add(w)
		}
		for _, r := range test.removes {
			intsets[0].Remove(r)
			intsets[1].Remove(r)
		}

		validateIntSets(t, intsets)

	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		words []int
	}{
		{[]int{1, 2, 3, 4, 200, 6, 7, 8}},
	}

	for _, test := range tests {
		intsets := newIntSets()
		for _, w := range test.words {
			intsets[0].Add(w)
			intsets[1].Add(w)
		}
		intsets[0].Clear()
		intsets[1].Clear()
		validateIntSets(t, intsets)
	}
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		words []int
	}{
		{[]int{}},
		{[]int{1, 2, 90}},
		{[]int{1, 2, 3, 4, 200, 6, 7, 8}},
	}

	for _, test := range tests {
		intsets := newIntSets()
		for _, w := range test.words {
			intsets[0].Add(w)
			intsets[1].Add(w)
		}
		felems := intsets[0].Copy()
		selems := intsets[1].Copy()
		validateIntSets(t, []IntSet{felems, selems})
	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		words []int
	}{
		{[]int{}},
		{[]int{1, 2, 90}},
		{[]int{1, 2, 3, 4, 200, 6, 7, 8}},
	}

	for _, test := range tests {
		intsets := newIntSets()
		for _, w := range test.words {
			intsets[0].Add(w)
			intsets[1].Add(w)
		}
		fs := intsets[0].String()
		ss := intsets[1].String()
		if fs != ss {
			t.Errorf("Error\n%v\n%v", fs, ss)
		}
	}
}

func TestAddAll(t *testing.T) {
	var tests = []struct {
		words  []int
		values []int
	}{
		{[]int{}, []int{1, 30, 400}},
		{[]int{1, 2, 32}, []int{}},
		{[]int{1, 2, 90}, []int{2}},
		{[]int{1, 2, 3, 4, 6, 7, 8}, []int{200}},
	}

	for _, test := range tests {
		intsets := newIntSets()
		for _, w := range test.words {
			intsets[0].Add(w)
			intsets[1].Add(w)
		}
		intsets[0].AddAll(test.values...)
		intsets[1].AddAll(test.values...)

		validateIntSets(t, intsets)

	}
}

func TestIntersectWith(t *testing.T) {
	var tests = []struct {
		xelements []int
		yelements []int
	}{
		{[]int{3}, []int{1, 30, 400}},
		{[]int{1, 2, 32}, []int{1, 2, 32}},
		{[]int{1, 2, 90}, []int{2}},
		{[]int{1, 2, 3, 4, 6, 7, 200}, []int{200, 2}},
	}

	for _, test := range tests {
		yintsets := newIntSets()
		xintsets := newIntSets()
		for _, e := range test.xelements {
			xintsets[0].Add(e)
			xintsets[1].Add(e)
		}
		for _, e := range test.yelements {
			yintsets[0].Add(e)
			yintsets[1].Add(e)
		}
		xintsets[0].IntersectWith(yintsets[0])
		xintsets[1].IntersectWith(yintsets[1])
		validateIntSets(t, xintsets)

	}
}

func TestDifferenceWith(t *testing.T) {
	var tests = []struct {
		xelements []int
		yelements []int
	}{
		{[]int{3}, []int{1, 30, 400}},
		{[]int{1, 2, 32}, []int{1, 2, 32}},
		{[]int{1, 2, 90}, []int{2}},
		{[]int{1, 2, 3, 4, 6, 7, 200}, []int{200, 2}},
	}

	for _, test := range tests {
		yintsets := newIntSets()
		xintsets := newIntSets()
		for _, e := range test.xelements {
			xintsets[0].Add(e)
			xintsets[1].Add(e)
		}
		for _, e := range test.yelements {
			yintsets[0].Add(e)
			yintsets[1].Add(e)
		}
		xintsets[0].DifferenceWith(yintsets[0])
		xintsets[1].DifferenceWith(yintsets[1])
		validateIntSets(t, xintsets)

	}
}

func TestSymmetricDifference(t *testing.T) {
	var tests = []struct {
		xelements []int
		yelements []int
	}{
		{[]int{3}, []int{1, 30, 400}},
		{[]int{1, 2, 32}, []int{1, 2, 32}},
		{[]int{1, 2, 90}, []int{2}},
		{[]int{1, 2, 3, 4, 6, 7, 200}, []int{200, 2, 5}},
	}

	for _, test := range tests {
		yintsets := newIntSets()
		xintsets := newIntSets()
		for _, e := range test.xelements {
			xintsets[0].Add(e)
			xintsets[1].Add(e)
		}
		for _, e := range test.yelements {
			yintsets[0].Add(e)
			yintsets[1].Add(e)
		}
		xintsets[0].SymmetricDifference(yintsets[0])
		xintsets[1].SymmetricDifference(yintsets[1])
		validateIntSets(t, xintsets)

	}
}

func TestUnionWith(t *testing.T) {
	var tests = []struct {
		xelements []int
		yelements []int
	}{
		{[]int{3}, []int{1, 30, 400}},
		{[]int{1, 2, 32}, []int{1, 2, 32}},
		{[]int{1, 2, 90}, []int{2}},
		{[]int{1, 2, 3, 4, 6, 7, 200}, []int{200, 2}},
	}

	for _, test := range tests {
		yintsets := newIntSets()
		xintsets := newIntSets()
		for _, e := range test.xelements {
			xintsets[0].Add(e)
			xintsets[1].Add(e)
		}
		for _, e := range test.yelements {
			yintsets[0].Add(e)
			yintsets[1].Add(e)
		}
		xintsets[0].UnionWith(yintsets[0])
		xintsets[1].UnionWith(yintsets[1])
		validateIntSets(t, xintsets)

	}
}
