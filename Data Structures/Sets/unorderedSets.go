package sets

import "fmt"

// IntUnorderedSet structure
type IntUnorderedSet struct {
	data map[int]struct{}
}

// NewIntUnorderedSet creates a new IntUnorderedSet
func NewIntUnorderedSet() *IntUnorderedSet {
	return &IntUnorderedSet{
		data: make(map[int]struct{}),
	}
}

// Add adds an integer to the set
func (s *IntUnorderedSet) Add(item int) {
	s.data[item] = struct{}{}
}

// Remove removes an integer from the set
func (s *IntUnorderedSet) Remove(item int) {
	delete(s.data, item)
}

// Contains checks if an integer exists in the set
func (s *IntUnorderedSet) Contains(item int) bool {
	_, exists := s.data[item]
	return exists
}

// Items returns all integers in the set (unordered)
func (s *IntUnorderedSet) Items() []int {
	keys := make([]int, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}
	return keys
}

// Size returns the number of items in the set
func (s *IntUnorderedSet) Size() int {
	return len(s.data)
}

func exampleUnordered() {
	set := NewIntUnorderedSet()
	set.Add(5)
	set.Add(3)
	set.Add(8)
	set.Add(5) // Duplicate addition, should be ignored

	fmt.Println("Set:", set.Items()) // Might print [5, 3, 8] or [3, 5, 8] or any other order

	set.Remove(3)
	fmt.Println("Set after removing 3:", set.Items()) // Might print [5, 8] or [8, 5] or any other order

	fmt.Println("Contains 5?", set.Contains(5))   // true
	fmt.Println("Contains 10?", set.Contains(10)) // false
}
