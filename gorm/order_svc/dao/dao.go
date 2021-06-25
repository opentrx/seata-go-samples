package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/gogf/gf/util/gconv"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Dao struct {
	DB *gorm.DB
}

//现实中涉及金额可能使用长整形，这里使用 float64 仅作测试，不具有参考意义

type SoMaster struct {
	Sysno                int64   `json:"sysNo"`
	SoId                 string  `json:"soID"`
	BuyerUserSysno       int64   `json:"buyerUserSysNo"`
	SellerCompanyCode    string  `json:"sellerCompanyCode"`
	ReceiveDivisionSysno int64   `json:"receiveDivisionSysNo"`
	ReceiveAddress       string  `json:"receiveAddress"`
	ReceiveZip           string  `json:"receiveZip"`
	ReceiveContact       string  `json:"receiveContact"`
	ReceiveContactPhone  string  `json:"receiveContactPhone"`
	StockSysno           int64   `json:"stockSysNo"`
	PaymentType          int32   `json:"paymentType"`
	SoAmt                float64 `json:"soAmt"`
	//10，创建成功，待支付；30；支付成功，待发货；50；发货成功，待收货；70，确认收货，已完成；90，下单失败；100已作废
	Status       int32      `json:"status"`
	OrderDate    time.Time  `json:"orderDate"`
	PaymentDate  *time.Time `json:"paymentDate"`
	DeliveryDate *time.Time `json:"deliveryDate"`
	ReceiveDate  *time.Time `json:"receiveDate"`
	Appid        string     `json:"appID"`
	Memo         string     `json:"memo"`
	CreateUser   *string    `json:"createUser"`
	GmtCreate    *time.Time `json:"gmtCreate"`
	ModifyUser   *string    `json:"modifyUser"`
	GmtModified  *time.Time `json:"gmtModified"`

	SoItems []*SoItem `gorm:"-"`
}

type SoItem struct {
	Sysno         int64   `json:"sysNo"`
	SoSysno       int64   `json:"soSysNo"`
	ProductSysno  int64   `json:"productSysNo"`
	ProductName   string  `json:"productName"`
	CostPrice     float64 `json:"costPrice"`
	OriginalPrice float64 `json:"originalPrice"`
	DealPrice     float64 `json:"dealPrice"`
	Quantity      int32   `json:"quantity"`
}

func (dao *Dao) CreateSO(ctx context.Context, soMasters []*SoMaster) ([]uint64, error) {
	result := make([]uint64, 0, len(soMasters))
	tx := dao.DB.WithContext(ctx).Begin(&sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})

	for _, soMaster := range soMasters {
		soid := NextID()
		so_master := &SoMaster{
			Sysno:                gconv.Int64(soid),
			SoId:                 gconv.String(soid),
			BuyerUserSysno:       soMaster.BuyerUserSysno,
			SellerCompanyCode:    soMaster.SellerCompanyCode,
			ReceiveDivisionSysno: soMaster.ReceiveDivisionSysno,
			ReceiveAddress:       soMaster.ReceiveAddress,
			ReceiveZip:           soMaster.ReceiveZip,
			ReceiveContact:       soMaster.ReceiveContact,
			ReceiveContactPhone:  soMaster.ReceiveContactPhone,
			StockSysno:           soMaster.StockSysno,
			PaymentType:          soMaster.PaymentType,
			SoAmt:                soMaster.SoAmt,
			Status:               soMaster.Status,
			OrderDate:            time.Now(),
			Appid:                soMaster.Appid,
			Memo:                 soMaster.Memo,
		}

		if err := tx.Omit(clause.Associations).Create(so_master).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		soItems := soMaster.SoItems
		for _, soItem := range soItems {
			soItemID := NextID()
			so_item := &SoItem{
				Sysno:         gconv.Int64(soItemID),
				SoSysno:       gconv.Int64(soid),
				ProductSysno:  soItem.ProductSysno,
				ProductName:   soItem.ProductName,
				CostPrice:     soItem.CostPrice,
				OriginalPrice: soItem.OriginalPrice,
				DealPrice:     soItem.DealPrice,
				Quantity:      soItem.Quantity,
			}
			if err := tx.Create(so_item).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}
		result = append(result, soid)
	}
	err := tx.Commit().Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NextID() uint64 {
	id, _ := uuid.NewUUID()
	return uint64(id.ID())
}
