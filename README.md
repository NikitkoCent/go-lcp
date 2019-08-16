# go-lcp
Effectively search the longest common prefix length for two specified suffixes of the given string.

Takes `O(N * log(N))` for preprocessing and `O(log(N))` on searching (`N` is length of the string).

# Usage example ([play.golang.org](https://play.golang.org/p/LdcJpaPtGAS))
```go
package main

import (
	"fmt"
	lcppkg "github.com/NikitkoCent/go-lcp"
)

func main() {
	const word = "abacaba"

	lcp := lcppkg.NewLongestCommonPrefix(word) // O(N * log N)

	// Suffix at 0 offset (first parameter) is 'abacaba'
	// Suffix at 4 offset (second parameter) is 'aba'
	// Their longest common prefix length is 3
	// So result = 3 will be printed
	fmt.Println(lcp.Get(0, 4))                 // O(log N)

	// 6 will be printed
	fmt.Println(lcp.Get(1, 1))

	// The next string will cause panic because one of parameters is out of bounds
	// fmt.Println(lcp.Get(0, 7))
}
```
