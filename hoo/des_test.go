package hoo

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestEncryptDES_ECB_PKCS5Padding(t *testing.T) {
	key := []byte("12345678")
	plain := []byte("des_ecb_pkcs5padding")
	secret, err := EncryptDES_ECB_PKCS5Padding(key, plain)
	if err != nil {
		t.Error(err)
	}
	str := base64.StdEncoding.EncodeToString(secret)
	fmt.Println(str)

	if str != "q4J+bU51FPiG7sLWpJrUarY+JvWYepkG" {
		t.Error("Encryption failure")
	}
}

func TestDecryptDES_ECB_PKCS5Padding(t *testing.T) {
	key := []byte("12345678")
	str := "q4J+bU51FPiG7sLWpJrUarY+JvWYepkG"
	secret, _ := base64.StdEncoding.DecodeString(str)
	plain, _ := DecryptDES_ECB_PKCS5Padding(key, secret)
	fmt.Println(string(plain))

	if string(plain) != "des_ecb_pkcs5padding" {
		t.Error("Decryption failure")
	}
}

func TestDecryptDES_ECB_PKCS5Padding2(t *testing.T) {
	key := []byte("********")
	str := "1f9xE7+7xrM=" // empty string
	secret, _ := base64.StdEncoding.DecodeString(str)
	plain, _ := DecryptDES_ECB_PKCS5Padding(key, secret)
	fmt.Printf("decrypted result is [%s]\n", string(plain))
}
