package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/transaction-wg/seata-golang/pkg/base/config_center/nacos"
	_ "github.com/transaction-wg/seata-golang/pkg/base/registry/nacos"
	"github.com/transaction-wg/seata-golang/pkg/client"
	"github.com/transaction-wg/seata-golang/pkg/client/config"
	"github.com/transaction-wg/seata-golang/pkg/client/tm"
)

import (
	"github.com/opentrx/seata-go-samples/aggregation_svc/svc"
)

var configPath = "/Users/scottlewis/dksl/temp/seata-samples/http/aggregation_svc/conf/client.yml"

func main() {
	r := gin.Default()
	config.InitConf()
	client.NewRpcClient()
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
