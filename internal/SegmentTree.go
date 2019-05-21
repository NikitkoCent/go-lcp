package SegmentTree

type SegmentTree interface {
	Get(minIndex, maxIndex uint64) uint64
}

func MakeMinSegmentTree(dataArray []uint64) SegmentTree {
	result := minSegmentTree{}
	result.size = uint64(len(dataArray))
	result.tree = make([]uint64, result.size*4)

	if result.size > 0 {
		result.initializeFrom(dataArray, 0)
	}

	return result
}

// Implementation

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

type minSegmentTree struct {
	tree []uint64
	size uint64
}

func (tree minSegmentTree) initializeFrom(dataArray []uint64, currentIndex uint64) {
	length := uint64(len(dataArray))

	if length == 1 {
		tree.tree[currentIndex] = dataArray[0]
	} else {
		dataMiddleIndex := (length-1)/2 + 1

		leftChildIndex := 2*currentIndex + 1
		rightChildIndex := leftChildIndex + 1

		tree.initializeFrom(dataArray[:dataMiddleIndex], leftChildIndex)
		tree.initializeFrom(dataArray[dataMiddleIndex:], rightChildIndex)

		tree.tree[currentIndex] = min(tree.tree[leftChildIndex], tree.tree[rightChildIndex])
	}
}

func (tree minSegmentTree) Get(minIndex, maxIndex uint64) uint64 {
	return tree.getImpl(minIndex, maxIndex, 0, 0, tree.size-1)
}

func (tree minSegmentTree) getImpl(minIndex, maxIndex, vertexIndex, rangeMinIndex, rangeMaxIndex uint64) uint64 {
	if (minIndex == rangeMinIndex) && (maxIndex == rangeMaxIndex) {
		return tree.tree[vertexIndex]
	}

	rangeMiddle := (rangeMinIndex + rangeMaxIndex) / 2

	if maxIndex <= rangeMiddle {
		return tree.getImpl(minIndex, maxIndex, vertexIndex*2+1, rangeMinIndex, rangeMiddle)
	} else if minIndex > rangeMiddle {
		return tree.getImpl(minIndex, maxIndex, vertexIndex*2+2, rangeMiddle+1, rangeMaxIndex)
	}

	return min(tree.getImpl(minIndex, rangeMiddle, vertexIndex*2+1, rangeMinIndex, rangeMiddle),
		tree.getImpl(rangeMiddle+1, maxIndex, vertexIndex*2+2, rangeMiddle+1, rangeMaxIndex))
}
