package main
func maxArea(height []int) int {
	left:= 0
	right:= len(height) - 1
	area := 0
	for left < right {
		if area <= findMin(height[left], height[right])*(right-left) {
			area = findMin(height[left], height[right]) * (right - left)
		}
		if height[left] < height[right] {
			left++
		} else {
			right++
		}
	}
	return area
}

func findMin(a, b int)	int {
	if a > b {
        return b
    }
    return a
}