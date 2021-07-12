module github.com/opentrx/seata-go-samples

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.2.0
	github.com/opentrx/mysql/v2 v2.0.0-rc
	github.com/dk-lockdown/harmonia v1.0.0
)

replace (
	github.com/dk-lockdown/harmonia => /Users/scottlewis/dksl/current/harmonia
	github.com/opentrx/mysql/v2 => /Users/scottlewis/dksl/current/mysql
)