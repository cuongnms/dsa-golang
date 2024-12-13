package main


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
