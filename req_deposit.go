package go_cheezeepay

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-cheezeepay/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
)

// https://pay-apidoc-en.cheezeebit.com/#p2p-payin-order
func (cli *Client) Deposit(req CheezeePayDepositReq) (*CheezeePayDepositResponse, error) {

	rawURL := cli.Params.DepositUrl

	// 1. 拿到请求参数，转为map
	var signDataMap map[string]interface{}
	mapstructure.Decode(req, &signDataMap)
	signDataMap["merchantsId"] = cli.Params.MerchantId
	signDataMap["pushAddress"] = cli.Params.DepositCallbackUrl
	signDataMap["takerType"] = "2"
	signDataMap["coin"] = "USDT"
	signDataMap["tradeType"] = "2"
	signDataMap["language"] = "en"
	//为了确保amount必然是int整数,这里做一下强保障
	amount, _ := decimal.NewFromString(req.DealAmount)
	signDataMap["dealAmount"] = amount.StringFixed(0)

	// 2. 计算签名,补充参数
	signStr, _ := cli.rsaUtil.GetSign(signDataMap, cli.Params.RSAPrivateKey) //私钥加密
	signDataMap["platSign"] = signStr

	var result CheezeePayDepositResponse

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(signDataMap).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp))
	cli.logger.Infof("PSPResty#cheezeepay#deposit->%+v", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp.StatusCode())
	}

	if resp.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp.Error(), resp.Body())
	}

	/*
		//验证签名
		if result.Code == "000000" {
			sign := result.PlatSign //收到的签名

			var signResultMap map[string]interface{}
			mapstructure.Decode(result, &signResultMap)
			delete(signResultMap, "platSign") //去掉，用余下的来计算签名

			verify, _ := cli.rsaUtil.VerifySign(signResultMap, cli.Params.RSAPublicKey, sign) //公钥解密
			if !verify {
				return nil, fmt.Errorf("sign verify failed")
			}
		}
	*/

	return &result, nil
}
