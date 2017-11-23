package test_go1

import (
	"fmt"
	a "github.com/a8uhnf/test-go2"
)

func HelloWorld() {
	fmt.Println("Hello World from test-go1. branch: test-dep1")
	a.HelloWorldMaster()
}

func HelloWorldNewWorld()  {
	fmt.Println("Hello World from New World!!!!! Branch: test-dep1")
	a.HelloWorldMaster()
}
