package easy

import "testing"

func TestReverse(t *testing.T) {
	test := struct {
		x int
		y int
	}{
		x: 645273,
		y: 372546,
	}
	rev := Reverse(test.x)
	if rev != test.y {
		t.Fatalf("%d reversed is %d failed", test.x, rev)
	}
	t.Logf("%d reversed is %d success\n", test.x, rev)
}
