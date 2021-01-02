package main

import (
	"fmt"
	"time"
	"math/rand"
	"testing"
)

func BubbleSort(arr *[]int)  {

	for i := 0; i<len(*arr) - 1; i++ {
		for j := 0; j < len(*arr) - i - 1; j++ {
			if  (*arr)[j] > (*arr)[j + 1] {
				(*arr)[j] += (*arr)[j+1]
				(*arr)[j+1] = (*arr)[j] - (*arr)[j+1]
				(*arr)[j] = (*arr)[j] - (*arr)[j+1]
			}
		}
	}
}



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


func TestBubblesort(t *testing.T)  {
	arr1 := []int {
		5, 1, 2, 5, 0, 4, 5, 6,
	}

	start := time.Now()
	BubbleSort(&arr1)
	end := time.Since(start)
	fmt.Println("Spended time for first array:  ", end.Microseconds(), "microsecond")

	arr2 := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr2[i] = rand.Int()
	}

	start = time.Now()
	BubbleSort(&arr2)
	end = time.Since(start)
	fmt.Println("Spended time for second array: ", end.Milliseconds(), "millisecond")

	for i := 0; i < 7; i++ {
		if arr1[i] > arr1[i+1] {
			t.Errorf("First array didn't sort")
		}
	}

	for i := 0; i<999; i++ {
		if arr2[i] > arr2[i+1] {
			t.Errorf("Second array didn't sort")
		}
	}

}



func TestMerge(t *testing.T)  {
	arr := []string{
		"bc", "b", "aa", "ab",  "a",
	}
	fmt.Println(arr)
	MergeSort(&arr, 0, len(arr) - 1)
	for i := 0; i < len(arr) - 1; i++ {
		if (arr[i] > arr[i+1])  {
			t.Errorf("Failed sorting")
		}
	}
	fmt.Println(arr)

}
