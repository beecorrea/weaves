package order

type Leq struct{}

func (leq Leq) Compare(a, b int) bool {
	return a <= b
}

func (leq Leq) Reflexivity(a int) bool {
	return leq.Compare(a, a)
}

func (leq Leq) Antisymmetry(a, b int) bool {
	if leq.Compare(a, b) && leq.Compare(b, a) && a != b {
		return false
	}
	return true
}

func (leq Leq) Transitivity(a, b, c int) bool {
	if leq.Compare(a, b) && leq.Compare(b, c) && !(a <= c) {
		return false
	}

	return true
}
