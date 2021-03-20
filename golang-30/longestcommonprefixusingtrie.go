package main

import (
	"fmt"
)

const (
	ALPHABET_SIZE = 26
)

var (
	index int = 0
)

type trieNode struct {
	childrens [ALPHABET_SIZE]*trieNode
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

func (t *trie) insert(word string) {
	wordLength := len(word)
	if t.root == nil {
		t.root = &trieNode{}
	}
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = &trieNode{}
		}
		current = current.childrens[index]
	}
	current.isWordEnd = true
}

func (t *trie) countChildren() int {
	count := 0
	for i := 0; i < ALPHABET_SIZE; i++ {
		if t.root.childrens[i] != nil {
			count += 1
			index = i
		}
	}
	return count
}

func (t *trie) findLongestPrefix() string {
	prefix := ""
	for t.countChildren() == 1 && t.root.isWordEnd == false {
		t.root = t.root.childrens[index]
		char := 97 + index
		prefix += string(char)
	}
	return prefix
}

func (t *trie) find(word string) bool {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}
	if current.isWordEnd {
		return true
	}
	return false
}

func main() {
	t := trie{}
	words := []string{"hello", "hey", "hellical", "hell", "helix", "he"}
	for i := 0; i < len(words); i++ {
		t.insert(words[i])
	}
	wordsToFind := []string{"hello", "hey", "hellical", "hell", "helix", "he", "hola", "iglo"}
	for i := 0; i < len(wordsToFind); i++ {
		found := t.find(wordsToFind[i])
		if found {
			fmt.Printf("found the word [\"%s\"]\n", wordsToFind[i])
		} else {
			fmt.Printf("not found the word [\"%s\"]\n", wordsToFind[i])
		}
	}
	fmt.Printf("\nthe longest prefix is [\"%s\"]\n", t.findLongestPrefix())
}
