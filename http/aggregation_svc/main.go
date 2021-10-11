package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/transaction-wg/seata-golang/pkg/client"
	"github.com/transaction-wg/seata-golang/pkg/client/config"
	"github.com/transaction-wg/seata-golang/pkg/client/tm"
	"path/filepath"
	"runtime"
)

import (
	"github.com/opentrx/seata-go-samples/aggregation_svc/svc"
)

func main() {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("caller err")
		return
	}

	configPath := filepath.Dir(path) + string(filepath.Separator) + "conf" + string(filepath.Separator) + "client.yml"
	fmt.Println(configPath)

	r := gin.Default()
	config.InitConf(configPath)
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
