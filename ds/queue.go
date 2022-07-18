package ds

type Queue struct {
	items List
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *Queue) PushBack(v any) *Element {
	l.items.lazyInit()
	return l.items.insertValue(v, l.items.root.prev)
}

// Front returns the first element of list l or nil if the list is empty.
func (l *Queue) FrontQueue() *Element {
	if l.items.len == 0 {
		return nil
	}
	return l.items.root.next
}
