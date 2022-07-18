package ds

type Stack struct {
	items List
}

// Front returns the first element of list l or nil if the list is empty.
func (l *Stack) Front() *Element {
	if l.items.len == 0 {
		return nil
	}
	return l.items.root.next
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *Stack) Remove(e *Element) any {
	if e.list == &l.items {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.items.remove(e)
	}
	return e.Value
}

// PushFront(e) inserts a new element e with value v at the front of list l and returns e.
func (l *Stack) PushFront(v any) *Element {
	l.items.lazyInit()
	return l.items.insertValue(v, &l.items.root)
}
