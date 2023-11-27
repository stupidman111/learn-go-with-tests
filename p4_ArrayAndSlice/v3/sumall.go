package main

//可变参数函数-参数个数可变的函数
func SumAll(numbersToSum...[]int) (sums []int) {
	//argsLen := len(numbersToSum)
	//sums = make([]int, argsLen)

	for _, numbers := range numbersToSum {
		//sums[i] = Sum(numbers)
		sums = append(sums, Sum(numbers))
	}
	return
}

func Sum(numbers []int) int {
	res := 0

	for _, v := range numbers {
		res += v
	}
	return res
}