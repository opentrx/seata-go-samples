# seata-go-samples

## Step0: setup TC server
```bash
git clone git@github.com:opentrx/seata-golang.git
cd seata-golang

vim ./cmd/profiles/dev/config.yml
# update storage.mysql.dsn
# update log.logPath

# create database `seata` on mysql server
# mysql> CREATE database if NOT EXISTS `seata` default character set utf8mb4 collate utf8mb4_unicode_ci;

cd cmd/tc
go run main.go start -config ../profiles/dev/config.yml
```

- ## AT mode example (gorm or http)
### Step1: setup aggregation_svc client
```bash
cd seata-go-samples/gorm
vim ./aggregation_svc/conf/client.yml
# update log.logPath

export ConfigPath="./aggregation_svc/conf/client.yml"
go run aggregation_svc/main.go
```

### Step2: setup order_svc client
```bash
cd seata-go-samples/gorm
vim ./order_svc/conf/client.yml
# update at.dsn
# update log.logPath

export ConfigPath="./order_svc/conf/client.yml"
go run order_svc/main.go
```

### Step3: setup product_svc client
```bash
cd seata-go-samples/gorm
vim ./product_svc/conf/client.yml
# update at.dsn
# update log.logPath

export ConfigPath="./product_svc/conf/client.yml"
go run product_svc/main.go
```

### Step4: access
- http://localhost:8003/createSoCommit
- http://localhost:8003/createSoRollback

- ## TCC mode example (tcc)
### Step1: setup client
```bash
cd seata-go-samples/tcc

export ConfigPath="./conf/config.yml"
go run main.go
```

### step2: access
- http://localhost:8080/commit
- http://localhost:8080/rollback