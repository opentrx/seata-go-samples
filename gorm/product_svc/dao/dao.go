package dao

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Inventory struct {
	Sysno           uint64
	ProductSysno    uint64
	AccountQty      int32
	AvailableQty    int32
	AllocatedQty    int32
	AdjustLockedQty int32
}

type Dao struct {
	DB *gorm.DB
}

type AllocateInventoryReq struct {
	ProductSysNo int64
	Qty          int32
}

func (dao *Dao) AllocateInventory(ctx context.Context, reqs []*AllocateInventoryReq) error {
	tx := dao.DB.WithContext(ctx).Begin(&sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})

	for _, req := range reqs {
		if err := tx.Model(&Inventory{}).
			Where("product_sysno = ? and available_qty >= ?", req.ProductSysNo, req.Qty).
			UpdateColumns(map[string]interface{}{
				"available_qty": gorm.Expr("available_qty - ?", req.Qty),
				"allocated_qty": gorm.Expr("allocated_qty + ?", req.Qty),
			}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	err := tx.Commit().Error
	if err != nil {
		return err
	}
	return nil
}
