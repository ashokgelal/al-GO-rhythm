package sort

func Sort(arr *[]MyType, compare func(a MyType, b MyType) bool){
	for i:= 1; i<len(*arr); i++ {
		j := i
		
		for (j>0) && compare((*arr)[j], (*arr)[j-1]){
			(*arr)[j], (*arr)[j-1] = (*arr)[j-1], (*arr)[j]
			j--
		}
	}
}


type Any interface{}

type MyType struct{
	data interface{}
}