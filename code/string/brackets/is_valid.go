package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "[)"
	fmt.Println(isValid(a))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	ss := strings.Split(s, "")
	isStart := map[string]bool{
		"{": true,
		"[": true,
		"(": true,
	}
	rightEnd := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	starts := []string{}
	for _,v:=range ss{
		if isStart[v]{
			starts = append(starts,v)
			continue
		}
		if len(starts) == 0 || rightEnd[starts[len(starts) - 1]] != v{
			return false
		}
		starts = starts[0:len(starts)-1]
	}
	if len(starts) > 0{
		return false
	}
	return true
}

func isValid2(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	rl := map[byte]byte{
		'[':']',
		'{':'}',
		'(':')',
	}
	stack := []byte{}
	for i := range s{
		if rl[s[i]]>0{
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 || rl[stack[len(stack)-1]] != s[i]{
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
