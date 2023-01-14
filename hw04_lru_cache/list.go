package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type list struct {
	first, last *ListItem
	len         int
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	li := &ListItem{Value: v, Prev: nil, Next: l.first}
	if l.first != nil {
		l.first.Prev = li
	}
	l.first = li
	if l.last == nil {
		l.last = li
	}
	l.len++
	return li
}

func (l *list) PushBack(v interface{}) *ListItem {
	li := &ListItem{Value: v, Prev: l.last, Next: nil}
	if l.last != nil {
		l.last.Next = li
	}
	l.last = li
	if l.first == nil {
		l.first = li
	}
	l.len++
	return li
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.first = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.last = i.Prev
	}
	i.Prev = nil
	i.Next = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i != l.first {
		l.Remove(i)
		l.PushFront(i.Value)
	}
}

func NewList() List {
	return new(list)
}
