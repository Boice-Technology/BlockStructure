package consensus

import ("math/big";
		"fmt";
		)

func EvaluateTargetBit(targetString string) string {
	targetInt,_ := new(big.Float).SetString(targetString)
	mantissa,_ := new(big.Float).SetPrec(1024).SetString("0")
	exponent := targetInt.MantExp(mantissa)
	remainder := exponent % 8
	exponent -= remainder

	remainderString := fmt.Sprint(remainder)
	remainderInt,_ := new(big.Int).SetString(remainderString, 10)

	two,_ := new(big.Int).SetString("2", 10)
	twoFiftySix,_ := new(big.Float).SetPrec(1024).SetString("256")
	valueOf3Bytes,_ := new(big.Float).SetPrec(1024).SetString("16777215")

	multiplicant,_ := new(big.Int).SetString("0", 10)

	multiplicant.Exp(two, remainderInt, nil)
	multiplicantString := fmt.Sprint(multiplicant)
	multiplicantFloat,_ := new(big.Float).SetPrec(1024).SetString(multiplicantString)

	mantissa.Mul(mantissa, multiplicantFloat)
	mantissaString := fmt.Sprint(mantissa)

	interFloat,_ := new(big.Float).SetPrec(1024).SetString(mantissaString)
	interMantissaString := mantissaString
	interFloat.Mul(interFloat, twoFiftySix)
	
	for interFloat.Cmp(valueOf3Bytes) == -1 {
		interMantissaString = fmt.Sprint(interFloat)
		exponent -= 8
		interFloat.Mul(interFloat, twoFiftySix)
	}

	exponent /= 8
	exponent += 3

	mantissa,_ = new(big.Float).SetString(interMantissaString)
	mantissaInt,_ := new(big.Int).SetString("0", 10)
	mantissa.Int(mantissaInt)

	mantissaString = fmt.Sprintf("%x", mantissaInt)
	exponentString := fmt.Sprintf("%x", exponent)

	targetBitForm := "0x" + exponentString + mantissaString

	return targetBitForm
}