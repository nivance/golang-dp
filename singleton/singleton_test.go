package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	s1 := New()
	s2 := New()
	if s1 != s2 {
		t.Errorf("create singleton failed")
	}
}
