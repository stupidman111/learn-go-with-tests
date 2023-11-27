package main

func Sum(numbers []int)  int {
	res := 0
	for _, v := range numbers {
		res += v
	}
	return res
}