package go_cheezeepay

import (
	"github.com/asaka1234/go-cheezeepay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	params CheezeePayInitParams

	ryClient *resty.Client
	logger   utils.Logger
	rsaUtil  utils.CheezeebitRSASignatureUtil
}

func NewClient(logger utils.Logger, params CheezeePayInitParams) *Client {
	return &Client{
		params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
		rsaUtil:  utils.CheezeebitRSASignatureUtil{},
	}
}
