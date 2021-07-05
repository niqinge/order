package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/niqinge/order/model"
	smodel "github.com/niqinge/order/server/server_model"
	"github.com/niqinge/order/utils"
	"net/http"
	"time"
)

func (s *Server) QueryOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		no := c.Param("order_no")
		if len(no) == 0 {
			s.httpErr(c, http.StatusBadRequest, "Not found no.")
			return
		}
		
		o, err := s.order.QueryByNo(no)
		if err != nil {
			s.httpErr(c, http.StatusBadGateway, err.Error())
			return
		}

		s.httpSuccess(c, &smodel.Order{
			ID:       o.ID,
			OrderNo:  o.OrderNo,
			UserName: o.UserName,
			Amount:   o.Amount,
			FileUrl:  o.FileUrl,
		})
	}
}

func (s *Server) QueryOrderList() gin.HandlerFunc {
	return func(c *gin.Context) {
		orders, err := s.order.QueryList(0, 1000)
		if err != nil {
			s.httpErr(c, http.StatusBadGateway, err.Error())
			return
		}
		
		s.httpSuccess(c, orders)
	}
}

func (s *Server) AddOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req *smodel.AddOrderReq
		if err := c.BindJSON(&req); err != nil {
			s.httpErr(c, http.StatusBadRequest, err.Error())
			return
		}
		
		if err := req.IsValid(); err != nil {
			s.httpErr(c, http.StatusBadRequest, err.Error())
			return
		}
		
		o := model.Order{
			OrderNo:  fmt.Sprintf("%s%d", time.Now().Format(utils.FormatDateTime), time.Now().UnixNano()),
			UserName: req.UserName,
			Amount:   req.Amount,
			Status:   true,
			FileUrl:  req.FileUrl,
		}
		
		if err := s.order.Create(&o); err != nil {
			s.httpErr(c, http.StatusBadGateway, err.Error())
			return
		}
		
		s.httpSuccess(c, &smodel.Order{
			ID:       o.ID,
			OrderNo:  o.OrderNo,
			UserName: o.UserName,
			Amount:   o.Amount,
			FileUrl:  o.FileUrl,
		})
	}
}

func (s *Server) UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req *smodel.Order
		if err := c.BindJSON(&req); err != nil {
			s.httpErr(c, http.StatusBadRequest, err.Error())
			return
		}
		
		if err := req.IsValid(); err != nil {
			s.httpErr(c, http.StatusBadRequest, err.Error())
			return
		}
		
		if err := s.order.UpdateByNo(req.OrderNo, map[string]interface{}{
			"amount":req.Amount,
		}); err != nil {
			s.httpErr(c, http.StatusBadGateway, err.Error())
			return
		}
		s.httpSuccess(c, nil)
	}
}

func (s *Server) Upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		no := c.PostForm("order_no")
		if len(no) == 0 {
			s.httpErr(c, http.StatusBadRequest, "Not found no.")
			return
		}
		
		mt, err := c.FormFile("file")
		if err != nil {
			s.httpErr(c, http.StatusBadRequest, err.Error())
			return
		}
		
		path := utils.FilePath(mt.Filename)
		if err = c.SaveUploadedFile(mt, path); err != nil {
			s.httpErr(c, http.StatusBadGateway, err.Error())
			return
		}
		
		if err = s.order.UpdateByNo(no, map[string]interface{}{
			"file_url":path,
		}); err != nil {
			s.httpErr(c, http.StatusBadGateway, err.Error())
			return
		}
		
		s.httpSuccess(c, path)
	}
}

func (s *Server) httpErr(c *gin.Context, code int, err string) {
	c.JSON(code, utils.NewHttpErrResp(err))
}

func (s *Server) httpSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, utils.NewHttpSuccessResp(data))
}
