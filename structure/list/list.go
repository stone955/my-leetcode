package list

import (
	"errors"
	"fmt"
)

// List 线性表接口
type List interface {
	Size() int
	Get(index int) (interface{}, error)
	Add(value interface{}) error            // 追加
	Set(index int, value interface{}) error // 更新
	Clear()
	Delete(index int) error
	String() string
	Iterable
}

// ArrayList 基于数组的线性表
type ArrayList struct {
	data []interface{} // 存储数据的数组
	size int           // 数组的大小
}

func NewArrayList() List {
	return &ArrayList{
		data: make([]interface{}, 10),
		size: 0,
	}
}

func (l *ArrayList) Size() int {
	return l.size
}

func (l *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("array out of index")
	}
	return l.data[index], nil
}

func (l *ArrayList) Add(value interface{}) error {
	// 检查数组容量是否足够
	l.ensureCapacity()
	l.data[l.size] = value
	l.size++
	return nil
}

func (l *ArrayList) ensureCapacity() {
	// 扩容
	if l.size == cap(l.data) {
		// 开辟新内存，长度是原数组的2倍
		newData := make([]interface{}, 2*l.size)
		// 拷贝数组
		copy(newData, l.data)
		l.data = newData
	}
}

func (l *ArrayList) Set(index int, value interface{}) error {
	if index < 0 || index >= l.size {
		return errors.New("array out of index")
	}
	l.data[index] = value
	return nil
}

func (l *ArrayList) Clear() {
	l.size = 0
	l.data = make([]interface{}, 10)
}

func (l *ArrayList) Delete(index int) error {
	if index < 0 || index >= l.size {
		return errors.New("array out of index")
	}
	l.data = append(l.data[:index], l.data[index+1:]...)
	l.size--
	return nil
}

func (l *ArrayList) Iterator() Iterator {
	return &ArrayListIterator{
		list:      l,
		currIndex: 0,
	}
}

func (l *ArrayList) String() string {
	return fmt.Sprint(l.data)
}

// LinkedList 基于单链表的线性表
type LinkedList struct {
}

func (l LinkedList) Size() int {
	panic("implement me")
}

func (l LinkedList) Get(index int) (interface{}, error) {
	panic("implement me")
}

func (l LinkedList) Add(value interface{}) error {
	panic("implement me")
}

func (l LinkedList) Set(index int, value interface{}) error {
	panic("implement me")
}

func (l LinkedList) Clear() {
	panic("implement me")
}

func (l LinkedList) Delete(index int) error {
	panic("implement me")
}

func (l LinkedList) String() string {
	panic("implement me")
}

type DoubleLinkedList struct {
}

func (l DoubleLinkedList) Size() int {
	panic("implement me")
}

func (l DoubleLinkedList) Get(index int) (interface{}, error) {
	panic("implement me")
}

func (l DoubleLinkedList) Add(value interface{}) error {
	panic("implement me")
}

func (l DoubleLinkedList) Set(index int, value interface{}) error {
	panic("implement me")
}

func (l DoubleLinkedList) Clear() {
	panic("implement me")
}

func (l DoubleLinkedList) Delete(index int) error {
	panic("implement me")
}

func (l DoubleLinkedList) String() string {
	panic("implement me")
}
