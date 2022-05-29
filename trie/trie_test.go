package trie

import (
	"fmt"
	"testing"
)

func TestPrefixTrie(t *testing.T) {
	word := "abc"
	obj := Constructor()
	obj.Insert(word)
	param_2 := obj.Search("ab")
	param_3 := obj.StartsWith("abcd")
	fmt.Println("param_2", param_2)
	fmt.Println("param_3", param_3)
}
