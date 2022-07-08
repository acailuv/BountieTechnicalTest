package utils

import (
	"bytes"
	"encoding/json"
	"strconv"
)

func ConvertToByteArray(data interface{}) []byte {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(data)

	return buffer.Bytes()
}

func ConvertDecimalToBinary(num int) string {
	var binarySequence []int

	for num != 0 {
		binarySequence = append(binarySequence, num%2)
		num = num / 2
	}

	if len(binarySequence) == 0 {
		return "0"
	}

	result := ""

	for _, v := range binarySequence {
		result += strconv.Itoa(v)
	}

	return result
}
