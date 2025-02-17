package main

import "fmt"

func main() {
	//============150==============================
	// root:= sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	// fmt.Println(root)
	// fmt.Println(maxSubArray([]int{-1, -2 , 3, 4}))
	// fmt.Println(combine(4,2))
	// values := []interface{}{
	// 	3,1,4,nil,2,
	// }
	// root := CreateTree(values)
	// fmt.Println(kthSmallest(root, 1))
	values:= []interface{}{1,2}
	list:= CreateLL(values)
	fmt.Println(rotateRight(list, 1))
}