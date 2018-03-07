package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	root := NewNode()
	root.Insert("apple")
	if root.Find("apple") != "apple" {
		t.Error("error Inserting and finding from trie")
	}
	root.Insert("app")
	if root.Find("app") != "app" {
		t.Error("error finding prefix match")
	}
	root.Insert("appocalypse")
	if root.Find("appocalypse") != "appocalypse" {
		t.Error("error finding prefix match")
	}
	r := root.FindPrefixMatches("app")
	fmt.Println(r)
}
