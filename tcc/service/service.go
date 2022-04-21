package service

import (
	"context"
	"fmt"

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

	businessActionContextB := &ctx.BusinessActionContext{
		RootContext:   rootContext,
		ActionContext: make(map[string]interface{}),
	}
	businessActionContextB.ActionContext["hello"] = "hello world,this is from BusinessActionContext B"

	resultA, err := TccProxyServiceA.Try(businessActionContextA, false)
	fmt.Printf("result A is :%v", resultA)
	if err != nil {
		return err
	}

	resultB, err := TccProxyServiceB.Try(businessActionContextB, false)
	fmt.Printf("result B is :%v", resultB)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) TCCCanceled(context context.Context) error {
	rootContext := context.(*ctx.RootContext)
	businessActionContextA := &ctx.BusinessActionContext{
		RootContext:   rootContext,
		ActionContext: make(map[string]interface{}),
	}
	businessActionContextA.ActionContext["hello"] = "hello world,this is from BusinessActionContext A"

	businessActionContextC := &ctx.BusinessActionContext{
		RootContext:   rootContext,
		ActionContext: make(map[string]interface{}),
	}
	businessActionContextC.ActionContext["hello"] = "hello world,this is from BusinessActionContext C"

	resultA, err := TccProxyServiceA.Try(businessActionContextA, false)
	fmt.Printf("result A is :%v", resultA)
	if err != nil {
		return err
	}

	resultC, err := TccProxyServiceC.Try(businessActionContextC, false)
	fmt.Printf("result C is :%v", resultC)
	if err != nil {
		return err
	}

	return nil
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
