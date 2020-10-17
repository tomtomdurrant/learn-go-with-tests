package main

import (
	"os"
	"io"
	"fmt"
)

const finalWord = "Go!"
const countdownStart = 3
func main()  {
	Countdown(os.Stdout)

}

func Countdown(out io.Writer)  {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}
