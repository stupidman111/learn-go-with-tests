package main

const REPEAT_TIMES = 5

func Repeat(restr string) string {
	res := ""
	for i := 0; i < REPEAT_TIMES; i++ {
		res += restr
	}
	return res
}