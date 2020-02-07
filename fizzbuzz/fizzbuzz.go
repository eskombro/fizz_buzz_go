package fizzbuzz

import (
	"fmt"
	"strconv"
)

func Fizzbuzz(int1 int, int2 int, limit int, str1 string, str2 string) string {
	var ret string
	i := 1
	for i <= limit {
		var result string
		if i%int1 == 0 {
			result = str1
		}
		if i%int2 == 0 {
			result += str2
		}
		if len(result) == 0 {
			result = strconv.Itoa(i)
		}
		i++
		if len(ret) == 0 {
			ret = result
		} else {
			ret = fmt.Sprintf("%s, %s", ret, result)
		}
	}
	return ret
}
