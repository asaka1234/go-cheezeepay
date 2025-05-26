package go_cheezeepay

// ---------------------------------------------
type CheezeePayDepositReq struct {
	CustomerMerchantsId string `json:"customerMerchantsId" mapstructure:"customerMerchantsId"` //商户侧的uid
	LegalCoin           string `json:"legalCoin" mapstructure:"legalCoin"`                     //法定货币. 只支持: INR(印度卢比) IDR(印尼盾)
	MerchantOrderId     string `json:"merchantOrderId" mapstructure:"merchantOrderId"`         //商户订单号
	DealAmount          string `json:"dealAmount" mapstructure:"dealAmount"`                   //数量
	Language            string `json:"language" mapstructure:"language"`                       //zh_hk Chinese；VI Vietnamese；en English；Indonesia Indonesian
	//以下sdk来赋值
	//MerchantsId string `json:"merchantsId" mapstructure:"merchantsId"` //商户id
	//PushAddress string `json:"pushAddress" mapstructure:"pushAddress"` //回调地址
	//TakerType   string `json:"takerType" mapstructure:"takerType"`     // Fixed: 2
	//Coin        string `json:"coin" mapstructure:"coin"`               //Fixed: USDT
	//TradeType   string `json:"tradeType" mapstructure:"tradeType"`     //Fixed: 2
	//PlatSign    string `json:"platSign" mapstructure:"platSign"`       //签名(rsa私钥加密)
}

//-------------------------------

type CheezeePayDepositResponse struct {
	Success  bool                           `json:"success" mapstructure:"success"`
	Code     string                         `json:"code" mapstructure:"code"` // 000000 成功
	Msg      string                         `json:"msg" mapstructure:"msg"`
	Data     *CheezeePayDepositResponseData `json:"data,omitempty" mapstructure:"data"`
	ErrorMsg string                         `json:"errorMsg,omitempty" mapstructure:"errorMsg"`
	PlatSign *string                        `json:"platSign,omitempty" mapstructure:"platSign"` //签名,需要校验. 要用rsa 公钥
}

type CheezeePayDepositResponseData struct {
	OrderId string `json:"orderId" mapstructure:"orderId"`
	Type    int    `json:"type" mapstructure:"type"` // 0 for new order, 1 for existing order
	Url     string `json:"url" mapstructure:"url"`
}

//--------------callback------------------------------

type CheezeePayDepositBackReq struct {
	MerchantsOrderId string                        `json:"merchantsOrderId" mapstructure:"merchantsOrderId"` //商户订单号
	OrderId          string                        `json:"orderId" mapstructure:"orderId"`                   //psp的订单号
	MerchantId       string                        `json:"merchantId" mapstructure:"merchantId"`             //商户号
	Data             *CheezeePayDepositBackReqData `json:"data,omitempty" mapstructure:"data"`
}

type CheezeePayDepositBackReqData struct {
	OrderId             string `json:"orderId" mapstructure:"orderId"`
	Status              string `json:"status" mapstructure:"status"` //4 for success, 5 for failure, 6 for failure (user has not operated for 6 hours), 7 for failure (price not accepted), 9 for failure (refund)
	Coin                string `json:"coin" mapstructure:"coin"`
	DealAmount          string `json:"dealAmount" mapstructure:"dealAmount"`
	DealQuantity        string `json:"dealQuantity" mapstructure:"dealQuantity"`
	EntrustOrderId      string `json:"entrustOrderId" mapstructure:"entrustOrderId"`
	FeeCoin             string `json:"feeCoin" mapstructure:"feeCoin"`
	LegalCoin           string `json:"legalCoin" mapstructure:"legalCoin"`
	Price               string `json:"price" mapstructure:"price"`
	TakerFee            string `json:"takerFee" mapstructure:"takerFee"`
	TakerId             string `json:"takerId" mapstructure:"takerId"`
	TakerName           string `json:"takerName" mapstructure:"takerName"`
	TradeType           string `json:"tradeType" mapstructure:"tradeType"`
	PayWayName          string `json:"payWayName" mapstructure:"payWayName"`
	Side                string `json:"side" mapstructure:"side"`
	CustomerMerchantsId string `json:"customerMerchantsId" mapstructure:"customerMerchantsId"`
}

// 给callback的response
type CheezeePayDepositBackResp struct {
	Code int `json:"code"` // 响应状态码  200成功 . 返回`{"code":"200"}`给psp
}

//==============================================

