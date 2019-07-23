package stringutil

import "strings"

func AllIndex(s string, substr string) (indexes []int) {
	index := 0
	for {
		i := strings.Index(s[index:], substr)
		if i == -1 {
			break
		}
		indexes = append(indexes, index+i)
		index += i + 1
	}
	return
}
