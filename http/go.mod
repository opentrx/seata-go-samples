module github.com/opentrx/seata-go-samples

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.2.0
	github.com/opentrx/mysql v1.0.0-pre
	github.com/opentrx/seata-golang/v2 v2.0.0
)

replace (
	github.com/opentrx/mysql => /Users/scottlewis/dksl/current/mysql
	github.com/opentrx/seata-golang/v2 => /Users/scottlewis/dksl/current/seata-golang/
)
