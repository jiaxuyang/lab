package intensity_segments

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

// Segment is abstraction of an interval in IntensitySegments
type Segment struct {
	From   int // from value of the interval, value range is (-infinity, infinity)
	Amount int // amount of this interval
}

// IntensitySegments
// Time complexity of Add() is O(n);
// Time complexity of Set() is O(n);
// Future improvement: use an ordered map implemented by red-black tree, which will
// make the time complexity decrease to O(log n).
type IntensitySegments struct {
	sync.RWMutex            // lock for concurrency safety
	segments     []*Segment // intervals ordered(asc) by From field
}

// NewIntensitySegments constructor of IntensitySegments
func NewIntensitySegments() *IntensitySegments {
	return &IntensitySegments{}
}

// Add  amount to a left close, right open segment: [from, to).
// Constrains: from < to
func (s *IntensitySegments) Add(from, to, amount int) {
	if from >= to || amount == 0 {
		return
	}

	s.Lock()
	defer s.Unlock()

	// If there are no 'from' or 'to' nodes, add nodes, and the values of the nodes are calculated
	// based on the current segments.
	s.addSeg(from)
	s.addSeg(to)

	// Accumulate the amount for the nodes within the range [from, to).
	s.doAddAmount(from, to, amount)

	// Shrinking conditions:
	// 1. Nodes with prefix and amount = 0 are deleted;
	// 2. Consecutive equal nodes are deleted.
	s.shrink()
}

// Set amount to a left close, right open segment: [from, to).
// Similar to Add(), but set to amount ignore original value.
// Constrains: from < to
func (s *IntensitySegments) Set(from, to, amount int) {
	if from >= to {
		return
	}

	s.Lock()
	defer s.Unlock()

	// If there are no 'from' or 'to' nodes, add nodes, and the values of the nodes are calculated
	// based on the current segments.
	s.addSeg(from)
	s.addSeg(to)

	// Assign the amount for the nodes within the range [from, to).
	s.doSetAmount(from, to, amount)

	// Shrinking conditions:
	// 1. Nodes with prefix and amount = 0 are deleted;
	// 2. Consecutive equal nodes are deleted.
	s.shrink()
}

// ToString for print
func (s *IntensitySegments) ToString() string {
	// need read lock
	s.RLock()
	defer s.RUnlock()

	return s.print()
}

// Print str, no lock
func (s *IntensitySegments) print() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("[")
	for i, seg := range s.segments {
		if i != 0 {
			strBuilder.WriteString(",")
		}
		strBuilder.WriteString(fmt.Sprintf("[%d,%d]", seg.From, seg.Amount))
	}
	strBuilder.WriteString("]")
	return strBuilder.String()
}

// Filter [from, to) and add amount
func (s *IntensitySegments) doAddAmount(from, to, amount int) {
	for i := range s.segments {
		if key := s.segments[i].From; key >= from && key < to {
			s.segments[i].Amount += amount // do add
		}
	}
}

// Filter [from, to) and set amount
func (s *IntensitySegments) doSetAmount(from, to, amount int) {
	for i := range s.segments {
		if key := s.segments[i].From; key >= from && key < to {
			s.segments[i].Amount = amount // do set
		}
	}
}

// shrink the segments if needed
func (s *IntensitySegments) shrink() {
	// The shrink result saved here
	var shrink []*Segment
	// IsPrefixZero in order to mark leading zero segments
	isPrefixZero := true
	for i := range s.segments {
		// 1. Nodes with prefix and amount = 0 are deleted;
		if s.segments[i].Amount == 0 && isPrefixZero {
			continue
		}
		isPrefixZero = false

		// 2. Consecutive equal nodes are deleted.
		if i-1 >= 0 {
			if s.segments[i].Amount == s.segments[i-1].Amount {
				continue
			}
		}

		shrink = append(shrink, s.segments[i])
	}
	s.segments = shrink
}

// insert a new seg if not exist
func (s *IntensitySegments) addSeg(key int) {
	// Binary find: find key or first greater than key.
	// Time complexity is log(n).
	greater, exist := sort.Find(len(s.segments), func(i int) int {
		return key - s.segments[i].From
	})
	if exist {
		// Doesn't need to do anything
		return
	}

	// Calculate the original amount of this interval
	baseAmount := 0
	if greater-1 >= 0 {
		baseAmount = s.segments[greater-1].Amount
	}

	// Insert seg to slice
	seg := &Segment{From: key, Amount: baseAmount}
	s.segments = append(s.segments[:greater], append([]*Segment{seg}, s.segments[greater:]...)...)
}

// AddHelper for test
func AddHelper(inputs [][]int) []string {
	s := NewIntensitySegments()
	output := make([]string, 0)
	for _, input := range inputs {
		s.Add(input[0], input[1], input[2])
		output = append(output, s.ToString())
	}
	return output
}

// SetHelper for test
func SetHelper(inputs [][]int) []string {
	s := NewIntensitySegments()
	output := make([]string, 0)
	for _, input := range inputs {
		s.Set(input[0], input[1], input[2])
		output = append(output, s.ToString())
	}
	return output
}
