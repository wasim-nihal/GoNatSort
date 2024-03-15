// Copyright 2018 Syed Wasim Nihal (wasim-nihal). All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.
package gonatsort

type NatSorter []string

func (p NatSorter) Len() int           { return len(p) }
func (p NatSorter) Less(i, j int) bool { return Less(p[i], p[j]) }
func (p NatSorter) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func Less(a, b string) bool {
	for len(a) > 0 && len(b) > 0 {
		// get the length of common prefix
		p := lengthOfCommonPrefix(a, b)
		// If there is a common prefix, remove the prefix from both the strings
		if p > 0 {
			a = a[p:]
			b = b[p:]
		}
		if len(a) == 0 {
			return len(b) != 0
		}
		ia := firstNonNumericCharacterIndex(a)
		ib := firstNonNumericCharacterIndex(b)
		switch {
		// If both strings a and b, begin with a numeric character
		case ia > 0 && ib > 0:
			// remove leading zeros if any
			trimmedA, lenTrimmedA := removeLeadingZeros(a[:ia])
			trimmedB, lenTrimmedB := removeLeadingZeros(b[:ib])
			// check the the numeric weightage based on number of digits in the numeric substring
			if lenTrimmedA > lenTrimmedB {
				return false
			} else if lenTrimmedA < lenTrimmedB {
				return true
			}
			// if the length is same, compare them lexicographically
			if trimmedA != trimmedB {
				return trimmedA < trimmedB
			}
			// if the length is same and the value is also same, remove the substring from both the strings
			// and continue with next set of characters
			if ia != len(a) && ib != len(b) {
				a = a[ia:]
				b = b[ib:]
				continue
			}
		// all the other cases, compare them lexicographically
		default:
			return a < b
		}
	}
	return a < b
}

// lengthOfCommonPrefix, returns a length of common prefix between the given two strings a and b.
// Returns 0 if there is no common prefix
func lengthOfCommonPrefix(a, b string) int {
	i, j := 0, 0
	lenA, lenB := len(a), len(b)
	for i < lenA && j < lenB {
		if a[i] == b[j] {
			i++
			j++
		} else {
			break
		}
	}
	return i
}

// firstNonNumericCharacterIndex returns the index of first non-numeric character
func firstNonNumericCharacterIndex(s string) int {
	for i, c := range s {
		if c < '0' || c > '9' {
			return i
		}
	}
	return len(s)
}

// removeLeadingZeros removes all the leading zeros from the numeric string, and returns it with the new length
func removeLeadingZeros(s string) (string, int) {
	// if the numeric string does not begin with 0, return the same string with its length
	if s[0] != '0' {
		return s, len(s)
	}
	index := 0
	for index < len(s) && s[index] == '0' {
		index++
	}
	return s[index:], len(s[index:])
}
