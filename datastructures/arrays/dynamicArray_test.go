package arrays

import "testing"

func TestAppendElement(t *testing.T) {
	arr := CreateDyanmicArray()
	arr.Append(1)
	if arr.length != 1 {
		t.Fatalf("Array must have 1 length, %d returned", arr.length)
	}
	if arr.arr[arr.length-1] != 1 {
		t.Fatalf("Array must have element \"1\" at the end of array, %d returned", arr.arr[arr.length-1])
	}
}

func TestLength(t *testing.T) {
	arr := CreateDyanmicArray()
	if arr.Length() != arr.length {
		t.Fatalf("Recenlty created Array must have 0 of length, since no elements are inside the array")
	}
	arr.Append(1)
	if arr.Length() != 1 || arr.length != 1 {
		t.Fatalf("Array must have 1 of length, %d returned", arr.Length())
	}
}

func TestInsert(t *testing.T) {
	arr := CreateDyanmicArray()
	for i := 0; i < 10; i++ {
		arr.Append(i)
	}
	// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
	arr.Inserts(2, 3)
	// 0, 1, 2, 2, 3, 4, 5, 6, 7, 8, 9
	if arr.arr[3] != 2 {
		t.Fatalf("Element at position 3 must be 2, %d returned", arr.arr[3])
	}
	if arr.arr[4] != 3 {
		t.Fatalf("Element at position 4 must be 3, %d returned", arr.arr[4])
	}
}

func TestDelete(t *testing.T) {
	arr := CreateDyanmicArray()
	for i := 0; i < 10; i++ {
		arr.Append(i)
	}
	arr.Delete(7)
	if arr.Length() != 9 {
		t.Fatalf("Array must have 9 elements, %d returned", arr.Length())
	}
	for i := 0; i < 9; i++ {
		e := arr.arr[i]
		if e == 7 {
			t.Fatalf("Array must not have element 7, since it was removed from the array")
		}
	}
}
