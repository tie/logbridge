package logbridge_test

import (
	"log"
	"os"

	"github.com/tie/logbridge"
)

func ExampleWrap() {
	l1 := log.New(os.Stdout, "outer: ", log.Lshortfile|log.Lmsgprefix)
	l2 := logbridge.Wrap(l1, "inner: ", 0)

	l1.Println("hello")
	l2.Println("world")

	// Output:
	// example_test.go:14: outer: hello
	// example_test.go:15: outer: inner: world
}
