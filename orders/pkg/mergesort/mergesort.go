package mergesort

import (
	"github.com/beecorrea/orders/pkg/order"
)

type Mergesort struct {
}

func sortAndMerge(xs []int, ys []int, o order.Relation) []int {
	i := 0
	j := 0
	out := make([]int, 0)

	for i < len(xs) && j < len(ys) {
		if o.Compare(xs[i], ys[j]) {
			out = append(out, xs[i])
			i++
		} else {
			out = append(out, ys[j])
			j++
		}
	}

	for i < len(xs) {
		out = append(out, xs[i])
		i++
	}

	for j < len(ys) {
		out = append(out, ys[j])
		j++
	}

	return out
}

func mergesort(xs []int, o order.Relation) []int {
	if len(xs) < 2 {
		return xs
	}
	// Split in half
	left := mergesort(xs[:len(xs)/2], o)
	right := mergesort(xs[len(xs)/2:], o)
	// Sort subarrays
	return sortAndMerge(left, right, o)
}

func (ms Mergesort) Strategy() string {
	return "Mergesort"
}

func (ms Mergesort) Run(ps order.Poset) order.Poset {
	numbers := ps.Members()
	sorted := mergesort(numbers, ps.Order())
	return order.New(sorted, ps.Order())
}
