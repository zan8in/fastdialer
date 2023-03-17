package stringsutil

import "strings"

// EqualFoldAny returns true if s is equal to any specified substring
func EqualFoldAny(s string, ss ...string) bool {
	for _, sss := range ss {
		if strings.EqualFold(s, sss) {
			return true
		}
	}
	return false
}

// ContainsAny returns true is s contains any specified substring
func ContainsAny(s string, ss ...string) bool {
	for _, sss := range ss {
		if strings.Contains(s, sss) {
			return true
		}
	}
	return false
}
