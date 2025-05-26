package go_cheezeepay

// https://pay-apidoc-en.cheezeebit.com/#p2p-payout-notification
func (cli *Client) WithdrawCallback(req CheezeePayWithdrawBackReq, processor func(CheezeePayWithdrawBackReq) error) error {
	//TODO  貌似官方也没有验证签名

	//开始处理
	return processor(req)
}
