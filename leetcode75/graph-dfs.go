package main

import (
	"fmt"
	"slices"
	// "slices"
	// "slices"
)

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

	for i := 0; i < len(isConnected[0]); i++ {
		if !isChecked[i] {
			count++
			dfs(i)
		}
	}

	return count

}

/*
There are n cities numbered from 0 to n - 1 and n - 1 roads such that there is only one way to travel between two different cities (this network form a tree).
Last year, The ministry of transport decided to orient the roads in one direction because they are too narrow.
Roads are represented by connections where connections[i] = [ai, bi] represents a road from city ai to city bi.
This year, there will be a big event in the capital (city 0), and many people want to travel to this city.
Your task consists of reorienting some roads such that each city can visit the city 0. Return the minimum number of edges changed.
It's guaranteed that each city can reach city 0 after reorder.

Example 1:
Input: n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]
Output: 3
Explanation: Change the direction of edges show in red such that each node can reach the node 0 (capital).

Example 2:
Input: n = 5, connections = [[1,0],[1,2],[3,2],[3,4]]
Output: 2
Explanation: Change the direction of edges show in red such that each node can reach the node 0 (capital).
Example 3:

Input: n = 3, connections = [[1,0],[2,0]]
Output: 0

Constraints:

2 <= n <= 5 * 104
connections.length == n - 1
connections[i].length == 2
0 <= ai, bi <= n - 1
ai != bi
*/

func minReorder(n int, connections [][]int) int {
	adj := make([][]int, 0)
	for i := 0; i < n; i++ {
		adj = append(adj, []int{})
	}
	for _, city := range connections {
		adj[city[0]] = append(adj[city[0]], city[1])
		adj[city[1]] = append(adj[city[1]], -city[0])
	}

	isVisited := make([]bool, n)

	var dfs func(from int) int
	dfs = func(from int) int {
		count := 0
		isVisited[from] = true
		for _, val := range adj[from] {
			if !isVisited[abs(val)] {
				if val > 0 {
					count += dfs(abs(val)) + 1
				} else {
					count += dfs(abs(val))
				}
			}
		}
		return count
	}

	return dfs(0)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/*
You are given an array of variable pairs equations and an array of real numbers values, where equations[i] = [Ai, Bi] and values[i] represent the equation Ai / Bi = values[i].
Each Ai or Bi is a string that represents a single variable.

You are also given some queries, where queries[j] = [Cj, Dj] represents the jth query where you must find the answer for Cj / Dj = ?.

Return the answers to all queries. If a single answer cannot be determined, return -1.0.

Note: The input is always valid. You may assume that evaluating the queries will not result in division by zero and that there is no contradiction.

Note: The variables that do not occur in the list of equations are undefined, so the answer cannot be determined for them.

Example 1:

Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
Explanation:
Given: a / b = 2.0, b / c = 3.0
queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
return: [6.0, 0.5, -1.0, 1.0, -1.0 ]
note: x is undefined => -1.0

Example 2:
Input: equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
Output: [3.75000,0.40000,5.00000,0.20000]

Example 3:
Input: equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
Output: [0.50000,2.00000,-1.00000,-1.00000]


Constraints:
1 <= equations.length <= 20
equations[i].length == 2
1 <= Ai.length, Bi.length <= 5
values.length == equations.length
0.0 < values[i] <= 20.0
1 <= queries.length <= 20
queries[i].length == 2
1 <= Cj.length, Dj.length <= 5
Ai, Bi, Cj, Dj consist of lower case English letters and digits.

"a":{"b": 2.0},
"b":{"a": 0.5, "c": 3}
"c":{"b": 0.333}

*/

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	for index, val := range equations {
		_, existOut := graph[val[0]]
		if existOut {
			graph[val[0]][val[1]] = values[index]
		} else {
			tmp:= make(map[string]float64)
			tmp[val[1]] = values[index]
			graph[val[0]] = tmp
		}
		_, existOut = graph[val[1]]
		if existOut {
			graph[val[1]][val[0]] = 1/(values[index])
		} else {
			tmp := make(map[string]float64)
			tmp[val[0]] = 1/values[index]
			graph[val[1]] = tmp
		}
	}
	var dfs func(start, end string, visit []string) float64
	dfs = func(start , end string, visit []string) float64 {
		visit = append(visit, start)
		if val, ok:= graph[start][end]; ok {
			return val
		}
		for key, val := range graph[start] {
			if !slices.Contains(visit, key) {
				if dfsRs := dfs(key,end,visit); dfsRs != -1.0 {
					return dfsRs*val
				} 
			}
			
		}
		return -1.0
	}

	flRs:=make([]float64, 0)
	for _, q:= range queries {
		_, okStart:= graph[q[0]]
		_, okEnd:= graph[q[1]]
		visit:= make([]string, 0)

		if !okEnd || !okStart {
			flRs = append(flRs, -1.0)
		}else {
			flRs = append(flRs, dfs(q[0],q[1], visit))
		}
	}

	return flRs
}
