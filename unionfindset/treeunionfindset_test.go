package unionfindset

import (
	"testing"
)

func TestTreeUnionFindset(t *testing.T) {
	u := NewTreeFindset(10)
	if u.Find(1, 2) {
		t.Fatalf("1, 2 should be not union")
	}
	u.Union(1, 2)
	if !u.Find(1, 2) {
		t.Fatalf("1, 2 should be the union")
	}
	if u.Find(1, 3) {
		t.Fatalf("1, 3 should be not union")
	}
	u.Union(1, 3)
	if !u.Find(1, 3) {
		t.Fatalf("1, 3 should be the union")
	}
	if !u.Find(2, 3) {
		t.Fatalf("1, 3 should be the union")
	}
}