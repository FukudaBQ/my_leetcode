package leetcode17

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := make([]string, 0)
	path := make([]byte, 0)

	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(digits) {
			res = append(res, string(path))
			return
		}

		letters := phone[digits[index]]

		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])
			backtrack(index + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := letterCombinations("42")
	fmt.Printf("water is: %v", a)
}
