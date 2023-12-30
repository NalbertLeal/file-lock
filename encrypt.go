package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

var iv = []byte("123456789abcdefg")

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func encrypt(content []byte, key []byte) ([]byte, error) {
	aesChipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(aesChipher, iv)

	content = PKCS5Padding(content, aes.BlockSize)

	ciphertext := make([]byte, len(content))
	mode.CryptBlocks(ciphertext, content)

	return ciphertext, nil
}

func decrypt(content []byte, key []byte) ([]byte, error) {
	aesChipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}

	mode := cipher.NewCBCDecrypter(aesChipher, iv)

	result := make([]byte, len(content))
	mode.CryptBlocks(result, content)

	result = PKCS5Trimming(result)

	return result, nil
}
