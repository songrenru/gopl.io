func plusOne(digits []int) []int {
	lens := len(digits)
	for i := lens - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	res := make([]int, lens+1)
	res[0] = 1
	return res
}