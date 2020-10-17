package main

import (
	"io"
	"os"
	// "bytes"
	"fmt"
)

func Greet(writer io.Writer, name string)  {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main()  {
	Greet(os.Stdout, "Jim")
}