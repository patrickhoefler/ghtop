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
func (rankinator *Rankinator) Get(value int) (rank int, err error) {
	setSortOrder(rankinator, value)

	if rankinator.lastRank > 0 {
		err := checkSortOrder(rankinator, value)
		if err != nil {
			return 0, err
		}
	}

	if value == rankinator.lastValue {
		rankinator.sameValueCount++
		rank = rankinator.lastRank
		return
	}

	rank = rankinator.lastRank + rankinator.sameValueCount + 1

	rankinator.lastRank = rank
	rankinator.lastValue = value
	rankinator.sameValueCount = 0
	return
}

func setSortOrder(r *Rankinator, value int) {
	if r.sortOrder != unknown {
		return
	}

	if value > r.lastValue {
		r.sortOrder = ascending
	} else if value < r.lastValue {
		r.sortOrder = descending
	}
}

func checkSortOrder(r *Rankinator, value int) error {
	if value > r.lastValue && r.sortOrder == descending {
		return errors.New("sort order changed from descending to ascending")
	}

	if value < r.lastValue && r.sortOrder == ascending {
		return errors.New("sort order changed from ascending to descending")
	}

	return nil
}
