package main

import (
	"fmt"

	"github.com/yeasin2002/better-next-app/internal/util"
)

func main() {
	fmt.Println("better-next-app")
    data :=util.ValidateNpmPackageName("he llo")
 	fmt.Println(data)

}
