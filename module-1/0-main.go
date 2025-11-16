/**
* go run main.go
**/

package main;

import (
	"fmt"
	"runtime"
);

func main() {
	fmt.Println("Hello, Go world!");
	fmt.Printf("Running on %s OS with %s architecture\n", runtime.GOOS, runtime.GOARCH);
}