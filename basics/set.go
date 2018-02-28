// The set package implements a set of integers.
package set

import (
	"sort"
)

// Set defines a new set of integers.
type IntSet struct {
	members map[int]struct{}
}

// Add places a new member in the set.
func (s *IntSet) Add(i int) {
	_, ok := s.members[i]
	if !ok {
		s.members[i] = struct{}{}
	}
}

// Remove takes a member out of the set.
func (s *IntSet) Remove(i int) {
	delete(s.members, i)
}

// Contains returns true if i is a member of the set.
func (s *IntSet) Contains(i int) bool {
	_, ok := s.members[i]

	return ok
}

// Members returns a list of the members of the set.
func (s *IntSet) Members() []int {
	var members []int

	for i := range s.members {
		members = append(members, i)
	}

	sort.Ints(members)

	return members
}

// Size returns the number of members in the set.
func (s *IntSet) Size() int {
	return len(s.members)
}

// Equal returns true if the given set is the same length and contains all of the
// same members as this set.
func (s *IntSet) Equal(y *IntSet) bool {
	if s.Size() != y.Size() {
		return false
	}

	for _, m := range s.Members() {
		if !y.Contains(m) {
			return false
		}
	}

	return true
}

// Union returns a new set that contains all of the members of this set and
// all of the members of the given set y.
func (s *IntSet) Union(y *IntSet) *IntSet {
	u := NewIntSet([]int{})

	for _, m := range s.Members() {
		u.Add(m)
	}

	for _, m := range y.Members() {
		u.Add(m)
	}

	return u
}

// Intersection returns a new set that contains only those members that appear
// in both this set and the given set y.
func (s *IntSet) Intersection(y *IntSet) *IntSet {
	i := NewIntSet([]int{})

	for _, m := range s.Members() {
		if y.Contains(m) {
			i.Add(m)
		}
	}

	return i
}

// Difference returns a new set that contains only the members of this set
// that are not also members of the given set y.
func (s *IntSet) Difference(y *IntSet) *IntSet {
	d := NewIntSet([]int{})

	for _, m := range s.Members() {
		if !y.Contains(m) {
			d.Add(m)
		}
	}

	return d
}

// NewIntSet returns a new set of integers containing the unique integers in
// the members slice.
func NewIntSet(members []int) *IntSet {
	i := new(IntSet)
	i.members = make(map[int]struct{})

	for _, m := range members {
		i.Add(m)
	}

	return i
}

// Set defines a new set of strings.
type StringSet struct {
	members map[string]struct{}
}

// Add places a new member in the set.
func (s *StringSet) Add(i string) {
	_, ok := s.members[i]
	if !ok {
		s.members[i] = struct{}{}
	}
}

// Remove takes a member out of the set.
func (s *StringSet) Remove(i string) {
	delete(s.members, i)
}

// Contains returns true if i is a member of the set.
func (s *StringSet) Contains(i string) bool {
	_, ok := s.members[i]

	return ok
}

// Members returns a list of the members of the set.
func (s *StringSet) Members() []string {
	var members []string

	for i := range s.members {
		members = append(members, i)
	}

	sort.Strings(members)

	return members
}

// Size returns the number of members in the set.
func (s *StringSet) Size() int {
	return len(s.members)
}

// Equal returns true if the given set is the same length and contains all of the
// same members as this set.
func (s *StringSet) Equal(y *StringSet) bool {
	if s.Size() != y.Size() {
		return false
	}

	for _, m := range s.Members() {
		if !y.Contains(m) {
			return false
		}
	}

	return true
}

// Union returns a new set that contains all of the members of this set and
// all of the members of the given set y.
func (s *StringSet) Union(y *StringSet) *StringSet {
	u := NewStringSet([]string{})

	for _, m := range s.Members() {
		u.Add(m)
	}

	for _, m := range y.Members() {
		u.Add(m)
	}

	return u
}

// Intersection returns a new set that contains only those members that appear
// in both this set and the given set y.
func (s *StringSet) Intersection(y *StringSet) *StringSet {
	i := NewStringSet([]string{})

	for _, m := range s.Members() {
		if y.Contains(m) {
			i.Add(m)
		}
	}

	return i
}

// Difference returns a new set that contains only the members of this set
// that are not also members of the given set y.
func (s *StringSet) Difference(y *StringSet) *StringSet {
	d := NewStringSet([]string{})

	for _, m := range s.Members() {
		if !y.Contains(m) {
			d.Add(m)
		}
	}

	return d
}

// NewStringSet returns a new set of integers containing the unique strings in
// the members slice.
func NewStringSet(members []string) *StringSet {
	i := new(StringSet)
	i.members = make(map[string]struct{})

	for _, m := range members {
		i.Add(m)
	}

	return i
}
