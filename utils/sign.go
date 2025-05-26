package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"sort"
	"strings"
)

const (
	DATA = "data"
)

// Generate signature
func GetSign(params map[string]interface{}, privateKey string) (string, error) {
	delete(params, "sign")
	textContent := GetContent(params)
	fmt.Println("Signing content:", textContent)
	return Sign(textContent, privateKey)
}

// Verify the signature
func VerifySign(params map[string]interface{}, publicKey string) (bool, error) {
	platSign, ok := params["platSign"].(string)
	if !ok {
		return false, errors.New("platSign not found or not a string")
	}
	delete(params, "platSign")

	textContent := GetContentNew(params)
	fmt.Println("Signature Verification:", textContent)
	return VerifySignature(textContent, platSign, publicKey)
}

//----------------------------

// The string for the reception signature
// 是value的拼接
func GetContent(params map[string]interface{}) string {
	keys := lo.Keys(params)
	sort.Strings(keys)

	var builder strings.Builder
	for _, name := range keys {
		//遍历每一个key
		if params[name] != nil && name != "payeeAccountInfos" {
			if name == "agentOrderBatch" {
				// builder.WriteString(JSON.toJSONString(params.get(name)))
			} else {
				//直接做value的拼接
				builder.WriteString(cast.ToString(params[name]))
			}
		}
	}
	return builder.String()
}

// Spell the string of the signature to be verified
func GetContentNew(params map[string]interface{}) string {
	keys := lo.Keys(params)
	sort.Strings(keys)

	var builder strings.Builder
	for _, name := range keys {
		if !isEmpty(params[name]) {
			if name == DATA {
				if dataValue, ok := params[name].(map[string]interface{}); ok {
					builder.WriteString(GetContentNew(dataValue))
				} else {
					builder.WriteString(fmt.Sprintf("%v", params[name]))
				}
			} else if name == "payeeAccountInfos" {
				if strValue, ok := params[name].(string); ok {
					builder.WriteString(strValue)
				}
			} else {
				if params[name] != nil {
					builder.WriteString(fmt.Sprintf("%v", params[name]))
				}
			}
		}
	}
	return builder.String()
}

// Check if value is empty
func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	switch v := value.(type) {
	case string:
		return v == ""
	case map[string]interface{}:
		return len(v) == 0
	case []interface{}:
		return len(v) == 0
	default:
		return false
	}
}

// Generate a signature using a private key
func Sign(textContent, privateKeyStr string) (string, error) {

	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return "", errors.New("failed to parse PEM block containing the private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	hashed := sha256.Sum256([]byte(textContent))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// Use public key to verify signature
func VerifySignature(textContent, signStr, publicKeyStr string) (bool, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return false, errors.New("failed to parse PEM block containing the public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return false, errors.New("not an RSA public key")
	}

	signature, err := base64.StdEncoding.DecodeString(signStr)
	if err != nil {
		return false, err
	}

	hashed := sha256.Sum256([]byte(textContent))
	err = rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return false, nil
	}

	return true, nil
}
