// Package proverb is useful to build a rhymed proverb starting from a sequence of words
package proverb

// Proverb returns a rhymed proverb (under a form of a list)
func Proverb(rhyme []string) []string {
	finalStr := []string{}
	for i,_ := range rhyme {
		if i != len(rhyme) - 1 {
			finalStr = append(finalStr, "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost.")
		} else {
			finalStr = append(finalStr, "And all for the want of a " + rhyme[0] + ".")
		}
	}
	return finalStr
}
