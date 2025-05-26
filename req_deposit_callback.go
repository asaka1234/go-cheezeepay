package go_cheezeepay

// https://pay-apidoc-en.cheezeebit.com/#p2p-payin-notification
func (cli *Client) DepositCallback(req CheezeePayDepositBackReq, processor func(CheezeePayDepositBackReq) error) error {
	//TODO  貌似官方也没有验证签名

	//开始处理
	return processor(req)
}
