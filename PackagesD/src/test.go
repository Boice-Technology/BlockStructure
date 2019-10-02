package main

import ("fmt";
		"string_padding";
		)

func main(){
	var s string = "ABCABC12"
	res := string_padding.LeftPadHexadecimalString(s,8)
	fmt.Println(res)
}