package arrayandslice

func Sum(array [5]int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

// sumAll is a variadic function that can take input of any num of slices of []int type
func sumAll(slices ...[]int) []int {
	// below is a nil slice as slice and map reference to an array and hash table underlying
	// var finalSum []int
	// if we wanna use the above declaration then it will panic as it is a nil
	// have 0 capacity so for this case we need append func
	// finalSum = append(finalSum, sum)
	// slice = append(slice, 1,2,3)
	//slice1 = append(slice1, slice2...) -> for appending 2 slices in a slice
	finalSum := make([]int, len(slices))
	for i, v := range slices {
		sum := 0
		for _, v2 := range v {
			sum += v2
		}
		finalSum[i] = sum
	}
	/*
		we can use logic as well
		for i, v := range slices {
				sliceSplitted := [1:] // this is complete slicing
				add := sum(slicesplitted)
				finalSum = append(finalsum, add)
		}

	*/
	return finalSum
}
