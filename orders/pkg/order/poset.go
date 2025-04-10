package order

import (
	"github.com/beecorrea/orders/pkg/fake"
)

type poset struct {
	members []int
	order   Relation
}

type Poset interface {
	Members() []int
	Order() Relation
	IsPartiallyOrdered() bool
	Sort(s Sort) Poset
}

func New(xs []int, order Relation) Poset {
	return poset{
		members: xs,
		order:   order,
	}
}

func Random(order Relation) Poset {
	xs := fake.RandomInts(100)
	return New(xs, order)
}

func (ps poset) Members() []int {
	return ps.members
}

func (ps poset) Order() Relation {
	return ps.order
}

func (ps poset) IsPartiallyOrdered() bool {
	var reflexivity, antisymmetry, transitivity bool
	for i := range ps.members {
		reflexivity = ps.order.Reflexivity(ps.members[i])
		for j := range ps.members {
			antisymmetry = ps.order.Antisymmetry(ps.members[i], ps.members[j])
			for k := range ps.members {
				transitivity = ps.order.Transitivity(ps.members[i], ps.members[j], ps.members[k])
				if !(reflexivity && antisymmetry && transitivity) {
					return false
				}
			}
		}
	}

	return true
}

func (ps poset) Sort(s Sort) Poset {
	return s.Run(ps)
}
