package main

import (
	"fmt"

	"github.com/mark-vrb/hello-package-go"
)

func main() {
	fmt.Printf(hellopackage.Reverse("Hello, world!"))
}
