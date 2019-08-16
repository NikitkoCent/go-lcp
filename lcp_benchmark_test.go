package lcp_test

import (
	lcppkg "github.com/NikitkoCent/go-lcp"
	"math/rand"
	"os"
	"testing"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const maxStringLength = 1000000

var (
	maxString     string
	firstIndexes  []uint
	secondIndexes []uint
)

func generateString(length uint) string {
	result := make([]byte, length)

	for i := range result {
		result[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(result)
}

func doNewLongestCommonPrefixBench(strLength uint, b *testing.B) {
	str := maxString[:strLength]

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lcppkg.NewLongestCommonPrefix(str)
	}
}

func doGetBench(strLength uint, b *testing.B) {
	str := maxString[:strLength]
	lcp := lcppkg.NewLongestCommonPrefix(str)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		firstIndex := firstIndexes[i % maxStringLength]
		secondIndex := secondIndexes[i % maxStringLength]
		lcp.Get(firstIndex % strLength, secondIndex % strLength)
	}
}


func BenchmarkNewLongestCommonPrefix10(b *testing.B) {
	doNewLongestCommonPrefixBench(10, b)
}

func BenchmarkNewLongestCommonPrefix100(b *testing.B) {
	doNewLongestCommonPrefixBench(100, b)
}

func BenchmarkNewLongestCommonPrefix1000(b *testing.B) {
	doNewLongestCommonPrefixBench(1000, b)
}

func BenchmarkNewLongestCommonPrefix10000(b *testing.B) {
	doNewLongestCommonPrefixBench(10000, b)
}

func BenchmarkNewLongestCommonPrefix100000(b *testing.B) {
	doNewLongestCommonPrefixBench(100000, b)
}

func BenchmarkNewLongestCommonPrefix1000000(b *testing.B) {
	doNewLongestCommonPrefixBench(1000000, b)
}


func BenchmarkGet10(b *testing.B) {
	doGetBench(10, b)
}

func BenchmarkGet100(b *testing.B) {
	doGetBench(100, b)
}

func BenchmarkGet1000(b *testing.B) {
	doGetBench(1000, b)
}

func BenchmarkGet10000(b *testing.B) {
	doGetBench(10000, b)
}

func BenchmarkGet100000(b *testing.B) {
	doGetBench(100000, b)
}

func BenchmarkGet1000000(b *testing.B) {
	doGetBench(1000000, b)
}


func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())

	maxString = generateString(1000000)
	firstIndexes = make([]uint, maxStringLength)
	secondIndexes = make([]uint, maxStringLength)

	for i := range firstIndexes {
		firstIndexes[i] = uint(i)
		secondIndexes[i] = uint(i)
	}

	rand.Shuffle(maxStringLength, func(i, j int) {
		firstIndexes[i], firstIndexes[j] = firstIndexes[j], firstIndexes[i]
	})
	rand.Shuffle(maxStringLength, func(i, j int) {
		secondIndexes[i], secondIndexes[j] = secondIndexes[j], secondIndexes[i]
	})

	os.Exit(m.Run())
}
