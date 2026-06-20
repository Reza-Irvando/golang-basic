package main

import "fmt"

func main() {
	// basic array
	var items [5]string = [5]string{"Apple", "Banana", "Cherry", "Mango", "Elderberry"}
	fmt.Println("Items in the array: ", items)
	fmt.Println("Length of the array: ", len(items))
	fmt.Println("Third item in the array: ", items[2])
	for i := 0; i < len(items); i++ {
		fmt.Printf("Item %d: %s\n", i, items[i])
	}
	items[1] = "Blueberry"
	fmt.Println("Updated array: ", items)
	for index, item := range items {
		fmt.Printf("Item %d: %s\n", index, item)
	}

	// array comparison
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}
	arr3 := [3]int{1, 2, 3}

	fmt.Println("arr1 == arr2:", arr1 == arr2)
	fmt.Println("arr1 == arr3:", arr1 == arr3)
	fmt.Println("arr1 != arr2:", arr1 != arr2)
	fmt.Println("arr1 != arr3:", arr1 != arr3)

	// array slice
	full:= [5]int{1, 2, 3, 4, 5}
	fmt.Println("Full array: ", full)
	s1:= full[:]
	fmt.Println("Slice [:]: ", s1)
	s2:= full[:3]
	fmt.Println("Slice [:3]: ", s2)
	s3:= full[2:]
	fmt.Println("Slice [2:]: ", s3)
	s4:= full[1:4]
	fmt.Println("Slice [1:4]: ", s4)

	// slice functions
	slice1 := make([]int, 3, 5)
	fmt.Println("Original slice: ", slice1)
	fmt.Println("Length of the slice: ", len(slice1))
	fmt.Println("Capacity of the slice: ", cap(slice1))
	slice1 = append(slice1, 60, 70)
	slice1[0] = 10
	fmt.Println("Appended slice1: ", slice1)
	slice2 := make([]int, 4)
	fmt.Println("Created slice2: ", slice2)
	fmt.Println("Length of the created slice2: ", len(slice2))
	fmt.Println("Capacity of the created slice2: ", cap(slice2))
	slice3 := copy(slice2, slice1)
	fmt.Println("Copied slice2: ", slice2)
	fmt.Println("Number of elements copied: ", slice3)

	// slice 4 = slice 5 then change slice 5
	slice4 := []int{1, 2, 3, 4, 5}
	slice5 := slice4[2:]
	slice5[2] = 100
	fmt.Println("slice4 after modifying slice5: ", slice4)
	fmt.Println("slice5 after modification: ", slice5)
}