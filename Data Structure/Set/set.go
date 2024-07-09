package main

import (
	"strconv"
	"strings"
)

// Set  is a struct that contains a map of type map[int]bool which is used to store the set of unique integers.
// NewSet() is a function that creates and returns a new instance of the Set  struct.
// Add(val int) adds the given integer value to the set.
// Contains(val int) checks if the set contains the given value.
// Remove(val int) removes the specified value from the set.
// Size() returns the number of elements in the set.
// IsEmpty() checks if the set is empty.
// Clear() clears the Set by creating a new empty map and assigning it to the set field of the receiver.
// Copy() creates a new Set by copying all the elements from the current Set.
// Values() returns all the values in the set as a slice of integers.
// IsSubset(other *Set) checks if the current set is a subset of the given set.
// IsSuperset(other *Set) checks if the given set is a superset of the current set.
// IsDisjoint(other *Set) checks if the given set is disjoint with the current set.
// Equal(other *Set) checks if two sets are equal by verifying if each is a subset of the other.
// Difference(other *Set) returns a new Set containing the elements that are in the receiver Set but not in the other Set.
// Intersection(other *Set) returns a new Set containing the elements that are in both the receiver Set and the other Set.
// Union(other *Set) computes the union of two sets.
// IterativeSet() iterates over the elements of the Set and prints them.
type Set struct {
	set map[int]bool
}

// NewSet creates and returns a new Set.
//
// No parameters.
// Returns a pointer to a Set.
func NewSet() *Set {
	return &Set{set: make(map[int]bool)}
}

// Add adds the given integer value to the set.
//
// Parameters:
// - val: the integer value to be added to the set.
//
// Return type: none.
func (s *Set) Add(val int) {
	if !s.set[val] {
		s.set[val] = true
	}
}

// Contains checks if the set contains the given value.
//
// Parameters:
// - val: the value to check.
//
// Returns:
// - true if the set contains the value, false otherwise.
func (s *Set) Contains(val int) bool {
	return s.set[val]
}

// Remove removes the specified value from the set.
//
// Parameters:
// - val: the value to be removed from the set.
//
// Return type: none.
func (s *Set) Remove(val int) {
	delete(s.set, val)
}

// Size returns the number of elements in the set.
//
// This function does not take any parameters.
// It returns an integer representing the size of the set.
func (s *Set) Size() int {
	return len(s.set)
}

// IsEmpty checks if the set is empty.
//
// It returns true if the set is empty, false otherwise.
func (s *Set) IsEmpty() bool {
	return len(s.set) == 0
}

// Clear clears the Set by creating a new empty map and assigning it to the set field of the receiver.
//
// No parameters.
// Returns nothing.
func (s *Set) Clear() {
	s.set = make(map[int]bool)
}

// Copy creates a new Set by copying all the elements from the current Set.
//
// It does not take any parameters.
// It returns a pointer to a new Set.
func (s *Set) Copy() *Set {
	newSet := NewSet()
	for val := range s.set {
		newSet.Add(val)
	}
	return newSet
}

// Values returns all the values in the set as a slice of integers.
//
// No parameters.
// Returns a slice of integers.
func (s *Set) Values() []int {
	values := make([]int, 0, len(s.set))
	for val := range s.set {
		values = append(values, val)
	}
	return values
}

// Union computes the union of two sets.
//
// Parameters:
// - other: the other Set to compute the union with.
// Returns:
// - A pointer to a new Set representing the union of the two sets.
func (s *Set) Union(other *Set) *Set {
	union := NewSet()
	for val := range s.set {
		union.Add(val)
	}
	for val := range other.set {
		union.Add(val)
	}
	return union
}

// Intersection returns a new Set containing the elements that are common to both the receiver Set and the other Set.
//
// Parameters:
// - other: a pointer to the other Set to intersect with.
//
// Returns:
// - a pointer to a new Set containing the common elements.
func (s *Set) Intersection(other *Set) *Set {
	intersection := NewSet()
	for val := range s.set {
		if other.Contains(val) {
			intersection.Add(val)
		}
	}
	return intersection
}

// Difference returns a new Set containing the elements that are in the receiver Set but not in the other Set.
//
// Parameters:
// - other: a pointer to the other Set to compare with.
//
// Returns:
// - a pointer to a new Set containing the elements that are in the receiver Set but not in the other Set.
func (s *Set) Difference(other *Set) *Set {
	difference := NewSet()
	for val := range s.set {
		if !other.Contains(val) {
			difference.Add(val)
		}
	}
	return difference
}

// SymmetricDifference computes the symmetric difference between two sets.
//
// The symmetric difference of two sets is the set of elements that are in either of the sets but not in their intersection.
//
// Parameters:
// - other: a pointer to the other Set to compute the symmetric difference with.
//
// Returns:
// - a pointer to a new Set representing the symmetric difference of the two sets.
func (s *Set) SymmetricDifference(other *Set) *Set {
	return s.Difference(other).Union(other.Difference(s))
}

// IsSubset checks if the current set is a subset of the given set.
//
// Parameters:
// - other: a pointer to the other Set to check if it contains all the elements of the current Set.
//
// Returns:
// - true if the current set is a subset of the given set, false otherwise.
func (s *Set) IsSubset(other *Set) bool {
	for val := range s.set {
		if !other.Contains(val) {
			return false
		}
	}
	return true
}

// IsSuperset checks if the given set is a superset of the current set.
//
// Parameters:
// - other: a pointer to the other Set to check if it is a superset of the current Set.
//
// Returns:
// - true if the given set is a superset of the current set, false otherwise.
func (s *Set) IsSuperset(other *Set) bool {
	return other.IsSubset(s)
}

// Equal checks if two sets are equal by verifying if each is a subset of the other.
//
// Parameters:
// - other: a pointer to the other Set to compare equality with.
// Returns:
// - true if the sets are equal, false otherwise.
func (s *Set) Equal(other *Set) bool {
	return s.IsSubset(other) && other.IsSubset(s)
}

// IsDisjoint checks if the given set is disjoint with the current set.
//
// Parameters:
// - other: a pointer to the other Set to check disjointness with.
//
// Returns:
// - true if the sets are disjoint, false otherwise.
func (s *Set) IsDisjoint(other *Set) bool {
	return s.Intersection(other).IsEmpty()
}

// String returns a string representation of the Set.
//
// It iterates over the set and converts each value to a string using the strconv.Itoa() function.
// The resulting strings are then joined together with a comma separator and enclosed in curly braces.
//
// Returns:
// - A string representation of the Set.
func (s *Set) String() string {
	values := make([]string, 0, len(s.set))
	for val := range s.set {
		values = append(values, strconv.Itoa(val))
	}
	return "{" + strings.Join(values, ", ") + "}"
}

// IterativeSet iterates over the elements of the Set and prints them.
//
// No parameters.
// No return type.
func (s *Set) IterativeSet() {
	for val := range s.set {
		println(val)
	}
}
