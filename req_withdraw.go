package go_cheezeepay

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-cheezeepay/utils"
	"github.com/mitchellh/mapstructure"
)

// https://pay-apidoc-en.cheezeebit.com/#p2p-payout-order
func (cli *Client) Withdraw(req CheezeePayWithdrawReq) (*CheezeePayWithdrawResp, error) {

	rawURL := cli.DepositURL
	// 1. 拿到请求参数，转为map
	var signDataMap map[string]interface{}
	mapstructure.Decode(req, &signDataMap)
	signDataMap["merchantsId"] = cli.MerchantID
	signDataMap["pushAddress"] = cli.DepositCallbackURL
	signDataMap["takerType"] = "2"
	signDataMap["coin"] = "USDT"
	signDataMap["tradeType"] = "1"

	// 2. 计算签名,补充参数
	signStr, _ := utils.GetSign(signDataMap, cli.RSAPrivateKey) //私钥加密
	signDataMap["platSign"] = signStr

	var result CheezeePayWithdrawResp

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(signDataMap).
		SetHeaders(getHeaders()).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	//fmt.Printf("result: %s\n", string(resp.Body()))

	if err != nil {
		return nil, err
	}

	if result.Code == "000000" {
		var signResultMap map[string]interface{}
		mapstructure.Decode(result, &signResultMap)

		verify, _ := utils.VerifySign(signResultMap, cli.RSAPublicKey) //公钥解密
		if !verify {
			return nil, fmt.Errorf("sign verify failed")
		}
	}

	return &result, nil
}
