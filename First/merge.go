package main

import (
	"fmt"
)


func Merge(arr *[]string, l int, r int)   {

	i := l
	j := (r + l)/2 + 1
	mergeArr := make([]string,  r - l + 1)
	index := 0

	for {
		if i > (r + l ) / 2 || j > r {
			break
		} else {
			if (*arr)[i] > (*arr)[j] {
				mergeArr[index] = (*arr)[j]
				index++
				j++
			} else {
				mergeArr[index] = (*arr)[i]
				index++
				i++
			}
		}
	}

	for k := i; k <= (r + l ) / 2 ; k++ {
		mergeArr[index] = (*arr)[i]
		index++
		i++

	}


	for k := j; k < r + 1; k++ {
		mergeArr[index] = (*arr)[j]
		index++
		j++

	}

	for k := 0; k < len(mergeArr); k++ {
		(*arr)[k + l] = mergeArr[k]
	}


}

func MergeSort(arr *[]string, l int, r int )  {
	if ( r - l + 1 > 2) {
		MergeSort(arr, l, (r + l)/2 )
		MergeSort(arr, (r + l)/2 + 1, r)
	}
	Merge(arr, l, r)


}

func main()  {
	arr := []string{
		"bc", "b", "aa", "ab",  "a",
	}
	fmt.Println(arr)
	MergeSort(&arr, 0, len(arr) - 1)
	fmt.Println(arr)

}

