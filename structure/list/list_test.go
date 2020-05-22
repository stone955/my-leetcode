package list

import (
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	l := NewArrayList()
	for i := 0; i < 20; i++ {
		_ = l.Add(i)
	}
	t.Logf("ArrayList String(): %v\n", l.String())
}

func TestArrayList_Get(t *testing.T) {
	l := NewArrayList()
	v, err := l.Get(1)
	if err != nil {
		t.Fatalf("ArrayList Get(1) error: %v\n", err)
	}
	t.Logf("ArrayList Get(1): %v\n", v)
}

func TestArrayList_Set(t *testing.T) {
	l := NewArrayList()
	_ = l.Add(1)
	_ = l.Add(2)
	_ = l.Add(3)
	t.Logf("ArrayList Before Set: %v\n", l.String())

	_ = l.Set(2, 33)
	t.Logf("ArrayList After Set: %v\n", l.String())
}

func TestArrayList_Size(t *testing.T) {
	l := NewArrayList()
	t.Logf("ArrayList Size(): %v\n", l.Size())
}

func TestArrayList_Clear(t *testing.T) {
	l := NewArrayList()
	_ = l.Add(1)
	_ = l.Add(2)
	_ = l.Add(3)
	t.Logf("ArrayList Before Clear: %v\n", l.String())

	l.Clear()
	t.Logf("ArrayList After Clear: %v\n", l.String())
}

func TestArrayList_Delete(t *testing.T) {
	l := NewArrayList()
	_ = l.Add(1)
	_ = l.Add(2)
	_ = l.Add(3)
	_ = l.Add(4)
	_ = l.Add(5)
	_ = l.Add(6)
	_ = l.Add(7)
	_ = l.Add(8)
	_ = l.Add(9)
	_ = l.Add(10)
	t.Logf("ArrayList Before Delete: %v\n", l.String())

	_ = l.Delete(1)
	t.Logf("ArrayList After Delete: %v\n", l.String())
}

func TestArrayList_String(t *testing.T) {
	l := NewArrayList()
	t.Logf("ArrayList String(): %v\n", l.String())
}

func TestDoubleLinkedList_Add(t *testing.T) {

}

func TestDoubleLinkedList_Get(t *testing.T) {

}

func TestDoubleLinkedList_Set(t *testing.T) {

}

func TestDoubleLinkedList_Size(t *testing.T) {

}

func TestLinkedList_Add(t *testing.T) {

}

func TestLinkedList_Get(t *testing.T) {

}

func TestLinkedList_Set(t *testing.T) {

}

func TestLinkedList_Size(t *testing.T) {

}
