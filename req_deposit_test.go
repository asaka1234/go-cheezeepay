package go_cheezeepay

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestDeposit(t *testing.T) {
	vlog := VLog{}
	//构造client
	cli := NewClient(vlog, &CheezeePayInitParams{MERCHANT_ID, RSA_PUBLIC_KEY, RSA_PRIVATE_KEY, DEPOST_URL, DEPOST_BACK_URL, WITHDRAW_URL, WITHDRAW_BACK_URL})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() CheezeePayDepositReq {
	return CheezeePayDepositReq{
		CustomerMerchantsId: "12345", //商户uid
		LegalCoin:           "INR",
		MerchantOrderId:     "8787791",
		DealAmount:          "400.00", //不能浮点数
	}
}

func TestViper(t *testing.T) {
	req := CheezeePayDepositBackReq{
		MerchantsOrderId: "11",
		OrderId:          "22",
		MerchantId:       "33",
		PlatSign:         "GOXVDKv/Q6rln0xP8lyQVbvLSkPjRlcBw1Idb8itFIBOvhC/U2ZXh7PmO9EzIRLuHlRbY49w57LL3LxOrir+Xd1gAOrWgU+GdMYPvfsxHy2HxFrfXqIeNF3miLKhX+3ASnPvibhQaGwHV02CT29Tz79dgd2LvSz+mhw2PsYOcYI8C/uldadCiE2bZNUZTYz58xN573JkNvgFtCqJ+UriBG6Q5FUAA9rjvZ9aIkx+Bsm9ue8XlgruJXZsks8xBWofUgmGQJjYsBl7dVPRcC4OSwzO3eCnxT9Mhk0JSIs/TsZQgOdCZvyrQppYBJHPeGfnxFuDU+dyY3SH/UFZ4xnAGw==",
		DataRaw:          "{\"coin\":\"USDT\",\"customerMerchantsId\":\"820002060\",\"dealAmount\":\"368.000000000000000000\",\"dealQuantity\":\"4.600000000000000000\",\"entrustOrderId\":2025021814074617800,\"extValues\":{},\"feeCoin\":\"USDT\",\"gmtCreate\":1750150128000,\"gmtEnd\":1750150294000,\"legalCoin\":\"INR\",\"merchantOrderId\":\"202506171148470593\",\"orderId\":1934895991828647936,\"payWayName\":\"[UPI]\",\"price\":\"80.000000000000000000\",\"side\":\"C2C\",\"status\":4,\"takerFee\":\"0.165600000000000000\",\"takerId\":\"CH1000114300000013\",\"tradeType\":2}"}

	viper.SetConfigType("json")
	err := viper.ReadConfig(bytes.NewBuffer([]byte(req.DataRaw)))
	if err != nil {
		return
	}

	var cc CheezeePayDepositBackReqData
	viper.Unmarshal(&cc)
	req.Data = cc

	fmt.Printf("==>%+v\n", req.Data)
}
