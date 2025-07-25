package go_cheezeepay

import (
	"bytes"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// https://pay-apidoc-en.cheezeebit.com/#p2p-payin-notification
func (cli *Client) DepositCallback(req CheezeePayDepositBackReq, processor func(CheezeePayDepositBackReq) error) error {
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
	viper.SetConfigType("json")
	err := viper.ReadConfig(bytes.NewBuffer([]byte(req.DataRaw)))
	if err != nil {
		return err
	}
	var data CheezeePayDepositBackReqData
	viper.Unmarshal(&data)
	req.Data = data

	//开始处理
	return processor(req)
}
