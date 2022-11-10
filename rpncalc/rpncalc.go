package rpncalc

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"math"
	"regexp"
	"strconv"
)

var processorBinary = map[string]func(a float64, b float64) float64{
	"+": func(a float64, b float64) float64 { return a + b },
	"-": func(a float64, b float64) float64 { return a - b },
	"*": func(a float64, b float64) float64 { return a * b },
	"/": func(a float64, b float64) float64 { return a / b },
	"^": func(a float64, b float64) float64 { return math.Pow(a, b) },
}

var processorUnary = map[string]func(a float64) float64{
	"!": func(a float64) float64 { return a },
}

func Rpn(in string, st *stack.Stack) string {
	re := regexp.MustCompile(`[0-9]+|[\+\-\*\/\^]`)
	args := re.FindAllString(in, -1)
	fmt.Printf("%v", args)
	for _, arg := range args {
		if fnBinary, ok := processorBinary[arg]; ok {
			a := st.Pop()
			b := st.Pop()
			if a == nil || b == nil {
				return "ERR"
			}
			st.Push(fnBinary(b.(float64), a.(float64)))
		} else if fnUnary, ok := processorUnary[arg]; ok {
			a := st.Pop()
			if a == nil {
				return "ERR"
			}
			st.Push(fnUnary(a.(float64)))
		} else {
			if i, err := strconv.ParseFloat(arg, 64); err == nil {
				st.Push(i)
			} else {
				return "ERR"
			}
		}
	}
	res := st.Peek()
	if res == nil {
		return "ERR"
	}
	return fmt.Sprintf("%f", res.(float64))
}
