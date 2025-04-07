package main

import "fmt"

/*
Given the head of a linked list,
rotate the list to the right by k places.

Example 1:

Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]
Example 2:

Input: head = [0,1,2], k = 4
Output: [2,0,1]

Constraints:

The number of nodes in the list is in the range [0, 500].
-100 <= Node.val <= 100
0 <= k <= 2 * 109
  
1 2 3 4 5 
k=2

5.Next = 1
k = 2%5 = 2
k = 5-2 = 3



*/
func rotateRight(head *ListNode, k int) *ListNode {
	fmt.Println()
	if head == nil || k == 0 || head.Next == nil{
		return head
	}
	lenList:= 1
	last:= head
	for last.Next != nil {
		lenList++
		last = last.Next
	}
	times:= k % lenList
	if times == 0 {
		return head
	}
	
	last.Next = head
	for i:=0; i < lenList - times; i++ {
		last =last.Next
	}
	head = last.Next
	last.Next = nil 
	return head
}

/*
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. 
Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.

Example 1:

Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

1 -> 2

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4
 
Constraints:

1 <= capacity <= 3000
0 <= key <= 104
0 <= value <= 105
At most 2 * 105 calls will be made to get and put.
*/
type DListNode struct {
	Key int
	Val int
	Next *DListNode
	Prev *DListNode
}

type LRUCache struct {
	MapNode map[int]*DListNode
	Head *DListNode
	Tail *DListNode
	cap int
}


func Constructor(capacity int) LRUCache {
	mapNode := make(map[int]*DListNode)
	head := &DListNode{-1, -1, nil, nil}
	tail := &DListNode{-1, -1, nil, nil}
	head.Next = tail
	tail.Prev = head
	cap:= capacity
	return LRUCache{MapNode: mapNode, Head: head, Tail: tail, cap: cap}
}

func (this *LRUCache) DeleteNode(node *DListNode) {
	tmpPrev:= node.Prev
	tmpNext:= node.Next

	tmpPrev.Next = tmpNext
	tmpNext.Prev = tmpPrev

}

func (this *LRUCache) AddNode(node *DListNode) {
	tmp := this.Head.Next
	tmp.Prev = node
	this.Head.Next = node
	node.Next = tmp
	node.Prev = this.Head
}

func (this *LRUCache) Get(key int) int {
    node, exist := this.MapNode[key]
	if !exist {
		return -1
	}else {
		delete(this.MapNode, key)
		this.DeleteNode(node)
		this.AddNode(node)
		this.MapNode[key] = this.Head.Next
		return node.Val
	}
}

func (this *LRUCache) Put(key int, value int)  {
    node, exist:= this.MapNode[key]
	if exist {
		delete(this.MapNode, key)
		this.DeleteNode(node)
	}
	if len(this.MapNode) == this.cap {
		delete(this.MapNode, this.Tail.Prev.Key)
		this.DeleteNode(this.Tail.Prev)
	}

	addNode :=&DListNode{Key: key, Val: value, Next: nil, Prev: nil}
	this.AddNode(addNode)
	this.MapNode[key] = this.Head.Next
}