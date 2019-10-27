package consensus

import ("math/big";
		"fmt";
		)

func RetargetDifficulty(currentTargetDifficulty string, timeInterval int64) (string) {
	exponentHex := currentTargetDifficulty[2:4] 
	coefficientHex := currentTargetDifficulty[4:] 

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

	timeIntervalString := fmt.Sprint(timeInterval)
	timeIntervalInt,_ := new(big.Int).SetString(timeIntervalString,10)

	var idealTimeInterval int64 = 2016 * 10 * 60
	idealTimeIntervalString := fmt.Sprint(idealTimeInterval)
	idealTimeIntervalInt,_ := new(big.Int).SetString(idealTimeIntervalString,10)

	mulRes,_ := new(big.Int).SetString("0",10)
	mulRes.Mul(timeIntervalInt, targetInt)

	newTargetInt,_ := new(big.Int).SetString("0",10)
	newTargetInt.Div(mulRes, idealTimeIntervalInt)
	newTarget := fmt.Sprint(newTargetInt)

	newTargetBitForm := EvaluateTargetBit(newTarget)

	return newTargetBitForm
}