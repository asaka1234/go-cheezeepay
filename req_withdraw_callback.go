package go_cheezeepay

import (
	"bytes"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// https://pay-apidoc-en.cheezeebit.com/#p2p-payout-notification
func (cli *Client) WithdrawCallback(req CheezeePayWithdrawBackReq, processor func(CheezeePayWithdrawBackReq) error) error {
	//验证签名
	sign := req.PlatSign //收到的签名

	var signResultMap map[string]interface{}
	mapstructure.Decode(req, &signResultMap)
	delete(signResultMap, "platSign") //去掉，用余下的来计算签名

	verify, _ := cli.rsaUtil.VerifySign(signResultMap, cli.Params.RSAPublicKey, sign) //私钥加密
	if !verify {
		return fmt.Errorf("sign verify failed")
	}

	//映射一下
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer([]byte(req.DataRaw)))
	if err != nil {
		return err
	}
	return viper.Unmarshal(req.Data)
	
	//开始处理
	return processor(req)
}
