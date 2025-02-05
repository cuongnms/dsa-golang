package main

import (
	"fmt"
	"sort"
)

type TrieNode struct {
	node      [26]*TrieNode
	isNodeEnd bool
	count int
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if cur.node[idx] == nil {
			cur.node[idx] = &TrieNode{}
		}
		cur = cur.node[idx]
	}
	cur.count++
	cur.isNodeEnd = true
}

func (this *Trie) Print() {
	queue := make([]*TrieNode, 0)
	cur := this.root
	queue = append(queue, cur)
	count := 0
	for len(queue) > 0 {
		for i := len(queue) - 1; i >= 0; i-- {
			node := queue[0]
			queue = queue[1:]
			for j := 0; j < 26; j++ {
				if node.node[j] != nil {
					fmt.Println(j)
					if !node.node[j].isNodeEnd {
						queue = append(queue, node.node[j])
					}
				}
			}
		}
		count++
	}
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for i := 0; i < len(word); i++ {
		if cur.node[word[i]-'a'] != nil {
			cur = cur.node[word[i]-'a']
		} else {
			return false
		}
	}
	return cur.isNodeEnd

}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for i := 0; i < len(prefix); i++ {
		if cur.node[prefix[i]-'a'] != nil {
			cur = cur.node[prefix[i]-'a']
		} else {
			return false
		}
	}
	return true
}

func (this *Trie) dfsWithPrefix(wordSearch string) []string {
	cur := this.root
	for i:=0; i < len(wordSearch) ; i++ {
		if cur.node[wordSearch[i] - 'a'] != nil {
			cur = cur.node[wordSearch[i] - 'a']
		}
	}

	rsArr := make([]string, 0)
	if cur == this.root {
		return rsArr
	}
	if cur.count > 0 {
		rsArr = append(rsArr, wordSearch)
	}

	var recursive func(word string, node *TrieNode)
	recursive = func(word string, node *TrieNode) {
		if node == nil {
			return
		}
		if len(rsArr) == 3 {
			return
		}
		for i := 0; i < 26; i++ {
			if node.node[i] != nil {
				if node.node[i].isNodeEnd {
					rsArr = append(rsArr, word+string(i+'a'))
					if node.node[i].count > 0 {
						recursive(word + string(i+'a'), node.node[i])
					}
				} else {
					recursive(word + string(i+'a'), node.node[i])
				}
			}
		}
	}
	
	recursive(wordSearch, cur)
	return rsArr
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

/*
You are given an array of strings products and a string searchWord.
Design a system that suggests at most three product names from products after each character of searchWord is typed.
Suggested products should have common prefix with searchWord.
If there are more than three products with a common prefix return the three lexicographically minimums products.
Return a list of lists of the suggested products after each character of searchWord is typed.

Example 1:

Input: products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"
Output: [["mobile","moneypot","monitor"],["mobile","moneypot","monitor"],["mouse","mousepad"],["mouse","mousepad"],["mouse","mousepad"]]
Explanation: products sorted lexicographically = ["mobile","moneypot","monitor","mouse","mousepad"].
After typing m and mo all products match and we show user ["mobile","moneypot","monitor"].
After typing mou, mous and mouse the system suggests ["mouse","mousepad"].
Example 2:

Input: products = ["havana"], searchWord = "havana"
Output: [["havana"],["havana"],["havana"],["havana"],["havana"],["havana"]]
Explanation: The only word "havana" will be always suggested while typing the search word.

Constraints:

1 <= products.length <= 1000
1 <= products[i].length <= 3000
1 <= sum(products[i].length) <= 2 * 104
All the strings of products are unique.
products[i] consists of lowercase English letters.
1 <= searchWord.length <= 1000
searchWord consists of lowercase English letters.
*/
type Triee struct {
	children [26]*Triee
	words []string
}

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	root := &Triee{}
	rsArr:= make([][]string,0)
	for _, product:=range products {
		cur:= root
		for _, c:= range product {
			if cur.children[c-'a'] == nil {
				cur.children[c-'a'] = &Triee{}
			}
			cur = cur.children[c-'a']
			cur.words = append(cur.words, product)
		}
		
	}

	child := root

	for _, ch := range searchWord {
		if child != nil && child.children[ch-'a'] != nil {
			child = child.children[ch-'a']
			if len(child.words) <=3 {
				rsArr = append(rsArr, child.words)				
			}else {
				rsArr = append(rsArr, child.words[:3])
			}
			
		}else {
			child = nil
			rsArr = append(rsArr, []string{})
		}
	}

	return rsArr

}


