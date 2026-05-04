package models

type TraderListResponse struct {
	Code        string       `json:"code"`
	Msg         string       `json:"msg"`
	RequestTime int64        `json:"requestTime"`
	Data        []TraderInfo `json:"data"`
}

type TraderInfo struct {
	TraderId           string       `json:"traderId"`
	TraderName         string       `json:"traderName"`
	CanTrace           string       `json:"canTrace"`
	WinRate            string       `json:"winRate,omitempty"`
	ROI                string       `json:"roi,omitempty"`
	MaxLimit           string       `json:"maxLimit"`
	FollowCount        string       `json:"followCount"`
	TraderStatus       string       `json:"traderStatus"`
	CurrentTradingList []string     `json:"currentTradingList"`
	ColumnList         []ColumnItem `json:"columnList"`
	TotalFollowers     string       `json:"totalFollowers"`
	ProfitCount        string       `json:"profitCount"`
	LossCount          string       `json:"lossCount"`
	TradeCount         string       `json:"tradeCount"`
	TraderPic          string       `json:"traderPic"`
	MaxCallbackRate    string       `json:"maxCallbackRate"`
	FollowerTotalProfit string      `json:"followerTotalProfit"`
	TradeDays          string       `json:"tradeDays"`
}

type ColumnItem struct {
	Describe string `json:"describe"`
	Value    string `json:"value"`
}

type CurrentTrackResponse struct {
	Code string        `json:"code"`
	Msg  string        `json:"msg"`
	Data []PositionInfo `json:"data"`
}

type PositionInfo struct {
	Symbol       string `json:"symbol" bson:"symbol"`
	PosSide      string `json:"posSide" bson:"posSide"`
	OpenPriceAvg string `json:"openPriceAvg" bson:"openPriceAvg"`
	OpenLeverage string `json:"openLeverage" bson:"openLeverage"`
	TrackingNo   string `json:"trackingNo" bson:"trackingNo"`
	CTime        string `json:"cTime" bson:"cTime"`
    TraderId     string `json:"traderId" bson:"traderId"`
}
