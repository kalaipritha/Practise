package main

import (
	"fmt"
	s "strings"
)

type priority_queue struct {
	item     string
	priority int
}

var i = 0

//i have specofied the array length make it as dynmic
var pq [3]priority_queue

func GetCount(str string, i string) int {
	return s.Count(str, i)
}

func Insert(item string, priority int) {
	k := i
	if i == 0 {
		pq[i].item = item
		pq[i].priority = priority
		i++
	} else {
		for j := i - 1; j >= 0; j-- {
			if priority > pq[j].priority {
				pq[k].item = pq[j].item
				pq[k].priority = pq[j].priority
				k--
			} else {
				pq[k].item = item
				pq[k].priority = priority
				i++
				return
			}
		}
		pq[k].item = item
		pq[k].priority = priority
		i++

	}
}

func Pop() (string, int) {
	ele := pq[0].item
	priority := pq[0].priority - 1
	for j := 1; j < i; j++ {
		pq[j-1] = pq[j]
		pq[j].item = ""
		pq[j].priority = 0
	}
	i--
	return ele, priority
}

func main() {
	m := make(map[string]int)
	j := 1
	flag := true
	str := "aaabba"
	for _, s := range str {
		if m[string(s)] == 0 {
			m[string(s)] = 1
			Insert(string(s), GetCount(str, string(s)))
		}
	}
	arr := make([]string, len(str))
	temp, temp_prio := Pop()
	arr[0] = temp
	for i >= 0 && j < len(arr) {
		ele, prio := Pop()
		if s.Compare(ele, temp) == 0 && prio > 0 {
			flag = false
			break
		}
		arr[j] = ele
		if temp_prio > 0 {
			Insert(temp, temp_prio)
		}
		temp = ele
		temp_prio = prio
		j++
	}

	if flag == true {
		fmt.Println(arr)
	} else {
		fmt.Println("not possible")
	}
}
