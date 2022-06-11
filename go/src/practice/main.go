// practice project main.go
package main

import (
	"fmt"
	//"sort"
)

/*
func main() {
		sort.Ints(a)

		fmt.Println(a)
	a := []int{0, 1, 0, 1, 2, 2, 0, 1}
	var temp = 0
	var mid int = 0
	var hi int = len(a) - 1
	var lo int = 0

	switch lo {
	case 1:
		temp = a[lo]
		a[lo] = a[mid]
		a[mid] = temp
		lo++
		mid++
		break
	case 2:
		mid++
		break
	case 3:
		temp = a[mid]
		a[mid] = a[hi]
		a[hi] = temp
		hi--
		break
	}
	fmt.Println(a, temp)

}*/

//try this for simple swamp

/*  temp = A[i]
    A[i] = A[j]
    A[j] = temp  //for swapping */

/*
  static void sort012(int a[], int arr_size)
    {
        int lo = 0;
        int hi = arr_size - 1;
        int mid = 0, temp = 0;
        while (mid <= hi) {
            switch (a[mid]) {
            case 0: {
                temp = a[lo];
                a[lo] = a[mid];
                a[mid] = temp;
                lo++;
                mid++;
                break;
            }
            case 1:
                mid++;
                break;
            case 2: {
                temp = a[mid];
                a[mid] = a[hi];
                a[hi] = temp;
                hi--;
                break;
            }
            }
        }
    }

    /* Utility function to print array arr[] */
/*   static void printArray(int arr[], int arr_size)
{
    int i;
    for (i = 0; i < arr_size; i++)
        System.out.print(arr[i] + " ");
    System.out.println("");
}

/*Driver function to check for above functions*/
/*   public static void main(String[] args)
    {
        int arr[] = { 0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1 };
        int arr_size = arr.length;
        sort012(arr, arr_size);
        System.out.println("Array after seggregation ");
        printArray(arr, arr_size);
    }
}

*/
func main() {
	var a, b, c = 3, 4, "foo"
	fmt.Printf("a is of type %T\n", a, b, c)
	//fmt.Println(bc)
	fmt.Println(i, j, k)
}

const (
	i = 7
	j
	k
)
