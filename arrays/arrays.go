package arrays

func Sum(numbers []int) int {

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, val := range numbersToSum {
		sums = append(sums, Sum(val))
	}
	return sums
}

func SumAllTails(tailsToSum ...[]int) []int {
	var tailSums []int
	for _, val := range tailsToSum {
		if len(val) == 0 {
			tailSums = append(tailSums, 0)
		} else {
			tail := val[1:]
			tailSums = append(tailSums, Sum(tail))
		}
	}
	return tailSums
}
