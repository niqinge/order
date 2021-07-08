package dao

import (
	"fmt"
	"github.com/niqinge/order/dao/db"
	"github.com/niqinge/order/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type OrderTestSuite struct {
	suite.Suite
	order *model.Order
	dao OrderDAO
}

func (s *OrderTestSuite) SetupSuite() {
	s.T().Log("SetupSuite")
	s.dao = NewOrderDAO(db.NewOrm())
	
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
	require.NoError(s.T(), s.dao.Create(s.order))
}

func (s *OrderTestSuite) TestQueryByNo() {
	s.T().Log("TestQueryByNo")
	o, err := s.dao.QueryByNo(s.order.OrderNo)
	require.NoError(s.T(), err, "条件不成立时输出")
	require.NotNil(s.T(), o, "条件不成立时输出")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestOrderTestSuite(t *testing.T) {
	suite.Run(t, new(OrderTestSuite))
}

