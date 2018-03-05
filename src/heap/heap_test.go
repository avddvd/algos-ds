package heap

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	mH := NewMinHeap()
	for _, elem := range []int{5, 7, 3, 11} {
		mH.Insert(elem)
	}

	if mH.GetMin() != 3 {
		t.Error("error building heap or retrieving min")
	}
	if mH.currSize != 3 {
		t.Error("error retrieving min from min heap")
	}
	if mH.GetMin() != 5 {
		t.Error("error building heap or retrieving min")
	}
	m := mH.GetMin()
	if m != 7 {
		t.Error("error retreiving min", mH.heapList)
	}
	m = mH.GetMin()
	if m != 11 {
		t.Error("error retreiving min", m)
	}
	if mH.currSize != 0 {
		t.Error("error retreiving min", mH.currSize, mH.heapList)
	}
	mH.BuildHeap([]int{5, 7, 3, 11})
	fmt.Println(mH.heapList)
}

func TestMaxHeap(t *testing.T) {
	mH := NewMaxHeap()
	for _, elem := range []int{5, 7, 3, 11} {
		mH.Insert(elem)
	}
	fmt.Println(mH)
	if mH.GetMax() != 11 {
		t.Error("error building heap or retrieving min")
	}
	if mH.currSize != 3 {
		t.Error("error retrieving min from min heap")
	}
	if mH.GetMax() != 7 {
		t.Error("error building heap or retrieving min")
	}
	m := mH.GetMax()
	if m != 5 {
		t.Error("error retreiving min", mH.heapList)
	}
	m = mH.GetMax()
	if m != 3 {
		t.Error("error retreiving min", m)
	}
	if mH.currSize != 0 {
		t.Error("error retreiving min", mH.currSize, mH.heapList)
	}

	mH.BuildHeap([]int{5, 7, 3, 11})
	fmt.Println(mH.heapList)

}
