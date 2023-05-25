package common

func CompareNumber(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
