package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/opentrx/seata-golang/v2/pkg/client"
	"github.com/opentrx/seata-golang/v2/pkg/client/config"
	"github.com/opentrx/seata-golang/v2/pkg/client/tm"
	"github.com/opentrx/seata-golang/v2/pkg/util/log"

	"net/http"
	_ "net/http/pprof"

	"github.com/opentrx/seata-go-samples/aggregation_svc/svc"
)

func init() {
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
}

func main() {
	r := gin.Default()

	configPath := os.Getenv("ConfigPath")
	conf := config.InitConfiguration(configPath)

	log.Init(conf.Log.LogPath, conf.Log.LogLevel)
	client.Init(conf)

	tm.Implement(svc.ProxySvc)

	r.GET("/createSoCommit", func(c *gin.Context) {

		if err := svc.ProxySvc.CreateSo(c, false); err == nil {
			c.JSON(200, gin.H{
				"success": true,
				"message": "success",
			})
		} else {
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}

	})

	r.GET("/createSoRollback", func(c *gin.Context) {

		if err := svc.ProxySvc.CreateSo(c, true); err == nil {
			c.JSON(200, gin.H{
				"success": true,
				"message": "success",
			})
		} else {
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}
	})

	r.Run(":8003")
}
