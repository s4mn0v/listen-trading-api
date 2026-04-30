package client

import (
	"github.com/s4mn0v/listen-trading-api/internal/common"
)

type MixBrokerClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *MixBrokerClient) Init() *MixBrokerClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *MixBrokerClient) QueryTraders(pageSize, pageNo string) (string, error) {
	params := map[string]string{
		"pageSize": pageSize,
		"pageNo":   pageNo,
	}
	// Endpoint V2 para listar traders
	resp, err := p.BitgetRestClient.DoGet("/api/v2/copy/mix-broker/query-traders", params)
	return resp, err
}
