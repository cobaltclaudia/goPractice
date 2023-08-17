package main

import (
	"fmt"
	"math"

	"github.com/cobaltclaudia/goPractice/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.5))
	fmt.Println(math.Ceil(2.5))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("olleh"))
}
