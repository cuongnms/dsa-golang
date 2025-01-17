package main

import "fmt"

type TrieNode struct {
	node [26]*TrieNode
	isNodeEnd bool
}

type Trie struct {
    root *TrieNode
}


func Constructor() Trie {
    return Trie{root: &TrieNode{}}
}


func (this *Trie) Insert(word string)  {
	cur := this.root
	for i:=0 ; i < len(word); i++ {
		idx := word[i] - 'a'
		if cur.node[idx] == nil {
			cur.node[idx] = &TrieNode{}
		}
		cur = cur.node[idx]
	}
	cur.isNodeEnd = true
}

func (this *Trie) Print() {
	queue := make([]*TrieNode, 0)
	cur := this.root
	queue = append(queue, cur)
	count:=0 
	for len(queue) > 0 {	
		fmt.Println("step: " , count)
		for i:=len(queue)-1; i>=0 ; i-- {
			node:= queue[0]
			queue = queue[1:]
			for j:= 0 ; j < 26; j++ {
				if node.node[j] != nil {
					fmt.Println(j)
					queue = append(queue, node.node[j])
				}
			}
		}
		count++
	}
}

// insert apple
// search app
func (this *Trie) Search(word string) bool {
	cur := this.root
    for i:=0 ; i < len(word);i++ {
		if cur.node[word[i]-'a'] != nil {
			cur = cur.node[word[i]-'a']
		}else {
			return false
		}
	}
	return cur.isNodeEnd

}


func (this *Trie) StartsWith(prefix string) bool {
    cur := this.root
	
	for i:=0; i < len(prefix); i++ {
		if cur.node[prefix[i]-'a'] != nil {
			cur = cur.node[prefix[i]-'a']
		}else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
