package order

import (
	"fmt"
	"testing"
)

type Relation interface {
	Compare(a, b int) bool
	Reflexivity(a int) bool
	Antisymmetry(a, b int) bool
	Transitivity(a, b, c int) bool
}

type Sort interface {
	Strategy() string
	Run(ps Poset) Poset
}

func AssertPartiallyOrdered(t *testing.T) func(strategy Sort, poset Poset) {
	return func(strategy Sort, poset Poset) {
		testName := fmt.Sprintf("Should sort using %s", strategy.Strategy())

		t.Run(testName, func(t *testing.T) {
			actual := poset.Sort(strategy)
			if len(actual.Members()) != len(poset.Members()) {
				t.Errorf("actual and original have different amount of elements")
			}

			t.Run("Assert properties of partial order relations", func(t *testing.T) {
				isPartiallyOrdered := poset.IsPartiallyOrdered()
				if !isPartiallyOrdered {
					t.Errorf("not an poset")
				}
			})
		})
	}
}
