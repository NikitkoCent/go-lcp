// Package LCP provides data structure which allows effectively calculate
// longest common prefix of two any suffixes of specified string.
package LCP

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

type lcpImpl struct {

}

func (lcp *lcpImpl) Get(firstPrefixIndex uint64, secondPrefixIndex uint64) uint64 {
	return 0
}
