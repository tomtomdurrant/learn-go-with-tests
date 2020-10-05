package arrays


func Sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll(numbersToSum ...[]int) (results []int) {
	var sums []int
	for _,number := range numbersToSum {
		sums = append(sums, Sum(number))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) (results []int) {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) < 1 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}