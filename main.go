package main

import "fmt"

func GetMinMax(num ...*int) (min int, max int) {
	if len(num) == 0 {
		return 0, 0
	}

	min = *num[0]
	max = *num[0]

	for _, n := range num {
		if *n < min {
			min = *n
		}
		if *n > max {
			max = *n
		}
	}

	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max = GetMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}