package main

import "fmt"

type TrieNode struct {
	node [26]*TrieNode
}

type Trie struct {
    root *TrieNode
}


func Constructor() Trie {
    return Trie{root: &TrieNode{}}
}


func (this *Trie) Insert(word string)  {
	cur := this.root
	for ch:=range word {
		idx:= ch - 'a'
		if cur.node[idx] == nil {
			cur.node[idx] = &TrieNode{}
		}
		cur = cur.node[idx]
	}
}

func (this *Trie) Print() {
	fmt.Println(this)
}

func (this *Trie) Search(word string) bool {
    return false
}


func (this *Trie) StartsWith(prefix string) bool {
    return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
