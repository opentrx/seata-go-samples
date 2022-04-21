package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	ctx "github.com/opentrx/seata-golang/v2/pkg/client/base/context"
	"github.com/opentrx/seata-golang/v2/pkg/client/base/model"
)

type Service struct {
}

func (svc *Service) TCCCommitted(context context.Context) error {
	rootContext := context.(*ctx.RootContext)
	businessActionContextA := &ctx.BusinessActionContext{
		RootContext:   rootContext,
		ActionContext: make(map[string]interface{}),
	}
	// 业务参数全部放到 ActionContext 里
	businessActionContextA.ActionContext["hello"] = "hello world,this is from BusinessActionContext A"

	resultA, err := TccProxyServiceA.Try(businessActionContextA, false)
	fmt.Printf("result A is :%v", resultA)
	if err != nil {
		return err
	}

	req2, err := http.NewRequest("GET", "http://localhost:8082/try", nil)
	if err != nil {
		return err
	}
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("xid", rootContext.GetXID())

	client := &http.Client{}
	resultB, err2 := client.Do(req2)
	if err2 != nil {
		return err2
	}
	fmt.Printf("result B is :%v", resultB)

	return nil
}

func (svc *Service) TCCCanceled(context context.Context) error {
	rootContext := context.(*ctx.RootContext)
	businessActionContextA := &ctx.BusinessActionContext{
		RootContext:   rootContext,
		ActionContext: make(map[string]interface{}),
	}
	businessActionContextA.ActionContext["hello"] = "hello world,this is from BusinessActionContext A"

	resultA, err := TccProxyServiceA.Try(businessActionContextA, false)
	fmt.Printf("result A is :%v", resultA)
	if err != nil {
		return err
	}

	req2, err := http.NewRequest("GET", "http://localhost:8082/try", nil)
	if err != nil {
		return err
	}
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("xid", rootContext.GetXID())

	client := &http.Client{}
	resultB, err2 := client.Do(req2)
	if err2 != nil {
		return err2
	}
	fmt.Printf("result B is :%v", resultB)

	return errors.New("should cancel")
}

var service = &Service{}

type ProxyService struct {
	*Service
	TCCCommitted func(ctx context.Context) error
	TCCCanceled  func(ctx context.Context) error
}

func (svc *ProxyService) GetProxyService() interface{} {
	return svc.Service
}

func (svc *ProxyService) GetMethodTransactionInfo(methodName string) *model.TransactionInfo {
	return methodTransactionInfo[methodName]
}

var methodTransactionInfo = make(map[string]*model.TransactionInfo)

func init() {
	methodTransactionInfo["TCCCommitted"] = &model.TransactionInfo{
		TimeOut:     60000000,
		Name:        "TCC_TEST_COMMITTED",
		Propagation: model.Required,
	}
	methodTransactionInfo["TCCCanceled"] = &model.TransactionInfo{
		TimeOut:     60000000,
		Name:        "TCC_TEST_CANCELED",
		Propagation: model.Required,
	}
}

var ProxySvc = &ProxyService{
	Service: service,
}
