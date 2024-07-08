package iteration

func Repeat(value string, repeat int) string {
	var repeated string
	for i := 0; i < repeat; i++ {
		repeated += value
	}
	return repeated
}
