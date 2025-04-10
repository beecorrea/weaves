package mergesort

import (
	"testing"

	"github.com/beecorrea/orders/pkg/order"
)

func TestMergeSort(t *testing.T) {
	strategy := Mergesort{}
	o := order.Leq{}
	numbers := order.Random(o)
	order.AssertPartiallyOrdered(t)(strategy, numbers)
}
