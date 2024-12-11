package main

import (
	"fmt"
)

/*
You are given an m x n matrix maze (0-indexed) with empty cells (represented as '.') and walls (represented as '+').
You are also given the entrance of the maze, where entrance = [entrancerow, entrancecol] denotes the row and column of the cell you are initially standing at.

In one step, you can move one cell up, down, left, or right. You cannot step into a cell with a wall, and you cannot step outside the maze.
Your goal is to find the nearest exit from the entrance. An exit is defined as an empty cell that is at the border of the maze.
The entrance does not count as an exit.

Return the number of steps in the shortest path from the entrance to the nearest exit, or -1 if no such path exists.

Example 1:

Input: maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance = [1,2]
Output: 1
Explanation: There are 3 exits in this maze at [1,0], [0,2], and [2,3].
Initially, you are at the entrance cell [1,2].
- You can reach [1,0] by moving 2 steps left.
- You can reach [0,2] by moving 1 step up.
It is impossible to reach [2,3] from the entrance.
Thus, the nearest exit is [0,2], which is 1 step away.
Example 2:


Input: maze = [["+","+","+"],[".",".","."],["+","+","+"]], entrance = [1,0]
Output: 2
Explanation: There is 1 exit in this maze at [1,2].
[1,0] does not count as an exit since it is the entrance cell.
Initially, you are at the entrance cell [1,0].
- You can reach [1,2] by moving 2 steps right.
Thus, the nearest exit is [1,2], which is 2 steps away.
Example 3:

Input: maze = [[".","+"]], entrance = [0,0]
Output: -1
Explanation: There are no exits in this maze.

Constraints:

maze.length == m
maze[i].length == n
1 <= m, n <= 100
maze[i][j] is either '.' or '+'.
entrance.length == 2
0 <= entrancerow < m
0 <= entrancecol < n
entrance will always be an empty cell.
*/
func nearestExit(maze [][]byte, entrance []int) int {
	width:= len(maze)-1
	height:=len(maze[0])-1
	graphMaze := make([][]int,0)
	for row, valRow:= range maze {
		for col, valCol:= range valRow {
			if valCol == '.'{
				graphMaze = append(graphMaze, []int{row,col})
			}
		}
	}
	fmt.Println(graphMaze)
	queue:= make([][]int, 0)
	queue = append(queue, entrance)
	visited:=make([][]int,0)
	count:=0
	for len(queue) > 0 {
		for i:=len(queue)-1; i>=0;i-- {
			checkPoint:= queue[0]
			queue = queue[1:]
			visited = append(visited, checkPoint)
			fmt.Println("check point: ", checkPoint)
			if (checkPoint[0]==0 || checkPoint[1] ==0 || checkPoint[0] == height || checkPoint[0] == width || checkPoint[1] == height || checkPoint[1] == width) && (checkPoint[0] != entrance[0] || checkPoint[1] != entrance[1]) {
				return count
			}
			for _, val:= range graphMaze {
				if (val[0] == checkPoint[0] + 1 && val[1] == checkPoint[1]) || (val[0] == checkPoint[0] -1 && val[1] == checkPoint[1]) ||
				(val[0] == checkPoint[0] && val[1] == checkPoint[1] - 1 ) || (val[0] == checkPoint[0]  && val[1] == checkPoint[1] + 1){
					if !isVisited(visited, val) {
						queue = append(queue, val)
					}
				}
			}
		}
		count++
	}

	return -1
}

func isVisited(arr [][]int, checkValue []int) bool {
	for _, val:=range arr {
		if val[0] == checkValue[0] && val[1] == checkValue[1] {
			return true
		}
	}

	return false
}