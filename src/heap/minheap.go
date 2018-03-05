package heap

type MinHeap struct {
	heapList []int
	currSize int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		heapList: []int{0},
		currSize: 0,
	}
}

func (minH *MinHeap) Insert(data int) {
	minH.heapList = append(minH.heapList, data)
	minH.currSize += 1
	minH.PercUp(minH.currSize)
}

func (minH *MinHeap) GetMin() int {
	min := minH.heapList[1]
	minH.heapList[1] = minH.heapList[minH.currSize]
	minH.heapList = minH.heapList[0:minH.currSize]
	minH.currSize -= 1
	minH.PercDown(1)
	return min
}

func (minH *MinHeap) PercDown(i int) {
	for i*2 <= minH.currSize {
		mC := minH.GetMinChild(i)
		if minH.heapList[mC] < minH.heapList[i] {
			minH.heapList[mC], minH.heapList[i] = minH.heapList[i], minH.heapList[mC]
		}
		i = i * 2
	}
}

func (minH *MinHeap) GetMinChild(i int) int {
	var minIndex int
	if 2*i+1 > minH.currSize {
		minIndex = 2 * i
	} else {
		if minH.heapList[2*i] > minH.heapList[2*i+1] {
			minIndex = 2*i + 1
		} else {
			minIndex = 2 * i
		}
	}
	return minIndex
}

func (minH *MinHeap) PercUp(i int) {
	for i/2 > 0 {
		if minH.heapList[i] < minH.heapList[i/2] {
			minH.heapList[i], minH.heapList[i/2] = minH.heapList[i/2], minH.heapList[i]
		}
		i = i / 2
	}
}

func (minH *MinHeap) BuildHeap(list []int) {
	i := len(list) / 2
	for _, elem := range list {
		minH.heapList = append(minH.heapList, elem)
	}
	minH.currSize = len(list)
	for i > 0 {
		minH.PercDown(i)
		i = i - 1
	}
}
