package list

import "errors"

type Iterator interface {
	HasNext() bool              // 是否有下一个
	Next() (interface{}, error) // 下一个
	Remove()                    // 删除
}

type Iterable interface {
	Iterator() Iterator
}

type ArrayListIterator struct {
	list      *ArrayList // 数组指针
	currIndex int        // 当前索引
}

func (it *ArrayListIterator) HasNext() bool {
	return it.currIndex >= 0 && it.currIndex < it.list.size
}

func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("has not next")
	}
	value, err := it.list.Get(it.currIndex)
	if err != nil {
		return nil, err
	}
	it.currIndex++
	return value, nil
}

func (it *ArrayListIterator) Remove() {
	it.currIndex--
	_ = it.list.Delete(it.currIndex)
}
