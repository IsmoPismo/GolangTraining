// composite literal; slice literal
package main

import "fmt"

func main() {
	x := []int{8, 55, -23, -1}
	fmt.Println(x)

	y := make([]int, 0, 3)
	y = append(y, 8)
	y = append(y, -44)
	y = append(y, 5)
	fmt.Println(y)
	fmt.Println(len(y))
}
