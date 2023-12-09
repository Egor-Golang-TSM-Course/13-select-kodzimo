package main

import "time"

func main() {
	var c1 Chief = Chief{Name: "Greg", CookingTime: 3}
	var c2 Chief = Chief{Name: "John", CookingTime: 6}
	var c3 Chief = Chief{Name: "Mike", CookingTime: 9}

	go c1.cooking()
	go c2.cooking()
	go c3.cooking()

	time.Sleep(10 * time.Second)
}
