// liga sepak bola

package main

import "fmt"

func main(){
	var input1, input2, input3, input4, hasil1, hasil2 int

	fmt.Scan(&input1, &input2, &input3, &input4)
	x := input3 == input1
	x1 := input3 == input2
	x2 := input3 == input4
	if ((input1 < input2) && (input1 < input3) && (input1 < input4)) && ((input2 > input1) && (input2 > input3) && (input2 > input4)) {
		hasil1 = input2
		hasil2 = input1
		fmt.Print(hasil1, hasil2)
	} else if ((input1 < input2) && (input1 < input3) && (input1 < input4)) && ((input3 > input1) && (input3 > input2) && (input3 > input4)) {
		hasil1 = input3
		hasil2 = input1
		fmt.Print(hasil1, hasil2)
	} else if ((input1 < input2) && (input1 < input3) && (input1 < input4)) && ((input4 > input1) && (input4 > input2) && (input4 > input3)) {
		hasil1 = input4
		hasil2 = input1
		fmt.Print(hasil1, hasil2)
	} else if ((input2 < input1) && (input2 < input3) && (input2 < input4)) && ((input1 > input2) && (input1 > input3) && (input1 > input4)) {
		hasil1 = input1
		hasil2 = input2
		fmt.Print(hasil1, hasil2)
	} else if ((input2 < input1) && (input2 < input3) && (input2 < input4)) && ((input3 > input2) && (input3 > input1) && (input3 > input4)) {
		hasil1 = input3
		hasil2 = input2
		fmt.Print(hasil1, hasil2)
	} else if ((input2 < input1) && (input2 < input3) && (input2 < input4)) && ((input4 > input2) && (input4 > input1) && (input4 > input3)) {
		hasil1 = input4
		hasil2 = input2
		fmt.Print(hasil1, hasil2)
	} else if ((input3 < input1) && (input3 < input2) && (input3 < input4)) && ((input1 > input2) && (input1 > input3) && (input1 > input4)) {
		hasil1 = input1
		hasil2 = input3
		fmt.Print(hasil1, hasil2)
	} else if ((input3 < input1) && (input3 < input2) && (input3 < input4)) && ((input2 > input1) && (input2 > input3) && (input2 > input4)) {
		hasil1 = input2
		hasil2 = input3
		fmt.Print(hasil1, hasil2)
	} else if ((input3 < input1) && (input3 < input2) && (input3 < input4)) && ((input4 > input2) && (input4 > input1) && (input4 > input3)) {
		hasil1 = input4
		hasil2 = input3
		fmt.Print(hasil1, hasil2)
	} else if ((input4 < input1) && (input4 < input3) && (input4 < input2)) && ((input3 > input2) && (input3 > input1) && (input3 > input4)) {
		hasil1 = input3
		hasil2 = input4
		fmt.Print(hasil1, hasil2)
	} else if ((input4 < input1) && (input4 < input3) && (input4 < input2)) && ((input2 > input1) && (input2 > input3) && (input2 > input4)) {
		hasil1 = input2
		hasil2 = input4
		fmt.Print(hasil1, hasil2)
	} else if ((input4 < input1) && (input4 < input3) && (input4 < input2)) && ((input1 > input2) && (input1 > input3) && (input1 > input4)) {
		hasil1 = input1
		hasil2 = input4
		fmt.Print(hasil1, hasil2)
	} else if ((input1 == input2) && (input1 == input3) && (input1 == input4)) && ((input2 == input1) && (input2 == input3) && (input2 == input4)) && ((x) && (x1) && (x2)) && ((input3 == input1) && (input3 == input2) && (input3 == input4)) && ((input4 == input1) && (input4 == input3) && (input4 == input2))  {
		hasil1 = input1
		hasil2 = input1
		fmt.Print(hasil1, hasil2)
	}
}