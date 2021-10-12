module github.com/opentrx/seata-go-samples

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.2.0
	github.com/transaction-wg/seata-golang v1.0.0-rc2 //需要改为dev新版本的
)

replace github.com/transaction-wg/seata-golang => ../../seata-golang
