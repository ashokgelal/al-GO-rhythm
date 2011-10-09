package list

import (
	"testing"
)

type MyData struct{
	data int
}

func (me *MyData) Equals(other Any) bool{
	return me.data == other.(MyData).data
}


func checkListLen(t *testing.T, l *LinkedList, n int){
	if an := l.Len(); an != n{
		t.Errorf("l.Len() = %d, want %d", an, n)
	}
}

func TestList(t *testing.T){
	l := New()
	checkListLen(t, l, 0)
	
	l.Append(&(MyData{1}))
	checkListLen(t, l, 1)
}

