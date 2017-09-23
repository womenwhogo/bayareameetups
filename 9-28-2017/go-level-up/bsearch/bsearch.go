package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 1 {
		log.Fatalf("Usage: <search> pass an integer to search")
	}
	searchFor, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("Usage: <search> pass an integer to search")
		os.Exit(1)
	}
	arr := []int{5, 2, 6, 3, 1, 4, 23} // unsorted
	sort.Ints(arr)
	fmt.Println(arr)
	found := false
	iter := 0
	low, mid, high := 0, 0, len(arr)-1
	for low <= high {
		mid = (low + high) / 2
		if searchFor == arr[mid] {
			found = true
			break
		} else if searchFor > arr[mid] {
			low = low + 1
		} else if searchFor < arr[mid] {
			high = high - 1
		}
		iter += 1
		fmt.Println(low, mid, high, arr[mid])
	}
	if found == true {
		fmt.Println("Found:", mid, arr[mid])
	}
}
