package main

import (
	"fmt"
)

func main() {

	var originalCount int = 10
	fmt.Println("I started with", originalCount)
	var eatenCount = originalCount - 6
	fmt.Println("Some jerk ate", eatenCount, "apples")
	eatenCount = eatenCount + 2

	fmt.Println("There are ", eatenCount, "apples left")

}
