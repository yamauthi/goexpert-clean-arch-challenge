package database

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/yamauthi/goexpert-clean-arch-challenge/internal/entity"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Exec("DELETE FROM orders")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestListOrders_WhenList_ThenShouldListAllOrders() {
	repo := NewOrderRepository(suite.Db)

	for i := 1; i <= 8; i++ {
		//Generate orders to insert
		order, err := entity.NewOrder("ORDER_ID-"+strconv.Itoa(i), 80.0, 20.0)
		suite.NoError(err)
		suite.NoError(order.CalculateFinalPrice())
		err = repo.Save(order)
		suite.NoError(err)
	}

	resultList, err := repo.List()
	suite.NoError(err)

	i := 1
	for _, order := range resultList {
		suite.Equal(order.ID, "ORDER_ID-"+strconv.Itoa(i))
		suite.Equal(order.Price, 80.0)
		suite.Equal(order.Tax, 20.0)
		suite.Equal(order.FinalPrice, 100.0)
		i++
	}

}
