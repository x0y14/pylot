package pylot

import "strings"

func StringsContain(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

func Indent(nest int) string {
	return strings.Repeat("  ", nest)
}
