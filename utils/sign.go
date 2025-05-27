package utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"sort"
	"strings"
)

type CheezeebitRSASignatureUtil struct{}

func (util *CheezeebitRSASignatureUtil) GetSign(paramsMap map[string]interface{}, privateKey string) (string, error) {
	delete(paramsMap, "platSign")
	textContent := util.GetContent(paramsMap)
	return util.Sign(textContent, privateKey)
}

func (util *CheezeebitRSASignatureUtil) VerifySign(paramsMap map[string]interface{}, publicKey string, sign string) (bool, error) {
	delete(paramsMap, "platSign")
	textContent := util.GetContent(paramsMap)
	return util.Verify(textContent, sign, publicKey)
}

//-------------------------------------------------------------

func (util *CheezeebitRSASignatureUtil) GetContent(paramsMap map[string]interface{}) string {
	// Get sorted keys
	keys := lo.Keys(paramsMap)
	sort.Strings(keys)

	var pairs []string
	lo.ForEach(keys, func(x string, index int) {
		value := ""
		if x != "payeeAccountInfos" {
			if x == "agentOrderBatch" {
				valueByte, _ := json.Marshal(paramsMap[x])
				value = string(valueByte)
			} else {
				value = cast.ToString(paramsMap[x])
			}
		}

		if value != "" {
			pairs = append(pairs, value)
		}
	})

	queryString := strings.Join(pairs, "")
	fmt.Printf("[rawString]%s\n", queryString)

	return queryString
}

func (util *CheezeebitRSASignatureUtil) Sign(message, privateKeyString string) (string, error) {

	signResult, err := SignSHA256RSA([]byte(message), privateKeyString)
	if err != nil {
		fmt.Printf("==sign===>%s\n", err.Error())
	}
	return signResult, nil
}

func (util *CheezeebitRSASignatureUtil) Verify(message, signatureString, publicKeyString string) (bool, error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyString)
	if err != nil {
		return false, err
	}

	// Parse public key
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		return false, err
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return false, errors.New("not an RSA public key")
	}

	signature, err := base64.StdEncoding.DecodeString(signatureString)
	if err != nil {
		return false, err
	}

	hashed := sha256.Sum256([]byte(message))
	err = rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return false, nil // Verification failed, but not an error condition
	}

	return true, nil
}
