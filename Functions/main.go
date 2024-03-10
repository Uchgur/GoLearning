package main

func main() {
	//Example of function syntax
	Syntax()

	//Example of using variadic parameter (random number of input argumetns of the same type)
	Variadic()

	//Example of unfurling slice (means getting data from a slice. For example in this function we getting all numbers from a slice and passing them like individual numbers)
	Unfurling()

	//Example of using defer statement (it's invokes function only when calling function execute return, reaching it's end or panic)
	Defer()

	//Example of creating methods in go (functions that attached to the type)
	Method()

	//Example of using unterfaces in go
	InterAndPol()

	//Example of using stringer interface. Stringer interface change the way of printing
	//Also creating logs with this interface
	StringerInter()

	//Example of writing to the file
	WritingToFile()

	//Example of writing to byte buffer
	ByteBuffer()

	//Example of writing to either a file or a buffer
	FileAndBuffer()

	//Example of using anonymous functions
	AnonFunc()

	//Example of using func expressions (like delegates in c#)
	FuncExpression()

	//Example of returning function
	ReturningFunc()

	//Example of callback (means passing function as an argument)
	Callback()

	//Example of how func variable save the state of assigned function
	Closure()

	//Example of using recursion and comparinf it with non-recursive function
	Recursion()

	//Example of using wrapper function to read from file
	Wrapper()
}
