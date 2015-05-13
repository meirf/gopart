package gopart

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	var partitionTests = []struct {
		collectionLen int
		partitionSize int
		expRanges     []IdxRange
	}{
		// evenly split
		{9, 3, []IdxRange{{0, 3}, {3, 6}, {6, 9}}},
		// uneven partition
		{13, 5, []IdxRange{{0, 5}, {5, 10}, {10, 13}}},
		// large partition size
		{13, 19, []IdxRange{{0, 13}}},
		// nonpositive partiition size
		{7, 0, nil},
		// same size
		{19, 19, []IdxRange{{0, 19}}},
		// zero collection length
		{0, 19, nil},
	}

	for _, tt := range partitionTests {
		actChannel := Partition(tt.collectionLen, tt.partitionSize)
		var actRange []IdxRange
		for idxRange := range actChannel {
			actRange = append(actRange, idxRange)
		}

		if !reflect.DeepEqual(actRange, tt.expRanges) {
			t.Errorf("expected %d, actual %d", actRange, tt.expRanges)
		}
	}
}

func Benchmark100kPartitions(b *testing.B) { benchmarkParition(17*1e5+11, b) }
func Benchmark10kPartitions(b *testing.B)  { benchmarkParition(17*1e4+11, b) }
func Benchmark1kPartitions(b *testing.B)   { benchmarkParition(17*1e3+11, b) }
func Benchmark100Partitions(b *testing.B)  { benchmarkParition(17*1e2+11, b) }
func Benchmark10Partitions(b *testing.B)   { benchmarkParition(17*1e1+11, b) }

func benchmarkParition(length int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _ = range Partition(length, 17) {
		}
	}
}
