package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now().Weekday()
	year := time.Now().Year()
	fmt.Println(now,"\n",year)
	
	
	

}
