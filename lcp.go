// Package LCP provides data structure which allows effectively calculate
// longest common prefix of two any suffixes of specified string.
package lcp

import (
	. "lcp/internal"
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
	result := lcpImpl{}

	result.len = uint64(len(str))

	// O(N * log N)
	sufArr, lcpArr := makeSuffixAndLcpArrays(str)

	// O(N)
	result.sortedSuffixesPos = sufArr.makeInversed()

	// O(N)
	result.commonPrefixesLengths = MakeMinSegmentTree(lcpArr)

	return &result
}

// Implementation

const (
	charMin    = 0
	charMax    = math.MaxUint8
	charsCount = charMax - charMin + 1
)

type suffixArray []uint64
type equivClasses []uint64
type inverseSuffixArray []uint64
type lcpArray []uint64

type lcpImpl struct {
	len uint64
	sortedSuffixesPos     inverseSuffixArray
	commonPrefixesLengths SegmentTree
}

func (lcp *lcpImpl) Get(firstSuffixIndex uint64, secondSuffixIndex uint64) uint64 {
	if (firstSuffixIndex >= lcp.len) || (secondSuffixIndex >= lcp.len) {
		panic("Index out of bounds")
	}

	if firstSuffixIndex == secondSuffixIndex {
		return uint64(len(lcp.sortedSuffixesPos)) - firstSuffixIndex
	}

	minIndex := lcp.sortedSuffixesPos[firstSuffixIndex]
	maxIndex := lcp.sortedSuffixesPos[secondSuffixIndex]

	if minIndex > maxIndex {
		minIndex, maxIndex = maxIndex, minIndex
	}

	return lcp.commonPrefixesLengths.Get(minIndex, maxIndex - 1)
}


func makeSuffixAndLcpArrays(str string) (suffixArray, lcpArray) {
	strLen := uint64(len(str))

	sufArr := make(suffixArray, strLen)
	eqCl := make(equivClasses, strLen)

	if strLen == 0 {
		return sufArr, nil
	}

	classesCount := initSuffixArray(str, &sufArr, &eqCl)

	sortedSufArr := make(suffixArray, strLen)
	eqClTemp := make(equivClasses, strLen)

	sortingTable := make([]uint64, strLen)

	lcp := make(lcpArray, strLen - 1)
	lcpTemp := make(lcpArray, strLen - 1)

	rPos := make([]uint64, strLen)
	lPos := make([]uint64, strLen)

	// O(N * logN)
	for oldSubStrLen := uint64(1); oldSubStrLen < strLen; oldSubStrLen *= 2 {
		initPoses(rPos, lPos, sufArr, eqCl)

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

		lcpRmq := MakeMinSegmentTree(lcp)
		for i := range lcpTemp {
			subStr1Pos, subStr2Pos := sufArr[i], sufArr[i + 1]

			if eqCl[subStr1Pos] == eqCl[subStr2Pos] {
				str1Pos := (subStr1Pos + oldSubStrLen) % strLen
				str2Pos := (subStr2Pos + oldSubStrLen) % strLen

				lcpTemp[i] = Min(strLen, oldSubStrLen + lcpRmq.Get(lPos[eqCl[str1Pos]], rPos[eqCl[str2Pos]] - 1))
			} else {
				lcpTemp[i] = lcp[rPos[eqCl[subStr1Pos]]]
			}
		}

		lcp, lcpTemp = lcpTemp, lcp
		eqCl, eqClTemp = eqClTemp, eqCl
	}

	// avoiding circular substrings
	// O(N)
	for i := range lcp {
		lcp[i] = Min(lcp[i], Min(strLen - sufArr[i], strLen - sufArr[i + 1]))
	}

	return sufArr, lcp
}

// Initializes suffix array and array of equivalence classes
// Returns count of equivalence classes
// Complexity: O(N)
func initSuffixArray(str string, sufArr *suffixArray, eqCl *equivClasses) uint64 {
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
	classesCount := uint64(1)

	// O(N)
	for i := 1; i < len(*sufArr); i++ {
		if str[(*sufArr)[i]] != str[(*sufArr)[i-1]] {
			classesCount++
		}

		(*eqCl)[(*sufArr)[i]] = classesCount - 1
	}

	return classesCount
}

func initPoses(rPos []uint64, lPos []uint64, sufArr suffixArray, eqCl equivClasses) {
	for i := range rPos {
		rPos[eqCl[sufArr[i]]] = uint64(i)
	}

	for i := uint64(len(lPos)); i > 0; {
		i--
		lPos[eqCl[sufArr[i]]] = i
	}
}

// O(N)
func (sufArr suffixArray) makeInversed() inverseSuffixArray {
	result := make(inverseSuffixArray, len(sufArr))

	for i, val := range sufArr {
		result[val] = uint64(i)
	}

	return result
}
