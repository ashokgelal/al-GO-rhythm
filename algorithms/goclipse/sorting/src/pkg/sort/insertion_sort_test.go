package sort

import (
	"testing"
	"fmt"
)

func TestSortDesc(t *testing.T){
	t1 := &(MyType{3})
	t2 := &(MyType{2})
	t3 := &(MyType{4})
	t4 := &(MyType{1})
	t5 := &(MyType{8})
	t6 := &(MyType{5})
	t7 := &(MyType{6})
	v := []MyType{*t3, *t2, *t4, *t1, *t5, *t6, *t7}
	
	fmt.Printf("Before sorting: %v\n", v)
	Sort(&v, compare_desc)
	fmt.Printf("After sorting in desc. order: %v\n", v)
	
	for i:= 0; i<len(v)-1; i++{
		x := v[i].data.(int)
		if y := v[i+1].data.(int); x < y{
			t.Errorf("%d > %d", x, y)
			return
		}
	}
}

func TestSortAsc(t *testing.T){
	t1 := &(MyType{3})
	t2 := &(MyType{2})
	t3 := &(MyType{4})
	t4 := &(MyType{1})
	t5 := &(MyType{8})
	t6 := &(MyType{5})
	t7 := &(MyType{6})
	v := []MyType{*t3, *t2, *t4, *t1, *t5, *t6, *t7}
	
	fmt.Printf("Before sorting: %v\n", v)
	Sort(&v, compare_asc)
	fmt.Printf("After sorting in asc. order: %v\n", v)
	
	for i:= 0; i<len(v)-1; i++{
		x := v[i].data.(int)
		if y := v[i+1].data.(int); x > y{
			t.Errorf("%d > %d", x, y)
			return
		}
	}
}

func compare_asc(a MyType, b MyType) bool {
	return a.data.(int) < b.data.(int)
}

func compare_desc(a MyType, b MyType) bool {
	return a.data.(int) > b.data.(int)
}

