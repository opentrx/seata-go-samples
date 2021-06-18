package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/opentrx/seata-golang/v2/pkg/client"
	"github.com/opentrx/seata-golang/v2/pkg/client/config"
	"github.com/opentrx/seata-golang/v2/pkg/client/tm"

	"github.com/opentrx/seata-go-samples/aggregation_svc/svc"
)


func main() {
	r := gin.Default()

	configPath := os.Getenv("ConfigPath")
	conf := config.InitConfiguration(configPath)
	client.Init(conf)

	tm.Implement(svc.ProxySvc)

	r.GET("/createSoCommit", func(c *gin.Context) {

		svc.ProxySvc.CreateSo(c, false)

		c.JSON(200, gin.H{
			"success": true,
			"message": "success",
		})
	})

	r.GET("/createSoRollback", func(c *gin.Context) {

		svc.ProxySvc.CreateSo(c, true)

		c.JSON(200, gin.H{
			"success": true,
			"message": "success",
		})
	})

	r.Run(":8003")
}
