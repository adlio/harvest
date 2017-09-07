package harvest

func HaveSameFloat64Value(a, b *float64) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		// Only one is nil
		return false
	}
	return *a == *b
}
