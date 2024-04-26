package strings

func PrintOutStirng(arrayOfStrings []string) string {
	var s string
	for _, t := range arrayOfStrings {
		s = s + " " + t
	}
	return s
}
