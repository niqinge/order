package service

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/niqinge/order/mock"
	"github.com/niqinge/order/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
	"time"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type OrderTestSuite struct {
	suite.Suite
	order        *model.Order
	dao          *mock.MockOrderDAO
	orderService *OrderService
}

func (s *OrderTestSuite) SetupSuite() {
	s.T().Log("SetupSuite")
	
	s.dao = mock.NewMockOrderDAO(gomock.NewController(s.T()))
	s.orderService = NewOrderService(s.dao)

	s.order = &model.Order{
		OrderNo:  fmt.Sprintf("%d", time.Now().UnixNano()),
		UserName: "pual",
		Amount:   1000,
		Status:   false,
		FileUrl:  "FileUrl",
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (s *OrderTestSuite) SetupTest() {
	s.T().Log("SetupTest")
}

func (s *OrderTestSuite) TestCreateOrder() {
	s.T().Log("TestCreateOrder")
	createF := s.dao.EXPECT().Create(gomock.Any()).AnyTimes()
	s.Run("CreateOrder", func() {
		createF.Return(nil)
		require.NoError(s.T(), s.orderService.Create(s.order), "CreateOrderErr##error is not nil")
	} )
	
	s.Run("CreateOrderErr", func() {
		createF.Return(errors.New("create failed. "))
		require.Error(s.T(), s.orderService.Create(s.order), "CreateOrderErr##error is nil")
		// s.T().Error(err)
	})
}

func (s *OrderTestSuite) TestQueryByNo() {
	s.T().Log("TestQueryByNo")
	queryByNoF := s.dao.EXPECT().QueryByNo(gomock.Any()).AnyTimes()
	
	s.Run("QueryByNo", func() {
		queryByNoF.Return(&model.Order{
			Model:    gorm.Model{
				ID:        1234,
				CreatedAt: time.Now(),
			},
			OrderNo:  "testmock",
			UserName: "mock",
			Amount:   1000,
			Status:   true,
			FileUrl:  "./url",
		}, nil)
		o, err := s.orderService.QueryByNo(s.order.OrderNo)
		require.NoError(s.T(), err, "QueryByNo##error is not nil")
		require.NotNil(s.T(), o, "QueryByNo##order is nil")
		// s.T().Logf("QueryByNo:%+v", o)
	})
	
	s.Run("QueryByNoErr", func() {
		queryByNoF.Return(nil, errors.New("QueryByNo failed"))
		o, err := s.orderService.QueryByNo(s.order.OrderNo)
		require.Error(s.T(), err, "QueryByNoErr##error is nil")
		require.Nil(s.T(), o, "QueryByNoErr##Order is not nil")
	})
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestOrderTestSuite(t *testing.T) {
	suite.Run(t, new(OrderTestSuite))
}
