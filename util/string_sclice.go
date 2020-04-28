package util

import "regexp"

// Filter removes all matching elements from a string slice based on a regexp selector
func Filter(selector *regexp.Regexp, slice []string) (out []string, numRemoved int) {
	out, removed := FilterSplit(selector, slice)
	return out, len(removed)
}

// FilterSplit removes all matching elements from a string slice based on a regexp selector and returns all the removed items too
func FilterSplit(selector *regexp.Regexp, slice []string) (out, removed []string) {
	out = []string{}
	removed = []string{}
	for _, elem := range slice {
		if selector.MatchString(elem) {
			removed = append(removed, elem)
		} else {
			out = append(out, elem)
		}
	}
	return
}

// SelectMatching return all matching elements from a string slice based on a regexp selector
func SelectMatching(selector *regexp.Regexp, slice []string) (matching []string) {
	_, matching = FilterSplit(selector, slice)
	return
}
