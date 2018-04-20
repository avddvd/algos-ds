package interviewcake

import "math"

type Statistics struct {
	Max        int32
	Min        int32
	Mean       float32
	Mode       int32
	counts     map[int32]int32
	totalElems int
}

func NewStatistics() *Statistics {
	return &Statistics{
		Max:        math.MinInt32,
		Min:        math.MaxInt32,
		Mean:       0,
		Mode:       0,
		counts:     map[int32]int32{0: 0},
		totalElems: 0,
	}
}

func (s *Statistics) Insert(newElem int) {
	elem := int32(newElem)
	s.totalElems += 1
	if elem > s.Max {
		s.Max = elem
	}
	if elem < s.Min {
		s.Min = elem
	}
	if _, ok := s.counts[elem]; !ok {
		s.counts[elem] = 1
	} else {
		s.counts[elem] += 1
	}
	if s.counts[elem] > s.counts[s.Mode] {
		s.Mode = elem
	}
	s.Mean = (s.Mean + float32(elem)) / float32(s.totalElems)
}
