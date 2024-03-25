package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	myMap := make(map[string]int)
	myMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	values := make([]int, 0, len(myMap))
	keys := make([]string, 0, len(myMap))

	for k, v := range myMap {
		keys = append(keys, k)
		values = append(values, v)
	}

	slices.Sort(values)
	slices.Sort(keys)
	fmt.Println(values, keys)
	sort.Slice(
		values, func(i, j int) bool {
			return values[i] < values[j]
		},
	)

	sort.Slice(
		keys, func(i, j int) bool {
			return keys[i] < keys[j]
		},
	)

	fmt.Println(values, keys)
}
