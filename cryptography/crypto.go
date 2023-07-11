package cryptography

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"math/big"
)

func AESCBCEncrypt(plainText, key, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	plainText = pkcs5Padding(plainText, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	return cipherText
}

func AESCBCDecrypt(cipherText, key, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	plainText = pkcs5UnPadding(plainText)
	return plainText
}

func AESECBEncrypt(plainText, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	plainText = pkcs5Padding(plainText, block.BlockSize())
	blockMode := NewECBEncrypter(block)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	return cipherText
}

func AESECBDecrypt(cipherText, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockMode := NewECBDecrypter(block)
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	plainText = pkcs5UnPadding(plainText)
	return plainText
}

func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, paddingText...)
}

func pkcs5UnPadding(src []byte) []byte {
	n := len(src)
	unPadding := int(src[n-1])
	return src[:n-unPadding]
}

func RSAEncrypt(origData []byte, modulus string, exponent int64) string {
	bigOrigData := big.NewInt(0).SetBytes(origData)
	bigModulus, _ := big.NewInt(0).SetString(modulus, 16)
	bigRs := bigOrigData.Exp(bigOrigData, big.NewInt(exponent), bigModulus)
	return fmt.Sprintf("%0256x", bigRs)
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
