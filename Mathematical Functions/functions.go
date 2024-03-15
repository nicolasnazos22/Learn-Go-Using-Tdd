package main

func recursiveFactorial(num int64) int64 {
	if num < 0 {
		return -1
	}
	if num == 0 || num == 1 {
		return 1
	} else {
		return num * recursiveFactorial(num-1)
	}
}
func iterativeFactorial(num int64) int64 {
	if num < 0 {
		return -1
	}
	if num == 0 || num == 1 {
		return 1
	}
	i := num
	var j int64 = 1

	for i >= 1 {
		j *= i
		i--
	}
	return j

}
func eAproximation(num int64) float64 {
	if num < 0 {
		return -1
	} else if num == 0 {
		return 2
	} else {
		return 1 + 1/float64(recursiveFactorial(num))
	}

}
