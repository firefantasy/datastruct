package kmp

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestKmpFind(t *testing.T) {
	a := "I love golang"
	b := "golang"
	i := KmpFind(a, b)
	if i != 7 {
		t.Errorf("test failed")
	}
	b = "a"
	i = KmpFind(a, b)
	if i != 10 {
		t.Errorf("test failed")
	}
}

func TestGenLPS(t *testing.T) {
	a := "abcdabc"
	lps := genLPS(a)
	Convey("should equal", t, func() {
		So(lps, ShouldResemble, []int{0, 0, 0, 0, 1, 2, 3})
	})
	a = "abcdabcabc"
	lps = genLPS(a)
	Convey("should equal", t, func() {
		So(lps, ShouldResemble, []int{0, 0, 0, 0, 1, 2, 3, 1, 2, 3})
	})
}