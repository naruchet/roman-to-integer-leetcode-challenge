package main

import "fmt"

func main() {
}

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	var sum int = 0
	for i := 0; i < len(s); i++ {
		str := string(s[i])
		if len(s[i:]) > 1 {
			strP := string(s[i+1])
			var hasValue bool
			var strR string = fmt.Sprintf("%s%s", str, strP)
			if _, hasValue = romanMap[strR]; hasValue {
				sum += romanMap[strR]
				i += 1
				continue
			}
			sum += romanMap[str]
		} else {
			sum += romanMap[str]
		}
	}
	return sum
}
