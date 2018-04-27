package utils

// The Reverse string
func Reverse(str string) string {
	rts := ""
	for i := len(str) - 1; i >= 0; i-- {
		rts += string(str[i])
	}
	return rts
}
