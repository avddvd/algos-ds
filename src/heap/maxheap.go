package heap

type MaxHeap struct {
	heapList []int
	currSize int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		heapList: []int{0},
		currSize: 0,
	}
}

func (maxH *MaxHeap) Insert(data int) {
	maxH.heapList = append(maxH.heapList, data)
	maxH.currSize += 1
	maxH.PercUp(maxH.currSize)
}

func (maxH *MaxHeap) GetMax() int {
	max := maxH.heapList[1]
	maxH.heapList[1] = maxH.heapList[maxH.currSize]
	maxH.heapList = maxH.heapList[0:maxH.currSize]
	maxH.currSize -= 1
	maxH.PercDown(1)
	return max
}

func (maxH *MaxHeap) PercDown(i int) {
	for i*2 <= maxH.currSize {
		mC := maxH.GetMaxChild(i)
		if maxH.heapList[mC] > maxH.heapList[i] {
			maxH.heapList[mC], maxH.heapList[i] = maxH.heapList[i], maxH.heapList[mC]
		}
		i = i * 2
	}
}

func (maxH *MaxHeap) GetMaxChild(i int) int {
	var maxIndex int
	if 2*i+1 > maxH.currSize {
		maxIndex = 2 * i
	} else {
		if maxH.heapList[2*i] < maxH.heapList[2*i+1] {
			maxIndex = 2*i + 1
		} else {
			maxIndex = 2 * i
		}
	}
	return maxIndex
}

func (maxH *MaxHeap) PercUp(i int) {
	for i/2 > 0 {
		if maxH.heapList[i] > maxH.heapList[i/2] {
			maxH.heapList[i], maxH.heapList[i/2] = maxH.heapList[i/2], maxH.heapList[i]
		}
		i = i / 2
	}
}

func (maxH *MaxHeap) BuildHeap(list []int) {
	i := len(list) / 2
	for _, elem := range list {
		maxH.heapList = append(maxH.heapList, elem)
	}
	maxH.currSize = len(list)
	for i > 0 {
		maxH.PercDown(i)
		i = i - 1
	}
}
