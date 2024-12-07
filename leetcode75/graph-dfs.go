package main

import "fmt"

// import "fmt"

/*
There are n rooms labeled from 0 to n - 1 and all the rooms are locked except for room 0.
Your goal is to visit all the rooms. However, you cannot enter a locked room without having its key.
When you visit a room, you may find a set of distinct keys in it.
Each key has a number on it, denoting which room it unlocks, and you can take all of them with you to unlock the other rooms.
Given an array rooms where rooms[i] is the set of keys that you can obtain if you visited room i, return true if you can visit all the rooms, or false otherwise.

Example 1:

Input: rooms = [[1],[2],[3],[]]
Output: true
Explanation:
We visit room 0 and pick up key 1.
We then visit room 1 and pick up key 2.
We then visit room 2 and pick up key 3.
We then visit room 3.
Since we were able to visit every room, we return true.
Example 2:

Input: rooms = [[1,3],[3,0,1],[2],[0]]
Output: false
Explanation: We can not enter room number 2 since the only key that unlocks it is in that room.
Constraints:

n == rooms.length
2 <= n <= 1000
0 <= rooms[i].length <= 1000
1 <= sum(rooms[i].length) <= 3000
0 <= rooms[i][j] < n
All the values of rooms[i] are unique.
*/

func canVisitAllRooms(rooms [][]int) bool {
	lenRoom := len(rooms)
	visited := make([]bool, lenRoom)
	visited[0] = true
	var dfs func(roomIndex int)
	dfs = func(roomIndex int) {
		fmt.Println("roomIndex: ", roomIndex, " value: ", rooms[roomIndex])
		for _, indexOtherRoom := range rooms[roomIndex] {
			if visited[indexOtherRoom] {
				continue
			}
			visited[indexOtherRoom] = true
			if isVisitedAll(visited) {
				return
			} else {
				dfs(indexOtherRoom)
			}
		}
	}

	dfs(0)

	if isVisitedAll(visited) {
		return true
	} else {
		return false
	}
}

func isVisitedAll(visited []bool) bool {
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

/*
There are n cities. Some of them are connected, while some are not.
If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.
A province is a group of directly or indirectly connected cities and no other cities outside of the group.
You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.
Return the total number of provinces.

Example 1:
Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
Output: 2

Example 2:
Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
Output: 3

Constraints:
1 <= n <= 200
n == isConnected.length
n == isConnected[i].length
isConnected[i][j] is 1 or 0.
isConnected[i][i] == 1
isConnected[i][j] == isConnected[j][i]
*/

func findCircleNum(isConnected [][]int) int {
	lenInput := len(isConnected)
	isChecked := make([]bool, lenInput)
	var dfs func(currentCityIndex int)
	count := 0
	dfs = func(currentCityIndex int) {
		isChecked[currentCityIndex] = true
		for index, city := range isConnected[currentCityIndex] {
			if city == 1 && index != currentCityIndex && !isChecked[index] {
				dfs(index)
			}
		}
		
	}

	for i:=0; i < len(isConnected[0]); i++ {
		if !isChecked[i] {
			count++
			dfs(i)
		}
	}
	
	return count

}
