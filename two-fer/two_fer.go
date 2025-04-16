// The package twofer is useful when you need to share cookies with someone, either if you know their name or if you don't!
package twofer

import "fmt"

// ShareWith returns a formatted string telling who you are sharing a cookie with
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %s, one for me.", name)
	} else {
		return fmt.Sprintf("One for you, one for me.")
	}
}
