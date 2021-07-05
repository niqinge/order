package service

import (
	"github.com/niqinge/order/dao"
	"github.com/niqinge/order/model"
)

type OrderService struct {
	orderDao dao.OrderDAO
}

func NewOrderService(orderDao dao.OrderDAO) *OrderService {
	return &OrderService{orderDao: orderDao}
}

func (o *OrderService) Create(req *model.Order) error {
	return o.orderDao.Create(req)
}

func (o *OrderService) UpdateByNo(orderNo string, m map[string]interface{}) error {
	return o.orderDao.UpdateByNo(orderNo, m)
}

func (o *OrderService) QueryByNo(no string) (order *model.Order, err error) {
	return o.orderDao.QueryByNo(no)
}

func (o *OrderService) QueryList(page, size int) (orders []*model.Order, err error) {
	return o.orderDao.QueryList(page, size)
}



