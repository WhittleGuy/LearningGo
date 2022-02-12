package iteration

const iterations = 5

func Repeat(character string) (repeated string) {
	// var repeated string
	for i := 0; i < iterations; i++ {
		repeated += character
	}
	return
}