package Cities

import (
	"awesomeProject3/wordz"
)

var Hello = "This is package Cities"
var Prefix = "Random City: "
var Cities = []string{"New-York", "Los Angeles", "Washington DC", "Berlin", "Moscow", "Paris", "Minks"}
var Result = ""

func City() string {

	wordz.Words = Cities
	wordz.Hello = Hello
	wordz.Prefix = Prefix

	Result = wordz.Random()

	return Result
}
