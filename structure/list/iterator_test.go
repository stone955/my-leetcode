package list

import "testing"

func TestArrayListIterator_HasNext(t *testing.T) {
	l := NewArrayList()
	for i := 0; i < 10; i++ {
		_ = l.Add(i)
	}
	t.Logf("ArrayList String(): %v\n", l.String())

	t.Logf("ArrayList Iterator(): ")
	for it := l.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		t.Logf("%v ", item)
	}
}

func TestArrayListIterator_Next(t *testing.T) {
	l := NewArrayList()
	for i := 0; i < 10; i++ {
		_ = l.Add(i)
	}
	t.Logf("ArrayList String(): %v\n", l.String())

	t.Logf("ArrayList Iterator(): ")
	for it := l.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		t.Logf("%v ", item)
	}
}

func TestArrayListIterator_Remove(t *testing.T) {
	l := NewArrayList()
	for i := 0; i < 10; i++ {
		_ = l.Add(i)
	}
	t.Logf("ArrayList Before Remove: %v\n", l.String())

	for it := l.Iterator(); it.HasNext(); {
		_, _ = it.Next()
		it.Remove()
	}
	t.Logf("ArrayList After Remove: %v\n", l.String())
}
