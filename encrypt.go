package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/eiannone/keyboard"
)

const (
	passphrase = "jyczzae342;fsd"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("\n")
		fmt.Print("Type your string : ")
		scanner.Scan()
		userInput := scanner.Text()
		userInput = strings.TrimSpace(userInput)
		str, _ := encrypt(userInput)
		fmt.Println("\n-------------------------------------------")
		fmt.Print(userInput)
		fmt.Print(" ----->  ")
		fmt.Print(str)
		fmt.Println("\n-------------------------------------------")
		fmt.Println("")

		err := keyboard.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer keyboard.Close()
		fmt.Print("Continue y/n: ")
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		char = unicode.ToLower(char)
		if char == 121 {
			continue
		} else if char == 110 {
			break
		} else {
			break
		}
	}
}
func encrypt(plaintext string) (string, error) {
	key := sha256.Sum256([]byte(passphrase))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	encryptedText := append(nonce, ciphertext...)
	encryptedHex := hex.EncodeToString(encryptedText)
	return encryptedHex, nil
}
