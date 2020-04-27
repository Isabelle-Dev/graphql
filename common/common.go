package common

// Exists is a utility func which returns true or false if a string slice contains a
// specified string element
func Exists(element string, container []string) bool {
	for _, i := range container {
		if i == element {
			return true
		}
	}
	return false
}