type CheezeePayWithdrawReq struct {
	CustomerMerchantsId  string             `json:"customerMerchantsId" mapstructure:"customerMerchantsId"` //商户的userId
	LegalCoin            string             `json:"legalCoin" mapstructure:"legalCoin"`                     //法定货币. 只支持: INR(印度卢比) IDR(印尼盾)
	MerchantOrderId      string             `json:"merchantOrderId" mapstructure:"merchantOrderId"`         //商户订单号
	DealAmount           string             `json:"dealAmount" mapstructure:"dealAmount"`
	Language             string             `json:"language" mapstructure:"language"`                         //zh_hk Chinese；VI Vietnamese；en English；Indonesia Indonesian
	TakerName            string             `json:"takerName" mapstructure:"takerName"`                       //[Bank Transfer(India)]
	PayeeAccountType     string             `json:"payeeAccountType" mapstructure:"payeeAccountType"`         //Payment method, for example: [Bank Transfer(India)]
	PayeeAccountTypeName string             `json:"payeeAccountTypeName" mapstructure:"payeeAccountTypeName"` //Payment method name, for example: Bank Transfer(India)
	PayeeAccountInfos    []PayeeAccountInfo `json:"payeeAccountInfos" mapstructure:"payeeAccountInfos"`       //不参与签名计算！！！
	//sdk来做
	//MerchantsId string `json:"merchantsId" mapstructure:"merchantsId"`
	//PushAddress string `json:"pushAddress" mapstructure:"pushAddress"` //回调地址
	//TakerType string `json:"takerType" mapstructure:"takerType"` //Fixed: 2
	//Coin        string `json:"coin" mapstructure:"coin"`           //Fixed: USDT
	//TradeType   string `json:"tradeType" mapstructure:"tradeType"` //Fixed: 1
	//PlatSign    string `json:"platSign" mapstructure:"platSign"`   //签名
}

type PayeeAccountInfo struct {
	Field    string `json:"field" mapstructure:"field"`       //BANK_TRANSFER_INDIA_FIELD1
	Type     string `json:"type" mapstructure:"type"`         //text
	Required bool   `json:"required" mapstructure:"required"` //true
	Value    string `json:"value" mapstructure:"value"`       //***Account holder name***
}

type CheezeePayWithdrawResp struct {
	Success  bool          `json:"success" mapstructure:"success"`
	Code     string        `json:"code" mapstructure:"code"`
	Msg      string        `json:"msg" mapstructure:"msg"`
	Data     *ResponseData `json:"data" mapstructure:"data"`         //失败的话, null
	PlatSign string        `json:"platSign" mapstructure:"platSign"` //失败的话不返回该字段
}

type ResponseData struct {
	OrderId string `json:"orderId" mapstructure:"orderId"`
}

//-----callback---------------

type CheezeePayWithdrawBackReq struct {
	OrderId          string            `json:"orderId" mapstructure:"orderId"`                   //Platform order number.
	MerchantsOrderId string            `json:"merchantsOrderId" mapstructure:"merchantsOrderId"` //Merchant order number
	MerchantId       string            `json:"merchantId" mapstructure:"merchantId"`
	Data             WithdrawOrderData `json:"data" mapstructure:"data"`
}

type WithdrawOrderData struct {
	OrderId             string `json:"orderId" mapstructure:"orderId"`
	Status              string `json:"status" mapstructure:"status"` //4 for success, 5 for failure, 6 for failure (user has not operated for 6 hours), 7 for failure (price not accepted), 9 for failure (refund)
	Coin                string `json:"coin" mapstructure:"coin"`
	DealAmount          string `json:"dealAmount" mapstructure:"dealAmount"`
	DealQuantity        string `json:"dealQuantity" mapstructure:"dealQuantity"`
	EntrustOrderId      string `json:"entrustOrderId" mapstructure:"entrustOrderId"`
	FeeCoin             string `json:"feeCoin" mapstructure:"feeCoin"`
	LegalCoin           string `json:"legalCoin" mapstructure:"legalCoin"`
	Price               string `json:"price" mapstructure:"price"`
	TakerFee            string `json:"takerFee" mapstructure:"takerFee"`
	TakerId             string `json:"takerId" mapstructure:"takerId"`
	TakerName           string `json:"takerName" mapstructure:"takerName"`
	TradeType           string `json:"tradeType" mapstructure:"tradeType"`
	PayWayName          string `json:"payWayName" mapstructure:"payWayName"`
	Side                string `json:"side" mapstructure:"side"`
	CustomerMerchantsId string `json:"customerMerchantsId" mapstructure:"customerMerchantsId"` ////商户的userId
}

// 给callback的response
type CheezeePayWithdrawBackResp struct {
	Code int `json:"code"` // 响应状态码  200成功
}
