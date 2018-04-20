package interviewcake

import "testing"
import "fmt"

func TestStatistics(t *testing.T) {
	s := NewStatistics()
	for i := -5; i <= 5; i++ {
		s.Insert(i)
	}
	s.Insert(1)
	if s.Max != 5 {
		t.Error("error reading Max")
	}
	if s.Min != -5 {
		t.Error("error reading Min")
	}
	if s.Mode != 1 {
		t.Error("error reading Mode")
	}
	fmt.Println(s)
}
