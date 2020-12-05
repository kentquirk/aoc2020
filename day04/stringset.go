package main

// StringSet is intended to store a set of strings
type StringSet map[string]struct{}

// Add puts an item in the set
func (s *StringSet) Add(ary ...string) {
	for _, a := range ary {
		(*s)[a] = struct{}{}
	}
}

// Contains returns whether the value is in the set
func (s StringSet) Contains(n string) bool {
	_, ok := s[n]
	return ok
}

// Intersect creates a new StringSet that is the intersection of s and t
func (s StringSet) Intersect(t StringSet) StringSet {
	result := make(StringSet)
	for k := range s {
		if t.Contains(k) {
			result.Add(k)
		}
	}
	return result
}
