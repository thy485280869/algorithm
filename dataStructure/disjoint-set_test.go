package dataStructure

import "testing"

func TestNewUF(t *testing.T) {
	u := NewUF(3)
	if _, ok := u.(*UF); !ok {
		t.Error()
	}
}
