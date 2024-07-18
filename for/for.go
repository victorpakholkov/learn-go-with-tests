package iteration

func Repeat(char string, repeatCount int) (repeatedString string) {
	for i := 0; i < repeatCount; i++ {
		repeatedString += char
	}
	return repeatedString
}
