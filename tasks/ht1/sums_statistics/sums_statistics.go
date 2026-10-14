package main

import "fmt"

func SumsStatistics(sums []int) (min, max, sum int, avg float64, ok bool) {
	if len(sums) == 0 {
		return
	}
	min = sums[0]
	max = sums[0]
	for _, v := range sums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
	}
	avg = float64(sum) / float64(len(sums))
	ok = true
	return
}

func main() {
	sums := []int{12000, 7500, 31000, 4500}
	min, max, sum, avg, ok := SumsStatistics(sums)
	if !ok {
		fmt.Println("Ошибка: пустой срез")
	} else {
		fmt.Printf("Min: %d, Max: %d, Sum: %d, Avg: %g\n", min, max, sum, avg)
	}
	sums = []int{}
	min, max, sum, avg, ok = SumsStatistics(sums)
	if !ok {
		fmt.Println("Ошибка: пустой срез")
	} else {
		fmt.Printf("Min: %d, Max: %d, Sum: %d, Avg: %g\n", min, max, sum, avg)
	}
}
