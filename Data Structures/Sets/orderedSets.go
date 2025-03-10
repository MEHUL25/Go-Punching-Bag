package sets

import (
	"fmt"
	"sort"
)

// IntOrderedSet structure
type IntOrderedSet struct {
	data map[int]struct{}
}

// NewIntOrderedSet creates a new IntOrderedSet
func NewIntOrderedSet() *IntOrderedSet {
	return &IntOrderedSet{
		data: make(map[int]struct{}),
	}
}

// Add adds an integer to the set and ensures the set remains sorted
func (s *IntOrderedSet) Add(item int) {
	s.data[item] = struct{}{}
}

// Remove removes an integer from the set
func (s *IntOrderedSet) Remove(item int) {
	delete(s.data, item)
}

// Contains checks if an integer exists in the set
func (s *IntOrderedSet) Contains(item int) bool {
	_, exists := s.data[item]
	return exists
}

// Items returns all integers in the set in sorted order
func (s *IntOrderedSet) Items() []int {
	keys := make([]int, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

// Size returns the number of items in the set
func (s *IntOrderedSet) Size() int {
	return len(s.data)
}

func exampleOrdered() {
	set := NewIntOrderedSet()
	set.Add(5)
	set.Add(3)
	set.Add(8)
	set.Add(5) // Duplicate addition, should be ignored

	fmt.Println("Set:", set.Items()) // [5, 3, 8]

	set.Remove(3)
	fmt.Println("Set after removing 3:", set.Items()) // [5, 8]

	fmt.Println("Contains 5?", set.Contains(5))   // true
	fmt.Println("Contains 10?", set.Contains(10)) // false
}
