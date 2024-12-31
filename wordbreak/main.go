package main

import "fmt"

type TriNode struct {
	children map[rune]*TriNode
	isEnd    bool
}

type Trie struct {
	root *TriNode
}

func newTrie() *Trie {
	return &Trie{
		root: &TriNode{
			children: map[rune]*TriNode{},
			isEnd:    false,
		},
	}
}

func (t *Trie) insert(word string) {
	current := t.root
	for _, c := range word {
		if _, exists := current.children[c]; !exists {
			current.children[c] = &TriNode{
				children: map[rune]*TriNode{},
				isEnd:    false,
			}
		}
		current = current.children[c]
	}
	current.isEnd = true
}
func (t *Trie) Search(word string) bool {
	current := t.root
	for _, c := range word {
		if _, exists := current.children[c]; !exists {
			return false
		}
		current = current.children[c]
	}
	return current.isEnd
}

// a => a => a => is end => a => is End = true
func main() {

	fmt.Println(wordBreak("aaaaaaaaaas", []string{"aaaa", "aaa"}))
}

func wordBreak(s string, wordDic []string) bool {

	trie := newTrie()
	for _, word := range wordDic {
		trie.insert(word)
	}
	memo := map[int]bool{}
	return dfs(0, s, trie, memo)
}

func dfs(start int, s string, trie *Trie, memo map[int]bool) bool {
	if start == len(s) {
		return true
	}

	if val, ok := memo[start]; ok {
		return val
	}

	currentNode := trie.root

	for i := start; i < len(s); i++ {
		ch := rune(s[i])
		nextNode, exists := currentNode.children[ch]
		if !exists {
			return false
		}
		currentNode = nextNode
		if currentNode.isEnd { //aaa
			if dfs(i+1, s, trie, memo) {
				memo[start] = true
				return true
			}
		}
	}
	memo[start] = false
	return false
}
