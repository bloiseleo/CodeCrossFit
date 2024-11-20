package arrays

import "fmt"

type DynamicArray struct {
	arr    []int
	length int
}

/*
Creates a Dynamic Array that grows as the user inserts elements inside of it
O(1)
*/
func CreateDyanmicArray() *DynamicArray {
	return &DynamicArray{
		arr:    make([]int, 10),
		length: 0,
	}
}

/*
Grows the underlying array creating another array with twice the capacity of the original array.
*/
func (da *DynamicArray) growArray() {
	newArr := make([]int, da.length*2)
	copy(newArr, da.arr[:])
	da.arr = newArr
}

/*
Appends an element in the end of the DynamicArray. If the underlying array is full, so a new one is created and the element is inserted at the logical end of the array. O(n) when we need to grow the array capacity`s. O(1) when we do not need to grow the array.
*/
func (da *DynamicArray) Append(element int) {
	if len(da.arr) <= da.length {
		da.growArray()
	}
	da.arr[da.length] = element
	da.length++
}

/*
Inserts an element [element] at the index [position] inside the DyanmicArray. O(n)
*/
func (da *DynamicArray) Inserts(element, position int) {
	if position < 0 || position >= da.length {
		panic(fmt.Sprintf("Index [%d] out of bounds [%d]", position, da.length))
	}
	if position == da.length {
		da.Append(element)
		return
	}
	if len(da.arr) <= da.length {
		da.growArray()
	}
	for i := da.length - 1; i > position; i-- {
		da.arr[i+1] = da.arr[i]
	}
	da.arr[position+1] = da.arr[position]
	da.arr[position] = element
	da.length++
}

/*
	Returns the logical length of the array. O(1)
*/
func (da *DynamicArray) Length() int {
	return da.length
}

/*
Removes an element from the array based on the position informed. O(n) when it's removing from the middle of the array, but it's O(1) if it's removing from the tail.
*/
func (da *DynamicArray) Delete(position int) {
	if position < 0 || position >= da.Length() {
		panic(fmt.Sprintf("Index [%d] out of bounds [%d]", position, da.length))
	}
	if position == da.Length()-1 {
		da.length--
		return
	}
	for i := position; i < da.Length()-1; i++ {
		da.arr[i] = da.arr[i+1]
	}
	da.length--
}

/*
Creates a static copy of the DyanmicArray.
*/
func (da *DynamicArray) ToStaticArray() []int {
	return da.arr[0:da.length]
}
