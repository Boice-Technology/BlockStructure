package consensus

import ("block_structure";
		"math/big";
		"fmt";
		"string_padding";
		)

func EvaluateFunction(b *block_structure.Block) string {
	targetBits := b.Header.TargetDifficulty
	exponentHex := targetBits[2:4] 
	coefficientHex := targetBits[4:] 
	eight,_ := new(big.Int).SetString("8",10)
	three,_ := new(big.Int).SetString("3",10)
	exp_big,_ := new(big.Int).SetString(exponentHex,16)
	subtract := new(big.Int).Sub(exp_big,three)
	power := new(big.Int).Mul(subtract,eight) 
	coefficientInt,_ := new(big.Int).SetString(coefficientHex,16)
	right_exp,_ := new(big.Int).SetString("1",10)
	x,_ := new(big.Int).SetString("2",10)
	right_exp.Exp(x,power,nil)
	targetInt := new(big.Int).Mul(coefficientInt,right_exp)
	target := fmt.Sprintf("%x",targetInt)
	target = string_padding.LeftPadHexadecimalString(target,32)
	return target
}

//dependent