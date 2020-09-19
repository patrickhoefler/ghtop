package internal

import (
	"reflect"
	"testing"
)

func TestRankinator_Get(t *testing.T) {
	type fields struct {
		lastValue      int
		sameValueCount int
		lastRank       int
	}

	type args struct {
		value int
	}

	tests := []struct {
		name      string
		fields    fields
		args      args
		values    []int
		wantRanks []int
		expectErr bool
	}{
		{
			name:      "equal values in the middle",
			values:    []int{42, 69, 69, 123},
			wantRanks: []int{1, 2, 2, 4},
		},
		{
			name:      "equal values first",
			values:    []int{42, 42, 69, 123},
			wantRanks: []int{1, 1, 3, 4},
		},
		{
			name:      "equal values last",
			values:    []int{42, 69, 123, 123},
			wantRanks: []int{1, 2, 3, 3},
		},
		{
			name:      "multiple equal values",
			values:    []int{10, 42, 42, 69, 123, 123, 123, 314},
			wantRanks: []int{1, 2, 2, 4, 5, 5, 5, 8},
		},
		{
			name:      "values not ordered, initially ascending",
			values:    []int{42, 123, 69},
			expectErr: true,
		},
		{
			name:      "values not ordered, initially descending",
			values:    []int{123, 42, 69},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rankinator{
				lastValue:      tt.fields.lastValue,
				sameValueCount: tt.fields.sameValueCount,
				lastRank:       tt.fields.lastRank,
			}

			gotRanks := []int{}

			var err error

			for _, value := range tt.values {
				var rank int

				rank, err = r.Get(value)
				if err != nil {
					break
				}

				gotRanks = append(gotRanks, rank)
			}

			if tt.expectErr && err == nil {
				t.Fail()
				t.Logf("Exptected an error, but non occurred")
			} else if !tt.expectErr && err != nil {
				t.Fail()
				t.Logf("Expected no error, but got: %s", err)
			}

			if len(tt.wantRanks) > 0 && !reflect.DeepEqual(tt.wantRanks, gotRanks) {
				t.Errorf("Rankinator.Get() = %v, want %v", gotRanks, tt.wantRanks)
			}
		})
	}
}
