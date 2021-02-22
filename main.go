package main

import (
	"fmt"
	"strings"
)

func AboveBelow(nums []int, target int) {
	above := 0
	below := 0
	for _, value := range nums {
		if  value > target {
			above++
		}else{
			below++
		}
	}
	fmt.Printf("Above: %d\nBelow: %d\n",above, below)
}

func RotateString(s string, rotation int) string{
	if len(s) < 2 {
		return s
	}

	negative := false
	if rotation < 0 {
		negative = true
		rotation *=-1
	}

	for rotation > len(s) {//skip warps
		rotation -= len(s)
	}

	builder := strings.Builder{}
	if negative {
		builder.WriteString(s[rotation:])
		builder.WriteString(s[:rotation])
	}else{
		builder.WriteString(s[len(s)-rotation:])
		builder.WriteString(s[:len(s)-rotation])
	}
	return builder.String()


}


func main () {
	array := []int{1, 5, 2, 1, 10}
	target := 6
	fmt.Printf("Above below with array = %v, target = %d\n", array, target)
	AboveBelow(array, target)
	target = 3
	fmt.Printf("Above below with array = %v, target = %d\n", array, target)
	AboveBelow(array, target)
	
	s := "MyString"
	fmt.Println("Rotating string:",s)
	rotation := []int {2,10,-2,-11, 8}
	for _, r := range rotation {
		fmt.Println("Rotation:", r)
		fmt.Println(RotateString(s,r))

	}

}