package singleton

import (
	"testing"
)

func TestSingleton2(t *testing.T) {
	s1 := NewInstance()
	s2 := NewInstance()
	if s1 != s2 {
		t.Errorf("create singleton failed")
	}
}
