package iterations

func Repeat(char string, repeatNo int) string {
	var repeated string
	for i := 0; i < repeatNo; i++ {
		repeated += char
	}
	return repeated
}
