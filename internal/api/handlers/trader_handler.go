package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/s4mn0v/listen-trading-api/pkg/client"
)

func ListTraders(c *gin.Context) {
	// Inicializar cliente siguiendo el estilo SDK solicitado
	clientV2 := new(client.MixBrokerClient).Init()

	// Stateless Proxy: Sobrescribir llaves si vienen en los headers del usuario
	if key := c.GetHeader("X-USER-KEY"); key != "" {
		clientV2.BitgetRestClient.ApiKey = key
		clientV2.BitgetRestClient.SecretKey = c.GetHeader("X-USER-SECRET")
		clientV2.BitgetRestClient.Passphrase = c.GetHeader("X-USER-PASS")
	}

	pageSize := c.DefaultQuery("pageSize", "10")
	pageNo := c.DefaultQuery("pageNo", "1")

	resp, err := clientV2.QueryTraders(pageSize, pageNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(resp))
}
