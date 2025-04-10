package order

type Lt struct{}

func (lt Lt) Compare(a, b int) bool {
	return a < b
}

func (lt Lt) Reflexivity(a int) bool {
	return lt.Compare(a, a)
}

func (lt Lt) Antisymmetry(a, b int) bool {
	if lt.Compare(a, b) && lt.Compare(b, a) && a != b {
		return false
	}
	return true
}

func (lt Lt) Transitivity(a, b, c int) bool {
	if lt.Compare(a, b) && lt.Compare(b, c) && !lt.Compare(a, c) {
		return false
	}

	return true
}
