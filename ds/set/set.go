package set

import (
	e "aoc2020/err"
	"strconv"
)

// sets generally

type void struct{}

var member void

func IsMember[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

func CopySet[K comparable](set map[K]void) map[K]void {
	copiedSet := make(map[K]void)
	for k := range set {
		copiedSet[k] = member
	}
	return copiedSet
}

// IntSets

type IntSet = map[int]void

func StrArrToIntSet(arr []string) IntSet {
	set := make(IntSet)
	for _, val := range arr {
		num, err := strconv.Atoi(val)
		e.Check(err)
		set[num] = member

	}
	return set
}
