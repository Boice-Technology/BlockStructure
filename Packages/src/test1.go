package main

import ("consensus";
		"fmt";
		"math/big";
		)

func main(){
	fmt.Println(consensus.EvaluateTargetBit("565365551653621256325165326156253625327455436256"))
	two,_ := new(big.Int).SetString("2",10)
	four,_ := new(big.Int).SetString("4",10)
	res,_ := new(big.Int).SetString("0",10)
	res.Div(four, two)
	fmt.Println(res)
	res.Div(two,four)
	fmt.Println(res)
}