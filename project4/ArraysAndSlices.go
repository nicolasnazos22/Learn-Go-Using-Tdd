package main

func average(list []int64) float64 {
	return float64(sum(list)) / float64(len(list))

}

func sum(list []int64) int64 {

	var acc int64 = 0
	for _, number := range list {
		acc += number
	}
	return acc
}
func withoutRepeatedElements(list []int64) []int64 {
	var elements []int64
	for _, element := range list {
		if cont(element, elements) == false {
			elements = append(elements, element)
		}
	}
	return elements
}
func cont(element int64, list []int64) bool {
	for _, number := range list {
		if number == element {
			return true
		}
	}
	return false
}
