package main

import (
	ctx "github.com/opentrx/seata-golang/v2/pkg/client/base/context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/opentrx/seata-go-samples/service"
	"github.com/opentrx/seata-golang/v2/pkg/client"
	"github.com/opentrx/seata-golang/v2/pkg/client/config"
	"github.com/opentrx/seata-golang/v2/pkg/client/tcc"
)

func main() {
	r := gin.Default()

	configPath := os.Getenv("ConfigPath")
	config := config.InitConfiguration(configPath)
	client.Init(config)

	tcc.ImplementTCC(service.TccProxyServiceB)

	r.GET("/try", func(c *gin.Context) {
		rootContext := ctx.NewRootContext(c)
		rootContext.Bind(c.GetHeader("xid"))

		businessActionContextB := &ctx.BusinessActionContext{
			RootContext:   rootContext,
			ActionContext: make(map[string]interface{}),
		}
		businessActionContextB.ActionContext["hello"] = "hello world,this is from BusinessActionContext B"

		service.TccProxyServiceB.Try(businessActionContextB)

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8082")
}
