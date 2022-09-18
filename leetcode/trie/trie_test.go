package trie

import (
	"fmt"
	"testing"
)

func TestTrie_Search(t *testing.T) {
	tr := NewTrie()
	tr.Insert("test")
	fmt.Println(tr.Search("test"))
	fmt.Println(tr.StartWith("te"))
}
