package utffileconv

import (
	"encoding/base64"
	"fmt"
	"os"
	"unicode/utf16"
)

func CreateFileUTF16(name string) error {
	var bytes [2]byte
	const BOM = '\ufffe' //LE. for BE '\ufeff'
	data := `test string UTF-8`

	file, err := os.Create(name)
	if err != nil {
		fmt.Errorf("Can't open file. %v", err)
		return err
	}
	defer file.Close()

	bytes[0] = BOM >> 8
	bytes[1] = BOM & 255

	file.Write(bytes[0:])
	runes := utf16.Encode([]rune(data))
	for _, r := range runes {
		bytes[1] = byte(r >> 8)
		bytes[0] = byte(r & 255)
		file.Write(bytes[0:])
	}
	return nil
}

func CreateFile(nameFile string, message string) error {
	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Errorf("Can't open file. %v", err)
		return err
	}
	defer file.Close()
	file.Write([]byte(message))
	return nil
}

// Convert UTF8 characters string to UTF16LE Bytes array
func ConvertUTF8toUTF16LE(message string) []byte {
	utf16Runs := utf16.Encode([]rune(message))
	bytesToConv := make([]byte, len(utf16Runs)*2)

	for i, r := range utf16Runs {
		bytesToConv[i*2] = byte(r & 255)
		bytesToConv[i*2+1] = byte(r >> 8)
	}
	return bytesToConv
}

// Convert bytes to Base64 string
func ConvertBytesToBase64(bytesToConv []byte) string {
	base64Code := base64.StdEncoding.EncodeToString(bytesToConv)
	return base64Code
}
