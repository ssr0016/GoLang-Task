package main

import "fmt"

func task(input int) (int, error) {
	var result int

	factors := []int{3, 5}
	multiples := make(map[int]bool)

	for _, factor := range factors {
		for i := factor; i < input; i += factor {
			if !multiples[i] {
				multiples[i] = true
				result = result + i
			}
		}
	}

	return result, nil
}

func main() {

	result, err := task(20)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
