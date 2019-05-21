// Package LCP provides data structure which allows effectively calculate
// longest common prefix of two any suffixes of specified string.
package LCP

import (
	"math"
)

type LongestCommonPrefix interface {

	// Get the longest common prefix length for two specified suffixes of the string
	//
	// Arguments:
	//   firstSuffixIndex - position for the first suffix
	//   secondSuffixIndex - position for the second suffix
	//
	// Return value: position of the longest common prefix for specified suffixes
	//
	// Complexity: O(logN), where N is the string length
	Get(firstSuffixIndex uint64, secondSuffixIndex uint64) uint64
}

// Create longest common prefix data structure for the specified string
//
// Arguments:
//   string - the string, for suffixes of which longest common prefix will be calculated
//
// Complexity: O(N * logN), where N is the string length
func NewLongestCommonPrefix(str string) LongestCommonPrefix {
	return &lcpImpl{}
}

// Implementation

const (
	charMin    = 0
	charMax    = math.MaxUint8
	charsCount = charMax - charMin + 1
)

type lcpImpl struct {
}

func (lcp *lcpImpl) Get(firstPrefixIndex uint64, secondPrefixIndex uint64) uint64 {
	return 0
}

type SuffixArray []uint64
type EquivClasses []uint8

func makeSuffixArray(str string) (SuffixArray, EquivClasses) {
	strLen := uint64(len(str))

	sufArr := make(SuffixArray, strLen)
	eqCl := make(EquivClasses, strLen)

	if strLen == 0 {
		return sufArr, eqCl
	}

	classesCount := initSuffixArray(str, &sufArr, &eqCl)

	sortedSufArr := make(SuffixArray, strLen)
	eqClTemp := make(EquivClasses, strLen)

	sortingTable := [charsCount]uint64{}

	// O(N * logN)
	for oldSubStrLen := uint64(1); oldSubStrLen < strLen; oldSubStrLen *= 2 {
		for i := range sortedSufArr {
			if sufArr[i] < oldSubStrLen {
				sortedSufArr[i] = strLen
			} else {
				sortedSufArr[i] = 0
			}

			sortedSufArr[i] += sufArr[i] - oldSubStrLen
		}

		sortingTable := sortingTable[:classesCount]

		for i := range sortingTable {
			sortingTable[i] = 0
		}
		for i := range sortedSufArr {
			sortingTable[eqCl[sortedSufArr[i]]]++
		}

		// O(1)
		for i := 1; i < len(sortingTable); i++ {
			sortingTable[i] += sortingTable[i-1]
		}

		for i := strLen; i > 0; {
			i--
			tableIndex := eqCl[sortedSufArr[i]]
			sortingTable[tableIndex]--
			sufArr[sortingTable[tableIndex]] = sortedSufArr[i]
		}

		eqClTemp[sufArr[0]] = 0
		classesCount = 1

		for i := 1; i < len(sufArr); i++ {
			mid1, mid2 := (sufArr[i]+oldSubStrLen)%strLen, (sufArr[i-1]+oldSubStrLen)%strLen

			if (eqCl[sufArr[i]] != eqCl[sufArr[i-1]]) || (eqCl[mid1] != eqCl[mid2]) {
				classesCount++
			}

			eqClTemp[sufArr[i]] = classesCount - 1
		}

		eqCl, eqClTemp = eqClTemp, eqCl
	}

	return sufArr, eqCl
}

// Complexity: O(N)
func initSuffixArray(str string, sufArr *SuffixArray, eqCl *EquivClasses) uint8 {
	sortingTable := [charsCount]uint64{}

	// O(N)
	for _, ch := range str {
		sortingTable[ch]++
	}

	// O(1)
	for i := 1; i < len(sortingTable); i++ {
		sortingTable[i] += sortingTable[i-1]
	}

	// O(N)
	for i, ch := range str {
		sortingTable[ch]--
		(*sufArr)[sortingTable[ch]] = uint64(i)
	}

	(*eqCl)[(*sufArr)[0]] = 0
	classesCount := uint8(1)

	// O(N)
	for i := 1; i < len(*sufArr); i++ {
		if str[(*sufArr)[i]] != str[(*sufArr)[i-1]] {
			classesCount++
		}

		(*eqCl)[(*sufArr)[i]] = classesCount - 1
	}

	return classesCount
}

func MakeSuffixArray(str string) (SuffixArray, EquivClasses) {
	return makeSuffixArray(str)
}
