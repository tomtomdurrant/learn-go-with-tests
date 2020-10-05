package iteration

func Repeat(letter string, max int) string {
	var repeated string
	for i := 0; i < max; i++ {
		repeated += letter
	}

	return repeated
}