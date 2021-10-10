# http项目

## 模块
**aggregation_svc** ：seata系统中的TM，端口8003，接受/createSoCommit和/createSoRollback请求。

**order_svc**：seata系统中的RM，端口8002，接受/createSo请求，由**aggregation_svc**自动调用。

**product_svc**：seata系统中的另一个RM，端口8001，接受/allocateInventory请求，由**aggregation_svc**自动调用。

## 准备
1，请在执行sample前，将 **scripts** 文件夹下的表结构创建好。

2，本sample支持 **seata-server-1.2.0** 以上的版本

3，修改好3个模块目录下的**conf/client.yml**配置，大部分情况下，你只需要关心数据库连接的配置

## 步骤
1，启动好seata-sever作为TC

2，分别启动3个模块

3，请求 http://localhost:8003/createSoCommit ，查看数据中数据的commit情况

4，请求 http://localhost:8003/createSoRollback ，查看数据中数据的rollback情况

