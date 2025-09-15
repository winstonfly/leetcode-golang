package tree

import "strings"

// NO.208 前缀树
type Trie struct {
	m map[string]string
}

func Constructor1() Trie {
	return Trie{
		m: make(map[string]string),
	}
}

func (this *Trie) Insert(word string) {
	this.m[word] = word
}

func (this *Trie) Search(word string) bool {
	if this.m[word] == "" {
		return false
	}

	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range this.m {
		if strings.HasPrefix(v, prefix) {
			return true
		}
	}

	return false
}
