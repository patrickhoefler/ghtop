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
// Note that the values need to be supplied either in ascending or in descending order.
func (rankinator *Rankinator) Get(value int) (rank int, err error) {
	rankinator.setSortOrder(value)

	err = rankinator.checkSortOrder(value)
	if err != nil {
		return
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

func (rankinator *Rankinator) setSortOrder(value int) {
	// Don't set the order on the first run
	if rankinator.lastRank == 0 {
		return
	}

	// Don't change the sort order once it has been set
	if rankinator.sortOrder != unknown {
		return
	}

	if value > rankinator.lastValue {
		rankinator.sortOrder = ascending
	} else if value < rankinator.lastValue {
		rankinator.sortOrder = descending
	}
}

func (rankinator *Rankinator) checkSortOrder(value int) error {
	if value > rankinator.lastValue && rankinator.sortOrder == descending {
		return errors.New("sort order changed from descending to ascending")
	}

	if value < rankinator.lastValue && rankinator.sortOrder == ascending {
		return errors.New("sort order changed from ascending to descending")
	}

	return nil
}
