package fizzbuzz

import (
	"fmt"
	"strconv"
)

type FizzBuzzParams struct {
	Int1  int    `json:"int1" form:"int1" param:"int1"`
	Int2  int    `json:"int2" form:"int2" param:"int2"`
	Limit int    `json:"limit" form:"limit" param:"limit"`
	Str1  string `json:"str1" form:"str1" param:"str1"`
	Str2  string `json:"str2" form:"str2" param:"str2"`
}

func Fizzbuzz(params FizzBuzzParams) string {
	var ret string
	if params.Int1 <= 0 || params.Int2 <= 0 || params.Limit <= 0 {
		return ret
	}
	if len(params.Str1) == 0 || len(params.Str2) == 0 {
		return ret
	}
	for i := 1; i <= params.Limit; i++ {
		var result string
		if i%params.Int1 == 0 {
			result = params.Str1
		}
		if i%params.Int2 == 0 {
			result += params.Str2
		}
		if len(result) == 0 {
			result = strconv.Itoa(i)
		}
		if len(ret) == 0 {
			ret = result
		} else {
			ret = fmt.Sprintf("%s, %s", ret, result)
		}
	}
	AddRequest(params)
	return ret
}
