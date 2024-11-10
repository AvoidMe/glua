package interpreter

import "fmt"

type Function func(args []any)

func PrintFunction(args []any) {
	fmt.Println(args...)
}
