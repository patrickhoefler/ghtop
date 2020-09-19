package internal

import (
	"errors"
)

type sortOrder int

const (
	unknown sortOrder = iota
	ascending
	descending
)

//go:generate stringer -type=sortOrder

// Rankinator calculates the rank for you, so you don't have to
type Rankinator struct {
	lastValue      int
	sameValueCount int
	lastRank       int
	sortOrder      sortOrder
}

// Get returns the calculated rank.
// Please note that the values need to be supplied either in ascending or in descending order.
func (r *Rankinator) Get(value int) (rank int, err error) {
	if r.lastRank > 0 {
		if value > r.lastValue {
			if r.sortOrder == unknown {
				r.sortOrder = ascending
			} else if r.sortOrder == descending {
				err = errors.New("sort order changed from descending to ascending")
				return
			}
		} else if value < r.lastValue {
			if r.sortOrder == unknown {
				r.sortOrder = descending
			} else if r.sortOrder == ascending {
				err = errors.New("sort order changed from ascending to descending")
				return
			}
		}
	}

	if value == r.lastValue {
		r.sameValueCount++
		rank = r.lastRank
		return
	}

	rank = r.lastRank + r.sameValueCount + 1

	r.lastRank = rank
	r.lastValue = value
	r.sameValueCount = 0
	return
}
