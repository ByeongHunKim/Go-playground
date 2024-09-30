// https://go-tour-ko.appspot.com/basics/3
package main

import (
	"fmt"
	"math"
)

func main() {
	// math.pi로 하면 ./exported-names.go:9:19: undefined: math.pi 에러 발생
	fmt.Println(math.Pi)
}
