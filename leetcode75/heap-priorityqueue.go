package main

import "slices"

/*
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?



Example 1:

Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
Example 2:

Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4


Constraints:

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/

type Heap struct {
	arr  []int
}

func NewHeap() *Heap {
	arr := make([]int, 0)
	return &Heap{arr: arr}
}

func (h *Heap) heapifyUp(index int) {
	for index > 0 {
		parent:= (index-1)/2
		if h.arr[index]>=h.arr[parent]{
			break
		}
		h.arr[index], h.arr[parent] = h.arr[parent], h.arr[index]
		index = parent
	}
}

func (h *Heap) heapifyDown(index int) {
	lastIndex:=len(h.arr) - 1

	for {
		leftChild:=2*index+1
		rightChild:=2*index+2
		smallest:=index
		if leftChild <= lastIndex && h.arr[leftChild] < h.arr[smallest] {
			smallest = leftChild
		}

		if rightChild <= lastIndex && h.arr[rightChild] < h.arr[smallest] {
			smallest = rightChild
		}
		if smallest == index {
			break
		}
		h.arr[index], h.arr[smallest] = h.arr[smallest] , h.arr[index]
		index = smallest
	}

}

func (h *Heap) insert(val int) {
	h.arr = append(h.arr, val)
	h.heapifyUp(len(h.arr)-1)
}

func (h *Heap) pop() int {
	if len(h.arr) ==0 {
		return -1
	}

	min:=h.arr[0]
	h.arr[0] = h.arr[len(h.arr) - 1]
	h.arr= h.arr[:len(h.arr)-1]
	h.heapifyDown(0)
	return min
}

func findKthLargest(nums []int, k int) int {
	b:=NewHeap()
	for i:=0; i <k;i++ {
		b.insert(nums[i])
	}
	for i:=k; i < len(nums);i++{
		if nums[i] > b.arr[0] {
			b.pop()
			b.insert(nums[i])
		}
	}
	return b.arr[0]
}

/*
You have a set which contains all positive integers [1, 2, 3, 4, 5, ...].

Implement the SmallestInfiniteSet class:

SmallestInfiniteSet() Initializes the SmallestInfiniteSet object to contain all positive integers.
int popSmallest() Removes and returns the smallest integer contained in the infinite set.
void addBack(int num) Adds a positive integer num back into the infinite set, if it is not already in the infinite set.
 

Example 1:

Input
["SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"]
[[], [2], [], [], [], [1], [], [], []]
Output
[null, null, 1, 2, 3, null, 1, 4, 5]

Explanation
SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
smallestInfiniteSet.addBack(2);    // 2 is already in the set, so no change is made.
smallestInfiniteSet.popSmallest(); // return 1, since 1 is the smallest number, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 2, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 3, and remove it from the set.
smallestInfiniteSet.addBack(1);    // 1 is added back to the set.
smallestInfiniteSet.popSmallest(); // return 1, since 1 was added back to the set and
                                   // is the smallest number, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 4, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 5, and remove it from the set.
 
Constraints:

1 <= num <= 1000
At most 1000 calls will be made in total to popSmallest and addBack.
*/

type SmallestInfiniteSet struct {
    arr []int
}


func SmallestInfiniteSetConstructor() SmallestInfiniteSet {
	arr:= make([]int, 0)
	for i :=0; i<1000; i++{
		arr = append(arr, i+1)
	}
    return SmallestInfiniteSet{
		arr: arr,
	}
}

func (this *SmallestInfiniteSet) heapifyDownSet(index int) {
	lastIndex:= len(this.arr) - 1
	for {
		leftIndex:=2*index + 1
		rightIndex:=2*index + 2
		smallestValIndex:=index
		if leftIndex <= lastIndex && this.arr[leftIndex] < this.arr[smallestValIndex] {
            smallestValIndex = leftIndex
        }
		if rightIndex <= lastIndex && this.arr[rightIndex] < this.arr[smallestValIndex] {
            smallestValIndex = rightIndex
        }
		if smallestValIndex == index {
            break
        }
		this.arr[index], this.arr[smallestValIndex] = this.arr[smallestValIndex], this.arr[index]
		index= smallestValIndex

	}
}

func (this *SmallestInfiniteSet) heapifyUpSet(index int) {
	for index >0 {
		parentIndex:= (index-1)/2
		if this.arr[index] >= this.arr[parentIndex] {
			break
		}
		this.arr[index], this.arr[parentIndex] = this.arr[parentIndex], this.arr[index]
		index = parentIndex
	}
}


func (this *SmallestInfiniteSet) PopSmallest() int {
	rs:=this.arr[0]
	this.arr[0] = this.arr[len(this.arr)-1]
	this.arr= this.arr[:len(this.arr)-1]
	this.heapifyDownSet(0)
	return rs
}


func (this *SmallestInfiniteSet) AddBack(num int)  {
    if slices.Contains(this.arr, num) {
		return
	}else {
		this.arr = append(this.arr, num)
        this.heapifyUpSet(len(this.arr)-1)
	}
}


/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */