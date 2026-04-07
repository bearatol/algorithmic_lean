package main

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := ConstructorTrie()

	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("Search(apple) = false, want true")
	}
	if trie.Search("app") {
		t.Errorf("Search(app) = true, want false")
	}
	if !trie.StartsWith("app") {
		t.Errorf("StartsWith(app) = false, want true")
	}
}
