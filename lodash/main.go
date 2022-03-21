package main

import (
	"fmt"
	"github.com/samber/lo"
)

func main() {
	name := lo.Uniq([]string{"Samuel", "Marc", "Samuel"})
	fmt.Println(name)
}
