package main;

import "fmt";

//At the package level scope, assign the following values to the three variables
//a. for x assign 42
//b. for y assign “James Bond”
//c. for z assign true

var x = 42;
var y = "James Bond";
var z = true;

//In func main ::: use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
//returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
//print out the value stored by variable “s”

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z);
	fmt.Println(s)
}
