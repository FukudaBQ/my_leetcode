package leetcode43

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := len(num1), len(num2)
	res := make([]int, n1+n2)
	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			sum := mul + res[i+j+1]
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	i := 0
	for i < len(res) && res[i] == 0 {
		i++
	}
	result := ""
	for ; i < len(res); i++ {
		result += string(rune('0' + res[i]))
	}
	if result == "" {
		return "0"
	}
	return result
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := multiply("123", "456")
	fmt.Printf("is anagram: %v", a)
}
