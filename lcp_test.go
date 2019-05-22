package lcp_test

import (
	lcppkg "lcp"
	"testing"
)

type testCase struct {
	first, second uint64
	expected      uint64
}

func (tc testCase) assertValue(lcp lcppkg.LongestCommonPrefix, t *testing.T) {
	actual := lcp.Get(tc.first, tc.second)
	if actual != tc.expected {
		t.Errorf("Get(%d, %d) returned %d ; expected: %d", tc.first, tc.second, actual, tc.expected)
	}
}

func (tc testCase) assertPanic(lcp lcppkg.LongestCommonPrefix, t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Get(%d, %d) didn't panic, but it should", tc.first, tc.second)
		}
	}()

	lcp.Get(tc.first, tc.second)
}

func TestFromSRS(t *testing.T) {
	const testedString = "abacaba"

	cases := [...]testCase{
		{first: 0, second: 1, expected: 0}, // i.e. ""
		{first: 0, second: 2, expected: 1}, // i.e. "a"
		{first: 0, second: 3, expected: 0}, // i.e. ""
		{first: 0, second: 4, expected: 3}, // i.e. "aba"
		{first: 0, second: 5, expected: 0}, // i.e. ""
		{first: 0, second: 6, expected: 1}, // i.e. "a"
		{first: 0, second: 0, expected: 7}, // i.e. "abacaba"
		{first: 1, second: 5, expected: 2}, // i.e. "ab"
		{first: 2, second: 6, expected: 1}, // i.e. "a"
	}

	lcp := lcppkg.NewLongestCommonPrefix(testedString)

	for _, c := range cases {
		c.assertValue(lcp, t)
	}

	// second is out of bounds
	testCase{first: 0, second: 7, expected: 100}.assertPanic(lcp, t)

	// first is out of bounds
	testCase{first: 8, second: 0, expected: 100}.assertPanic(lcp, t)

	// both is out of bounds
	testCase{first: 7, second: 7, expected: 0}.assertPanic(lcp, t)
}

func TestEmpty(t *testing.T) {
	cases := [...]testCase{
		{first: 0, second: 0, expected: 100},
		{first: 1, second: 1, expected: 100},
		{first: 10, second: 10, expected: 100},
		{first: 0, second: 5, expected: 100},
		{first: 3, second: 0, expected: 100},
	}

	lcp := lcppkg.NewLongestCommonPrefix("")

	for _, c := range cases {
		c.assertPanic(lcp, t)
	}
}

func TestSamePrefixes(t *testing.T) {
	const testedString = "banana"
	const lenU64 = uint64(len(testedString))

	lcp := lcppkg.NewLongestCommonPrefix(testedString)

	for i := uint64(0); i < lenU64; i++ {
		testCase{first: i, second: i, expected: lenU64 - i}.assertValue(lcp, t)
	}
}

func TestFearOfBigWords(t *testing.T) {
	const testedString = "Hippopotomonstrosesquippedaliophobia" // fear of big words
	const lenU64 = uint64(len(testedString))

	pCases := [...]testCase{
		{first: 0, second: 0, expected: lenU64},
		{first: 3, second: 5, expected: 2},
		{first: 22, second: 2, expected: 2},
		{first: 12, second: 16, expected: 1},
		{first: 4, second: 29, expected: 2},
		{first: 7, second: 12, expected: 0},
		{first: 28, second: 15, expected: 0},
		{first: 35, second: 26, expected: 1},
	}

	nCases := [...]testCase{
		{first: 36, second: 26, expected: 100},
		{first: 26, second: 414, expected: 100},
	}

	lcp := lcppkg.NewLongestCommonPrefix(testedString)

	for _, c := range pCases {
		c.assertValue(lcp, t)
	}

	for _, c := range nCases {
		c.assertPanic(lcp, t)
	}
}
