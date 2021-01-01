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

