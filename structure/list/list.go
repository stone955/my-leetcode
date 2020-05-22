package list

// List 线性表接口
type List interface {
	Size() int
	Get(index int) (interface{}, error)
	Add(value interface{}) error
	Set(index int, value interface{}) error
}

// ArrayList 基于数组的线性表
type ArrayList struct {
	data []interface{} // 存储数据的数组
	size int           // 数组的大小
}

func (l ArrayList) Size() int {
	return l.size
}

func (l ArrayList) Get(index int) (interface{}, error) {
	panic("implement me")
}

func (l ArrayList) Add(value interface{}) error {
	panic("implement me")
}

func (l ArrayList) Set(index int, value interface{}) error {
	panic("implement me")
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
