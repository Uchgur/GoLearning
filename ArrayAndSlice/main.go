package main

func main() {
	//Example of using arrays
	Array()

	Ex40()

	//Example of using slices
	Slice()

	Ex41()

	//Example of using for range loops and blank identifiers for a slicce
	RangeBlank()

	//Example of Appending slices
	Appending()

	//Example of slicing slices(means taking only some part of it)
	Slicing()

	//Example of deleting value from a slice
	Deleting()

	//Example of using make function to create a slice
	MakeFunc()

	//Example of using multidimensional slices
	Multidimensional()

	//In this example we see that the variable is only a pointer to the part of a memory where our data stored
	SlicePointer()

	//In this example we see that b didn't change because copy function also copies data to new part of a memory
	SlicePointer2()

	//Comparing copy by operator and by function in one file
	SlicePointer3()
}
