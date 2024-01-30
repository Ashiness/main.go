package Digit

import (
	"awesomeProject3/wordz"
)

var Hello = "This is package Digit"
var Prefix = "Random Digit: "
var Digits = []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
var Result = ""

func Digit() string {

	wordz.Hello = Hello
	wordz.Prefix = Prefix
	wordz.Words = Digits

	Result = wordz.Random()

	return Result
}
