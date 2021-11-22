package main

import "fmt"

type trieNode struct {
	children map[rune]*trieNode
}

func newTrieNode() *trieNode {
	var t trieNode
	t.children = make(map[rune]*trieNode)
	return &t
}

func (tn *trieNode) search(word string) *trieNode {
	currNode := tn
	for _, char := range word {
		if _, ok := currNode.children[char]; ok {
			currNode = currNode.children[char]
		} else {
			return nil
		}
	}
	return currNode
}

func (tn *trieNode) insert(word string) {
	currNode := tn
	for _, char := range word {
		if _, ok := currNode.children[char]; ok {
			currNode = currNode.children[char]
		} else {
			newNode := newTrieNode()
			currNode.children[char] = newNode
			currNode = newNode
		}
	}
	currNode.children['*'] = nil
}

func (tn *trieNode) collectAllWords(word string, words *[]string) []string {
	for key, childNode := range tn.children {
		if key == '*' {
			*words = append(*words, word)
		} else {
			childNode.collectAllWords(word+string(key), words)
		}
	}
	return *words
}

func (tn *trieNode) autocomplete(prefix string) []string {
	currNode := tn.search(prefix)
	if currNode == nil {
		return nil
	}
	var words []string
	var word string
	return currNode.collectAllWords(word, &words)
}

func main() {
	root := newTrieNode()
	//fmt.Println(root)
	root.insert("dog")
	root.insert("apple")
	root.insert("doggy")
	root.insert("cat")
	root.insert("cactus")

	var words []string
	var word string
	fmt.Println(root.collectAllWords(word, &words))
	fmt.Println(root.autocomplete("do"))
}
