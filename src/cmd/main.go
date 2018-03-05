/*
	Package main.

	Practice of effective go with algorithms and datastructures
*/
package main

import (
	"fmt"
	//"github.com/avddvd/algos-ds/analysis"
	"github.com/avddvd/algos-ds/list"
	"github.com/avddvd/algos-ds/queue"
	//"github.com/avddvd/algos-ds/stack"
	//"github.com/avddvd/algos-ds/reverse"
	"github.com/avddvd/algos-ds/recursion"
	"github.com/avddvd/algos-ds/sorting"
)

func main() {
	fmt.Println(queue.RadixSort([]int{100, 199, 175, 156, 180, 103, 126}))
	l := new(list.List)
	for _, value := range "abcde" {
		l.Append(value)
	}
	fmt.Println(l.Size(), l)
	l1, l2 := l.Split()
	fmt.Println(l1, l1.Size())
	fmt.Println(l2, l2.Size())
	fmt.Println(l, l.Size())
	fmt.Println(recursion.PalindromeChecker("aibohphobia"))
	fmt.Println(recursion.ConstructMaze(4))
	maze := recursion.GetMaze()
	for _, r := range maze {
		for _, c := range r {
			fmt.Printf("%s ", c)
		}
		fmt.Printf("\n")
	}
	fmt.Println(recursion.EscapeMaze(maze, 2, 1))
	change := 63
	coinVals := []int{1, 2, 5, 10, 25}
	usedCoins, minCoins := recursion.CoinChange(change, coinVals)
	fmt.Println(recursion.GetCoinsUsed(usedCoins, change))
	fmt.Println(minCoins[change])
	arr := []int{5, 4, 3, 2, 1, 0, 100, -20}
	fmt.Println(sorting.MergeSort(arr))
	sorting.QuickSort(arr)
	fmt.Println(arr)
}
