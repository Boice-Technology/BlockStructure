// Left pad a hexadecimal string for specified number of bytes in result string.

package string_padding

func LeftPadHexadecimalString(hexString string, totalBytes int64) string {
	var length int = len(hexString)
	var totalHexChars int = int(totalBytes) * 2
	diff := totalHexChars - length
	var resultString string = ""
	for i := 0 ; i < diff ; i++ {
		resultString += "0"
	}
	resultString += hexString
	return resultString
}