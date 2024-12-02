package iteration

// Repeats the character, a given count times (equivalent to strings.Repeat,
// but slower, because it uses builder)
func Repeat(character string, count int) string {
	repeated := ""
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
