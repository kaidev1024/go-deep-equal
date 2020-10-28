package deepequal

func IsSliceEqual(a, b []int) bool {
	la := len(a)
	lb := len(b)
	if la != lb {
		return false
	}
	for ind, val := range a {
		if val != b[ind] {
			return false
		}
	}
	return true
}
