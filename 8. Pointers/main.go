package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	reverse(&a)
	fmt.Println(a)
}

func reverse(num *[4]int) {

	for i, v := range *num {
		(*num)[len(num)-1-i] = v
	}

}
