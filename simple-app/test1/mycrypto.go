package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func MyCryptoDemo() {
	data := "Hello,my friends"

	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	fmt.Printf("1.md5加密%s:%x\n", data, md5Hash.Sum(nil))

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(data))
	fmt.Printf("2.sha1加密%s:%x\n", data, sha1Hash.Sum(nil))

	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(data))
	fmt.Printf("3.sha256加密%s:%x\n", data, sha256Hash.Sum(nil))

	sha512Hash := sha512.New()
	sha512Hash.Write([]byte(data))
	fmt.Printf("4.sha512加密%s:%x\n", data, sha512Hash.Sum(nil))

	data1 := "my password is 123456"
	data2 := "my password is 123456"
	data3 := "my password is 1234567"

	hash1 := sha256.Sum256([]byte(data1))
	hash2 := sha256.Sum256([]byte(data2))
	hash3 := sha256.Sum256([]byte(data3))

	fmt.Printf("data1:%s，hash1:%x\n", data1, hash1)
	fmt.Printf("data2:%s，hash2:%x\n", data2, hash2)
	fmt.Printf("data3:%s，hash3:%x\n", data3, hash3)

	fmt.Printf("data3:%s,hex:%s\n", data3, hex.EncodeToString(hash3[:]))

	secret := []byte("my secret key")
	message := "imput message"
	hmacHash := hmac.New(sha256.New, secret)
	hmacHash.Write([]byte(message))
	fmt.Printf("HMAC-SHA256 of '%s' : %x\n", message, hmacHash.Sum(nil))

	secret1 := []byte("my secret key")
	message1 := "imput message"
	hmacHash1 := hmac.New(sha256.New, secret1)
	hmacHash1.Write([]byte(message1))
	fmt.Printf("HMAC-SHA256 of '%s' : %x\n", message1, hmacHash1.Sum(nil))

	secret2 := []byte("my secret key1")
	message2 := "imput message"
	hmacHash2 := hmac.New(sha256.New, secret2)
	hmacHash2.Write([]byte(message2))
	fmt.Printf("HMAC-SHA256 of '%s' : %x\n", message1, hmacHash1.Sum(nil))
}
