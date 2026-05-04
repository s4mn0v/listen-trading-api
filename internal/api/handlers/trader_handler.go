package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/s4mn0v/listen-trading-api/internal/models"
	"github.com/s4mn0v/listen-trading-api/internal/storage"
	"github.com/s4mn0v/listen-trading-api/logging/applogger"
	"github.com/s4mn0v/listen-trading-api/pkg/client"
)

// ListTraders
func ListTraders(c *gin.Context) {
	clientV2 := new(client.MixBrokerClient).Init()

	// Stateless Proxy: Keys inyected
	if key := c.GetHeader("X-USER-KEY"); key != "" {
		clientV2.BitgetRestClient.ApiKey = key
		clientV2.BitgetRestClient.SecretKey = c.GetHeader("X-USER-SECRET")
		clientV2.BitgetRestClient.Passphrase = c.GetHeader("X-USER-PASS")
	}

	raw, err := clientV2.QueryTraders(c.DefaultQuery("pageSize", "10"), c.DefaultQuery("pageNo", "1"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Bitget error"})
		return
	}

	var bgRes models.TraderListResponse
	json.Unmarshal([]byte(raw), &bgRes)

	// Save mongo in "traders" collection
	go storage.SaveTraders(bgRes.Data)

	// Intelligent mapping
	var nexusList []gin.H
	for _, t := range bgRes.Data {
		roi, mdd := "0", "0"
		for _, col := range t.ColumnList {
			if col.Describe == "ROI" {
				roi = col.Value
			}
			if col.Describe == "MDD" {
				mdd = col.Value
			}
		}

		fProfit, _ := strconv.ParseFloat(t.FollowerTotalProfit, 64)

		nexusList = append(nexusList, gin.H{
			"trader_id":       t.TraderId,
			"name":            t.TraderName,
			"avatar":          t.TraderPic,
			"follower_profit": fmt.Sprintf("%.2f USDT", fProfit),
			"roi":             roi + "%",
			"mdd":             mdd + "%",
			"is_honest":       fProfit > 0,
			"spots":           t.FollowCount + "/" + t.MaxLimit,
			"bias":            "LONG",
		})
	}

	c.JSON(200, gin.H{"status": "success", "data": nexusList})
}

// TraderDetail: Open positions by ID
func TraderDetail(c *gin.Context) {
	traderId := c.Param("id")
	clientV2 := new(client.MixBrokerClient).Init()

	// Stateless Proxy
	if key := c.GetHeader("X-USER-KEY"); key != "" {
		clientV2.BitgetRestClient.ApiKey = key
		clientV2.BitgetRestClient.SecretKey = c.GetHeader("X-USER-SECRET")
		clientV2.BitgetRestClient.Passphrase = c.GetHeader("X-USER-PASS")
	}

	raw, err := clientV2.QueryCurrentTrack(traderId)
	if err != nil {
		applogger.Error("Error Bitget: %v", err)
		c.JSON(500, gin.H{"error": "Bitget error"})
		return
	}

	fmt.Printf("DEBUG ID [%s] -> RESPONSE: %s\n", traderId, raw)

	var posRes models.CurrentTrackResponse
	json.Unmarshal([]byte(raw), &posRes)

	if posRes.Code != "00000" {
		c.JSON(400, gin.H{"status": "error", "code": posRes.Code, "msg": posRes.Msg})
		return
	}

	positions := posRes.Data
	if positions == nil {
		positions = []models.PositionInfo{}
	}

	go storage.SavePositions(traderId, positions)

	c.JSON(200, gin.H{
		"status":         "success",
		"trader_id":      traderId,
		"live_positions": positions,
	})
}
