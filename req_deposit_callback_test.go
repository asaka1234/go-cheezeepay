package go_cheezeepay

import (
	"fmt"
	"testing"
)

func TestDepositCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &CheezeePayInitParams{MERCHANT_ID, RSA_PUBLIC_KEY, RSA_PRIVATE_KEY, DEPOST_URL, DEPOST_BACK_URL, WITHDRAW_URL, WITHDRAW_BACK_URL})

	//发请求
	err := cli.DepositCallback(GenDepositCallbackRequestDemo(), processor)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
}

func GenDepositCallbackRequestDemo() CheezeePayDepositBackReq {
	return CheezeePayDepositBackReq{
		MerchantsOrderId: "202507091154070900", //商户uid
		OrderId:          "1942869868508745728",
		MerchantId:       "CH10001053",
		PlatSign:         "qmtwU0X2G1lKAuHaWgqgHEAdnATE0cG1dEJTcaqfD+XFViErAgxIu4DTYyES862MW/dRf2MWdZ0hxA/BvkcZAfd5dxrqV1blinNOACwO1xB94OLnK9yUp6NcnBIiCLurTm0OPIXMXUjqTi9nCdbnSnilKC8c8lYNxR24l06ahm1CawEugKexQ9zcFbbl5Z4B79I6k2en7vqp+QeYdHtlVsGulABsxo6VTx/hbMBydw79+0SY6f3ONV2NDEC2j8MTZyCqnRS9yAwAZoC0ielvGh8qhcmUkYJZbVlUT5U7GrNz+MPoMgoqTcKaNcwfBFOqS3pnLkfjgQKEon7w6llGaQ==",
		DataRaw:          "{\"coin\":\"USDT\",\"customerMerchantsId\":\"820000083\",\"dealAmount\":\"900.000000000000000000\",\"dealQuantity\":\"9.799651000000000000\",\"entrustOrderId\":2025051210331553100,\"extValues\":{},\"feeCoin\":\"USDT\",\"gmtCreate\":1752051248000,\"gmtEnd\":253370736000000,\"legalCoin\":\"INR\",\"merchantOrderId\":\"202507091154070900\",\"orderId\":1942869868508745728,\"payWayName\":\"[Paytm]\",\"price\":\"91.840000000000000000\",\"side\":\"C2C\",\"status\":5,\"takerFee\":\"0.352787000000000000\",\"takerId\":\"CH1000105300000700\",\"tradeType\":2}",
	}
}

func processor(CheezeePayDepositBackReq) error {
	return nil
}
