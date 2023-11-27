package main


func Repeat(reStr string, repeatCount int) string {
	var res string
	for i := 0; i < repeatCount; i++ {
		res += reStr
	}

	return res
}