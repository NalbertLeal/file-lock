package main

import (
	"fmt"
	"os"
)

func doesFileExists(pathToExam string) bool {
	info, err := os.Stat(pathToExam)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readFileContent(filepath string) ([]byte, error) {
	return os.ReadFile(filepath)
}

// Verify if the file is encripted by looking the first 9 bytes (these 9 bytes
// must contain the word encripted in ASCII, []byte("encrypted") )
func isEncrypted(content []byte) bool {
	return content[0] == 101 && content[1] == 110 &&
		content[2] == 99 && content[3] == 114 && content[4] == 121 &&
		content[5] == 112 && content[6] == 116 && content[7] == 101 &&
		content[8] == 100
}

func writeToFile(path string, content []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Can't overwrite file on path: " + path)
	}
	defer file.Close()

	file.Write(content)
	file.Sync()

	return nil
}

func encryptFileContent(path string, content []byte, key []byte) error {
	encryptedContent, err := encrypt(content, key)
	if err != nil {
		return err
	}
	contentWithFlagWrite := []byte("encrypted")
	contentWithFlagWrite = append(contentWithFlagWrite, encryptedContent...)

	writeToFile(path, contentWithFlagWrite)
	return nil
}

func decryptFileContent(path string, content []byte, key []byte) error {
	decryptedContent, err := decrypt(content, key)
	if err != nil {
		return err
	}

	writeToFile(path, decryptedContent)

	return nil
}
