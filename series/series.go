package series

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}
	subStrings := []string{}
	for i:=0; i < len(s); i++ {
		if i+n <= len(s) {
			subStrings = append(subStrings, s[i:i+n])
		}
	}
	return subStrings
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	} else {
		return s[:n], true
	}
}
 
func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return ""
	}
	return s[:n]
}
