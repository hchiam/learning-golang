package main

import (
	"fmt"
	t "time"
)

func main() {
	// single-line comment
	/* multi-line
	comment
	*/
	fmt.Println("Hello World")
	fmt.Println(t.Now())
	fmt.Println(`
--multiline
----string
------output
  ______________
 /   0      0   \
|    \______/    |
 \______________/
`)

	const earthsGravity = 9.80665
	fmt.Println("const earthsGravity =", earthsGravity)

	var anInteger int = 1 // int8, int16, int32, int64, uint8, unit16, uint32, unit64
	/*
		var anUnsignedInteger uint16 = 1
		var anotherInteger int8
		var aFloat float32 = 3.14 // there's also complex
		var aBoolean bool = true
	*/
	fmt.Println("anInteger =", anInteger)
}
