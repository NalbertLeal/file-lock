package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func createPassword(password string) string {
	const default_password = "XzFjaOhbMTHKN8Jk9rp3ulS2PLRqoEDs"

	if len(password) < 32 {
		finalPassword := password
		for i := len(password); i < 32; i++ {
			finalPassword += string(default_password[i])
		}

		return finalPassword
	}
	return password[0:32]
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Pass the file name and password key to be used")
		os.Exit(1)
	}

	dirpath, err := currentWorkingPath()
	if err != nil {
		fmt.Println("")
		os.Exit(1)
	}
	filename := os.Args[1]
	path := filepath.Join(dirpath, filename)

	if !doesFileExists(path) {
		fmt.Println("File  " + path + "  not found")
		os.Exit(1)
	}

	content, err := readFileContent(path)
	if err != nil {
		fmt.Println("Can't read the file content of " + path)
		os.Exit(1)
	}

	// key := []byte("N1PCdw3M2B1TfJhoaY2mL736p2vCUc47") //os.Args[2])
	password := os.Args[2]
	key := []byte(createPassword(password))
	if isEncrypted(content[0:9]) {
		contentWithNoFlag := content[9:]
		err = decryptFileContent(path, contentWithNoFlag, key)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		err = encryptFileContent(path, content, key)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
