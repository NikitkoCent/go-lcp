// Package lcp provides data structure which allows effectively calculate
// longest common prefix of two any suffixes of specified string.
package lcp

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
	Get(firstSuffixIndex uint, secondSuffixIndex uint) uint
}

// Create longest common prefix data structure for the specified string
//
// Arguments:
//   string - the string, for suffixes of which longest common prefix will be calculated
//
// Complexity: O(N * logN), where N is the string length
func NewLongestCommonPrefix(str string) LongestCommonPrefix {
	result := lcpImpl{}

	result.initialize(str)

	return &result
}

// Implementation

const (
	charMin    = 0
	charMax    = math.MaxUint8
	charsCount = charMax - charMin + 1
)

type suffixArray []uint
type equivClasses []uint

type lcpImpl struct {
	len uint
	allEquivClasses []equivClasses
}

func (lcp *lcpImpl) Get(firstSuffixIndex uint, secondSuffixIndex uint) uint {
	if (firstSuffixIndex >= lcp.len) || (secondSuffixIndex >= lcp.len) {
		panic("Index out of bounds")
	}

	if firstSuffixIndex == secondSuffixIndex {
		return lcp.len - firstSuffixIndex
	}

	result := uint(0)

	i := uint(len(lcp.allEquivClasses))
	subStrLen := uint(1) << i
	for ; i > 0; {
		i--
		subStrLen >>= 1

		if lcp.allEquivClasses[i][firstSuffixIndex] == lcp.allEquivClasses[i][secondSuffixIndex] {
			firstSuffixIndex += subStrLen
			secondSuffixIndex += subStrLen

			result += subStrLen

			if (firstSuffixIndex >= lcp.len) || (secondSuffixIndex >= lcp.len) {
				break
			}
		}
	}

	return result
}


// Builds suffix array and equivalence classes
//
// Complexity: O(N * log N)
func (lcp *lcpImpl) initialize(str string) {
	lcp.len = 0
	lcp.allEquivClasses = nil

	if len(str) < 1 {
		return
	}

	lcp.len = uint(len(str))

	if lcp.len == 1 {
		lcp.allEquivClasses = []equivClasses{{0}}
		return
	}

	lcp.allEquivClasses = make([]equivClasses, math.Ilogb(float64(lcp.len)) + 1)
	for i := range lcp.allEquivClasses {
		lcp.allEquivClasses[i] = make(equivClasses, lcp.len)
	}

	sufArr := make(suffixArray, lcp.len)
	sortingTable := make([]uint, lcp.len)

	classesCount := initSuffixArray(str, sufArr, lcp.allEquivClasses[0])

	sortedSufArr := make(suffixArray, lcp.len)

	for newEqClIndex, oldSubStrLen, newSubStrLen := uint(1), uint(1), uint(2);
		newSubStrLen < lcp.len;
		newEqClIndex, oldSubStrLen, newSubStrLen = newEqClIndex + 1, newSubStrLen, newSubStrLen * 2 {
		for i := range sortedSufArr {
			if sufArr[i] < oldSubStrLen {
				sortedSufArr[i] = lcp.len
			} else {
				sortedSufArr[i] = 0
			}

			sortedSufArr[i] += sufArr[i] - oldSubStrLen
		}

		sortingTable := sortingTable[:classesCount]
		oldEqClasses := lcp.allEquivClasses[newEqClIndex - 1]

		for i := range sortingTable {
			sortingTable[i] = 0
		}
		for i := range sortedSufArr {
			sortingTable[oldEqClasses[sortedSufArr[i]]]++
		}

		for i := 1; i < len(sortingTable); i++ {
			sortingTable[i] += sortingTable[i - 1]
		}

		for i := lcp.len; i > 0; {
			i--
			tableIndex := oldEqClasses[sortedSufArr[i]]
			sortingTable[tableIndex]--
			sufArr[sortingTable[tableIndex]] = sortedSufArr[i]
		}

		newEqClasses := lcp.allEquivClasses[newEqClIndex]

		newEqClasses[sufArr[0]] = 0
		classesCount = 1

		for i := uint(1); i < lcp.len; i++ {
			mid1, mid2 := (sufArr[i] + oldSubStrLen) % lcp.len, (sufArr[i - 1] + oldSubStrLen) % lcp.len

			if (oldEqClasses[sufArr[i]] != oldEqClasses[sufArr[i - 1]]) || (oldEqClasses[mid1] != oldEqClasses[mid2]) {
				classesCount++
			}

			newEqClasses[sufArr[i]] = classesCount - 1
		}
	}
}

// Initializes suffix array and array of equivalence classes
//
// Returns count of equivalence classes
//
// Complexity: O(N)
func initSuffixArray(str string, sufArr suffixArray, eqCl equivClasses) uint {
	sortingTable := [charsCount]uint{}

	// O(N)
	for _, ch := range str {
		sortingTable[ch]++
	}

	// O(1)
	for i := 1; i < len(sortingTable); i++ {
		sortingTable[i] += sortingTable[i - 1]
	}

	// O(N)
	for i, ch := range str {
		sortingTable[ch]--
		sufArr[sortingTable[ch]] = uint(i)
	}

	eqCl[sufArr[0]] = 0
	classesCount := uint(1)

	// O(N)
	for i := 1; i < len(sufArr); i++ {
		if str[sufArr[i]] != str[sufArr[i - 1]] {
			classesCount++
		}

		eqCl[sufArr[i]] = classesCount - 1
	}

	return classesCount
}
