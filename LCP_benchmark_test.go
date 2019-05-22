package LCP_test

import (
	"LCP"
	"math/rand"
	"os"
	"testing"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateString(length uint64) string {
	result := make([]byte, length)

	for i := range result {
		result[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(result)
}

func doNewLongestCommonPrefixBench(strLength uint64, b *testing.B) {
	str := generateString(strLength)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		LCP.NewLongestCommonPrefix(str)
	}
}

func doGetBench(strLength uint64, b *testing.B) {
	str := generateString(strLength)
	lcp := LCP.NewLongestCommonPrefix(str)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lcp.Get(strLength / 3, (strLength / 3) * 2)
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

	os.Exit(m.Run())
}
