package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"(.*password.*)"`)
	strings := []string{}
	for _,l := range lines {
		if re.MatchString(l) {
			strings = append(strings, l)
		}
	}
	return len(strings)
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	strings := []string{}
	for _,l := range lines {
		if re.FindStringSubmatch(l) != nil {
			submatch := re.FindStringSubmatch(l)[1]
			strings = append(strings, "[USR] " + submatch + " " + l)
		} else {
			strings = append(strings, l)
		}
	}
	return strings
}
