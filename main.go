package main

import (
	"fmt"
	"web-osLearn/router"
)

func main() {
	fmt.Println(1)
	r := router.InitRouter()
	fmt.Println(2)
	err := r.Run(fmt.Sprintf(":%d", 9264))
	fmt.Println(3)
	if err != nil {
		return
	}
}
