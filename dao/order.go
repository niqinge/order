package dao

import (
	"github.com/niqinge/order/model"
	"gorm.io/gorm"
)

type OrderDAO interface {
	Create(req *model.Order) error
	UpdateByNo(orderNo string, m map[string]interface{}) error
	QueryByNo(no string) (order *model.Order, err error)
	QueryList(page, size int) (orders []*model.Order, err error)
}

type orderDAO struct {
	db *gorm.DB
}

func NewOrderDAO(db *gorm.DB) *orderDAO {
	return &orderDAO{db: db}
}

func (o *orderDAO) Create(req *model.Order) error {
	return o.db.Create(req).Error
}

func (o *orderDAO) UpdateByNo(orderNo string, m map[string]interface{}) error {
	return o.db.Model(&model.Order{}).Where("order_no=?", orderNo).Updates(m).Error
}

func (o *orderDAO) QueryByNo(no string) (order *model.Order, err error) {
	err = o.db.Where("order_no=?", no).Last(&order).Error
	return
}

func (o *orderDAO) QueryList(page, size int) (orders []*model.Order, err error) {
	err = o.db.Offset(page * size).Limit(size).Find(&orders).Error
	return
}
