package iteration

import "strings"

func Repeat(value string, repeat int) string {
	value = strings.TrimSpace(value)
	var repeated string
	for i := 0; i < repeat; i++ {
		repeated += value
	}
	return repeated
}
