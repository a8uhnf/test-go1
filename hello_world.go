package test_go1

import (
	"fmt"
	a "github.com/a8uhnf/test-go2"
)

func HelloWorld() {
	fmt.Println("Hello World from test-go1")
	a.HelloWorldTestDep()
}
