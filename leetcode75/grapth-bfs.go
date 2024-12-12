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
	fmt.Println()
	width := len(maze) - 1
	height := len(maze[0]) - 1
	queue := make([][]int, 0)
	queue = append(queue, entrance)
	visited := make([][]bool, width+1)
	for i := range visited {
		visited[i] = make([]bool, height+1)
	}
	visited[entrance[0]][entrance[1]] = true

	count := 0
	for len(queue) > 0 {
		for i := len(queue) - 1; i >= 0; i-- {
			checkPoint := queue[0]
			queue = queue[1:]
			if (checkPoint[0] == 0 || checkPoint[1] == 0 || checkPoint[0] == width || checkPoint[1] == height) && (checkPoint[0] != entrance[0] || checkPoint[1] != entrance[1]) {
				return count
			}

			if checkPoint[0]+1 <= width && maze[checkPoint[0]+1][checkPoint[1]] == '.' && !visited[checkPoint[0]+1][checkPoint[1]] {
				visited[checkPoint[0]+1][checkPoint[1]] = true

				queue = append(queue, []int{checkPoint[0] + 1, checkPoint[1]})
			}

			if checkPoint[0]-1 >= 0 && maze[checkPoint[0]-1][checkPoint[1]] == '.' && !visited[checkPoint[0]-1][checkPoint[1]] {
				visited[checkPoint[0]-1][checkPoint[1]] = true

				queue = append(queue, []int{checkPoint[0] - 1, checkPoint[1]})
			}

			if checkPoint[1]+1 <= height && maze[checkPoint[0]][checkPoint[1]+1] == '.' && !visited[checkPoint[0]][checkPoint[1]+1] {
				visited[checkPoint[0]][checkPoint[1]+1] = true

				queue = append(queue, []int{checkPoint[0], checkPoint[1] + 1})
			}

			if checkPoint[1]-1 >= 0 && maze[checkPoint[0]][checkPoint[1]-1] == '.' && !visited[checkPoint[0]][checkPoint[1]-1] {
				visited[checkPoint[0]][checkPoint[1]-1] = true

				queue = append(queue, []int{checkPoint[0], checkPoint[1] - 1})
			}
		}
		count++
	}

	return -1
}

/*
You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

Example 1:

Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
Output: 4
Example 2:

Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
Example 3:

Input: grid = [[0,2]]
Output: 0
Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 10
grid[i][j] is 0, 1, or 2.
*/
func orangesRotting(grid [][]int) int {
	queue := make([][]int, 0)
	minutes := 0
	oranges := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				oranges++
			}
		}
	}
	if oranges == 0 {
		return 0
	}

	w := len(grid[0])
	h := len(grid)
	for len(queue) > 0 {
		for i := len(queue) - 1; i >= 0; i-- {
			check := queue[0]
			queue = queue[1:]
			if check[0]+1 <= h-1 && grid[check[0]+1][check[1]] == 1 {
				oranges--
				grid[check[0]+1][check[1]]++
				queue = append(queue, []int{check[0] + 1, check[1]})
			}
			if check[0]-1 >= 0 && grid[check[0]-1][check[1]] == 1 {
				oranges--
				grid[check[0]-1][check[1]]++
				queue = append(queue, []int{check[0] - 1, check[1]})
			}

			if check[1]+1 <= w-1 && grid[check[0]][check[1]+1] == 1 {
				oranges--
				grid[check[0]][check[1]+1]++
				queue = append(queue, []int{check[0], check[1] + 1})
			}

			if check[1]-1 >= 0 && grid[check[0]][check[1]-1] == 1 {
				oranges--
				grid[check[0]][check[1]-1]++
				queue = append(queue, []int{check[0], check[1] - 1})
			}
		}
		minutes++

	}

	if oranges > 0 {
		return -1
	} else {
		return minutes - 1
	}

}
