package list

import (
	"testing"
)

type MyData struct {
	data int
}

func (me *MyData) Equals(other Any) bool {
	return me.data == other.(*MyData).data
}

func checkListLen(t *testing.T, l *LinkedList, n int) {
	if an := l.Len(); an != n {
		t.Errorf("l.Len() = %d, want %d", an, n)
	}
}

func (me MyData) toString() string{
	return string (me.data)
}

func TestList(t *testing.T) {
	l := New()
	checkListPointers(t, l, []*Node{})
	checkListLen(t, l, 0)
	d1 := &(MyData{1})
	n1 := l.Append(d1)
	checkListPointers(t, l, []*Node{n1})
	checkListLen(t, l, 1)
	l.RemoveElement(d1)
	checkListPointers(t, l, []*Node{})
	checkListLen(t, l, 0)
	n1 = l.Append(d1)
	d2 := &(MyData{2})
	n2 := l.Append(d2)
	d3 := &(MyData{3})
	n3 := l.Append(d3)
	d4 := &(MyData{4})
	n4,_ := l.Insert(d4, 0)
	d5 := &(MyData{5})
	n5 := l.Append(d5)
	checkListPointers(t, l, []*Node{n4, n1, n2, n3, n5})
	checkListLen(t, l, 5)
}

func checkListPointers(t *testing.T, l *LinkedList, es []*Node) {
	if len(es) == 0 {
		if l.head.prev != l.head || l.head.next != l.head {
			t.Errorf("%v/%v should be pointing to %v", l.head.prev, l.head.next, l.head)
		}
		return
	}

	if l.head.next != es[0] {
		t.Errorf("l.head.next = %v, want %v", l.head.next, es[0])
	}

	if last := es[len(es)-1]; l.head.prev != last {
		t.Errorf("l.head.prev = %v, want %v", l.head.prev, last)
	}

	for i, e := range es {
		var e_prev = es[i]
		var e_next = e_prev.next

		if i == 0 {
			e_prev = l.head
		} else{
			e_prev = es[i-1]
		}
		
		if e.prev != e_prev {
			t.Errorf("node #%d (%v) has: \nprev=%v \nwant %v\n\n", i, e, e.prev, e_prev)
		}
		
		if i < len(es) - 1{
			e_next = es[i+1]
		} 
		
		if e.next != e_next {
			t.Errorf("node #%d (%v) has: \nnext=%v \nwant %v\n\n", i, e, e.next, e_next)
		}
	}
}
