package rabn

import (
	"math"
)

func commonPrefix(s1, s2 string) string {
	minLength := int(math.Min(float64(len(s1)), float64(len(s2))))

	for i := 0; i < minLength; i++ {
		if s1[i] != s2[i] {
			return s1[:i]
		}
	}
	return s1[:minLength]
}

func longestCommonPrefix(stringList []string, leftIndex, rightIndex int) string {
	if leftIndex == rightIndex {
		return stringList[leftIndex]
	}

	midIndex := (leftIndex + rightIndex) / 2
	longestCommonPrefixLeft := longestCommonPrefix(stringList, leftIndex, midIndex)
	longestCommonPrefixRight := longestCommonPrefix(stringList, midIndex+1, rightIndex)
	return commonPrefix(longestCommonPrefixLeft, longestCommonPrefixRight)
}

func FindLongestCommonPrefix(stringList []string) string {
	return longestCommonPrefix(stringList, 0, len(stringList) - 1)
}