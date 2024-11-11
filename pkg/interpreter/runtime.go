package interpreter

import "fmt"

type LuaTable map[any]any
type LuaFunction func(args []any)

func LuaPrintFunction(args []any) {
	fmt.Println(args...)
}
