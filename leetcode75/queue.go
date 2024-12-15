package main

import "fmt"

/*
You have a RecentCounter class which counts the number of recent requests within a certain time frame.
Implement the RecentCounter class:
RecentCounter() Initializes the counter with zero recent requests.
int ping(int t) Adds a new request at time t, where t represents some time in milliseconds, and returns the number of requests that has happened in the past 3000 milliseconds (including the new request).
Specifically, return the number of requests that have happened in the inclusive range [t - 3000, t].
It is guaranteed that every call to ping uses a strictly larger value of t than the previous call.

Example 1:

Input
["RecentCounter", "ping", "ping", "ping", "ping"]
[[], [1], [100], [3001], [3002]]
Output
[null, 1, 2, 3, 3]

Explanation
RecentCounter recentCounter = new RecentCounter();
recentCounter.ping(1);     // requests = [1], range is [-2999,1], return 1
recentCounter.ping(100);   // requests = [1, 100], range is [-2900,100], return 2
recentCounter.ping(3001);  // requests = [1, 100, 3001], range is [1,3001], return 3
recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002], range is [2,3002], return 3

Constraints:

1 <= t <= 109
Each test case will call ping with strictly increasing values of t.
At most 104 calls will be made to ping.
*/
type RecentCounter struct {
	queue []int
}

func RecentCounterConstructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	if len(this.queue) == 0 {
		this.queue = append(this.queue, t)
	} else {
		for len(this.queue) > 0 {
			if t-this.queue[0] > 3000 {
				this.queue = this.queue[1:]
			} else {
				break
			}
		}
		this.queue = append(this.queue, t)
	}
	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

/*
 In the world of Dota2, there are two parties: the Radiant and the Dire.

The Dota2 senate consists of senators coming from two parties. Now the Senate wants to decide on a change in the Dota2 game.
The voting for this change is a round-based procedure. In each round, each senator can exercise one of the two rights:
Ban one senator's right: A senator can make another senator lose all his rights in this and all the following rounds.
Announce the victory: If this senator found the senators who still have rights to vote are all from the same party, he can announce the victory and decide on the change in the game.
Given a string senate representing each senator's party belonging. The character 'R' and 'D' represent the Radiant party and the Dire party.
Then if there are n senators, the size of the given string will be n.

The round-based procedure starts from the first senator to the last senator in the given order.
This procedure will last until the end of voting. All the senators who have lost their rights will be skipped during the procedure.

Suppose every senator is smart enough and will play the best strategy for his own party.
Predict which party will finally announce the victory and change the Dota2 game. The output should be "Radiant" or "Dire".

Example 1:

Input: senate = "RD"
Output: "Radiant"
Explanation:
The first senator comes from Radiant and he can just ban the next senator's right in round 1.
And the second senator can't exercise any rights anymore since his right has been banned.
And in round 2, the first senator can just announce the victory since he is the only guy in the senate who can vote.
Example 2:

Input: senate = "RDD"
Output: "Dire"
Explanation:
The first senator comes from Radiant and he can just ban the next senator's right in round 1.
And the second senator can't exercise any rights anymore since his right has been banned.
And the third senator comes from Dire and he can ban the first senator's right in round 1.
And in round 2, the third senator can just announce the victory since he is the only guy in the senate who can vote.

Constraints:

n == senate.length
1 <= n <= 104
senate[i] is either 'R' or 'D'.
*/

// D: 0 1
// R: 2 3 4
// DDRRRR
type IntQueue struct {
	data []int
}
func NewIntQueue() *IntQueue {
    return &IntQueue{}
}
func (q *IntQueue) Len() int { return len(q.data) }

func (q *IntQueue) IsEmpty() bool { return len(q.data) == 0 }

func (q *IntQueue) PushToTail(i int) { 
    q.data = append(q.data, i)
}

func (q *IntQueue) Peek() (int, error) { 
    if q.IsEmpty() {
        return 0, fmt.Errorf("empty queue")
    }
    return q.data[0], nil
}

func (q *IntQueue) PopFromHead() (int, error) { 
    if q.IsEmpty() {
        return 0, fmt.Errorf("empty queue")
    }
    rs:=q.data[0]
    if len(q.data) >1 {
        q.data = q.data[1:]
    }else {
        q.data = q.data[:0]
    }
    return rs, nil
}


func predictPartyVictory(senate string) string {
    // D R D R R D
    // queue1 = 1 3 4
    // queue2 = 0 2 5
    queue1:= make([]int,0)
    queue2:= make([]int,0)

    for i:=0; i < len(senate); i++ {
        if senate[i] == 'R' {
            queue1 = append(queue1, i)
        } else {
            queue2 = append(queue2, i)
        }
    }
    nextIndex:= len(senate)
    for len(queue1) > 0 && len(queue2) > 0 {
        if queue1[0] < queue2[0] {
            queue1 = append(queue1, nextIndex)
        }else {
            queue2 = append(queue2, nextIndex)
        }
        nextIndex++
        queue1 = queue1[1:]
        queue2 = queue2[1:]
    }

    if len(queue1) > 0 {
        return "Radiant"
    } else {
        return "Dire"
    }

}
