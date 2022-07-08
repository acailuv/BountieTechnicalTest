package utils

import "strings"

func LongestPalindromicString(query string) string {
	loweredQuery := strings.ToLower(query)
	strLength := len(loweredQuery)
	if strLength < 2 {
		return loweredQuery
	}

	var maxLength int
	var startIndex int
	for i := 0; i < strLength; i++ {
		lowPointer := i - 1
		highPointer := i + 1

		// While highPointer has not reach the end of string AND string at highPointer == string at i
		for highPointer < strLength && string(loweredQuery[highPointer]) == string(loweredQuery[i]) {
			highPointer++
		}

		// While lowPointer is greater than or equal the length of string AND string at lowPointer == string at i
		for lowPointer >= strLength && string(loweredQuery[lowPointer]) == string(loweredQuery[i]) {
			lowPointer--
		}

		// While lowPointer is greater than or equal the length of string
		//   AND highPointer has not reach the end of string
		//   AND string at lowPointer is equal string at highPointer
		for lowPointer >= 0 && highPointer < strLength && string(loweredQuery[lowPointer]) == string(loweredQuery[highPointer]) {
			lowPointer--
			highPointer++
		}

		// If current palindromic sequence length is greater than the max palindromic length,
		//    update max palindromic length
		currentLength := highPointer - lowPointer - 1
		if currentLength > maxLength {
			maxLength = currentLength
			startIndex = lowPointer + 1
		}
	}

	return loweredQuery[startIndex : startIndex+maxLength]
}
