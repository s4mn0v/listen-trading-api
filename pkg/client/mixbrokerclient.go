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
	return p.BitgetRestClient.DoGet("/api/v2/copy/mix-broker/query-traders", params)
}

func (p *MixBrokerClient) QueryCurrentTrack(traderId string) (string, error) {
	params := map[string]string{
		"traderId":    traderId,
		"productType": "usdt-futures", // HARDCORE FIX LATER TO BE GLOBAL
	}
	return p.BitgetRestClient.DoGet("/api/v2/copy/mix-broker/query-current-traces", params)
}
