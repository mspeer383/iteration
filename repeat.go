package iteration

//Repeat repeates a given string and returns a string
func Repeat(c string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += c
	}
	return repeated
}
