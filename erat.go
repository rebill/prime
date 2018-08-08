package prime

func Erat(n int64) bool {
	if n <= 1 {
		return false // 1不是质数，且不考虑负数与0，故输入n<=1时输出为false
	}
	var i int64
	for i = 2; i*i <= n; i++ {
		if n % i == 0 {
			return false // 若整除时输出为false，否则输出为true
		}
	}

	return true
}
