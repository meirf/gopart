package gopart

// IdxRange specifies a single range. Start and Stop
// are the indexes in the larger collection at which this
// range begins and ends, respectively. Note that Stop
// is exclusive, wheras Start is inclusive.
type IdxRange struct {
	Start, Stop int
}

// Partition enables type-agnostic partitioning
// of anything indexable by specifying the length and
// the desired partition size of the indexable object.
// Consecutive index ranges are sent to the channel,
// each of which is the same size. The final range may
// be smaller than the others.
//
// This method should be used in a for...range loop.
// No results will be returned if the partition size is
// nonpositive. If the partition size is greater than the
// collection length, the range returned includes the
// entire collection.
func Partition(collectionLen, partitionSize int) chan IdxRange {
	c := make(chan IdxRange)
	if partitionSize <= 0 {
		close(c)
		return c
	}

	go func() {
		numFullPartitions := collectionLen / partitionSize
		var i int
		for ; i < numFullPartitions; i++ {
			c <- IdxRange{Start: i * partitionSize, Stop: (i + 1) * partitionSize}
		}

		if collectionLen%partitionSize != 0 {
			c <- IdxRange{Start: i * partitionSize, Stop: collectionLen}
		}

		close(c)
	}()
	return c
}
