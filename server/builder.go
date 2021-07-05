package server

import (
	"fmt"
	"github.com/niqinge/order/dao"
	"github.com/niqinge/order/dao/db"
	"github.com/niqinge/order/service"
	"gorm.io/gorm"
)

func NewServer(dst ...interface{}) *Server {
	orm := db.NewOrm()
	// create(orm, dst...)
	migrate(orm, dst...)
	
	orderDao := dao.NewOrderDAO(orm)

	orderService := service.NewOrderService(orderDao)

	return &Server{
		order: orderService,
	}
}

func create(db *gorm.DB, dst ...interface{}) {
	if err := db.Migrator().CreateTable(dst); err != nil {
		panic(err)
	}
	fmt.Println("hhhh")
}

func migrate(db *gorm.DB, dst ...interface{}) {
	fmt.Println("migrate", len(dst))
	if err := db.AutoMigrate(dst...); err != nil {
		panic(err)
	}
	fmt.Println("ffff")
}
