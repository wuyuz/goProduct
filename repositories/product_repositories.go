package repositories

import (
	"database/sql"
	"fmt"
	"strconv"
	"test-produce/common"
	"test-produce/datamodels"
)

// step1: 先开发我们的接口
type IProduct interface {
	// 连接数据库,判断是否成功
	Conn() error // 首先这是接口 method (参数列表) (返回值列表),这里表示传入() 返回error
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error) // 增删改查,这里查以主键查
	SelectAll() ([]*datamodels.Product, error)
}

// step2: 实现定义的接口
type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

// 结构体接口自检,默认是隐式实现的
func NewProductManager(table string, db *sql.DB) IProduct {
	// 此函数如果返回的ProductManager不符合接口是会报错的,达到自检的效果
	return &ProductManager{table: table, mysqlConn: db}
}

// 数据库连接
func (p *ProductManager) Conn() (err error) {
	// 确保是否有数据库连接
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}

	if p.table == "" {
		p.table = "product"
	}
	return
}

// 插入函数
func (p *ProductManager) Insert(product *datamodels.Product) (productId int64, err error) {

	// 判断连接是否存在
	if err = p.Conn(); err == nil {
		return
	}
	// 准备sql
	sql := "INSERT product SET productName=?,productNum=?,productImage=?,productUrl=?"
	// mysql的sql预编译，之后就可以执行Exec Query 等方法
	stmt, errSql := p.mysqlConn.Prepare(sql)
	if errSql != nil {
		return 0, errSql
	}

	// 传入参数
	result, errStm := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if errStm != nil {
		return 0, errStm
	}

	return result.LastInsertId()
}

// 删除函数
func (p *ProductManager) Delete(productID int64) bool {
	// 判断连接是否存在
	if err := p.Conn(); err == nil {
		return false
	}

	sql := "DELETE FROM product where id=?"
	stmt, errSql := p.mysqlConn.Prepare(sql)
	if errSql == nil {
		return false
	}

	// 执行删除
	_, errStmt := stmt.Exec(productID)
	if errStmt != nil {
		return false
	}
	return true
}

// 更新函数
func (p *ProductManager) Update(product *datamodels.Product) error {
	// 判断连接是否存在
	if err := p.Conn(); err == nil {
		return err
	}

	sql := "UPDATE product set productName=?,productNum=?,productImage=?, productUrl=? where id=" + strconv.FormatInt(product.ID, 10)

	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return err
	}
	return nil
}

// 根据商品id查询商品
func (p *ProductManager) SelectByKey(productID int64) (productResult *datamodels.Product, err error) {
	// 判断连接是否存在
	if err = p.Conn(); err == nil {
		return &datamodels.Product{}, err
	}

	sql := "SELETE * FROM " + p.table + "WHERE ID=" + strconv.FormatInt(productID, 10)

	row, errRow := p.mysqlConn.Query(sql)
	if errRow != nil {
		return &datamodels.Product{}, errRow
	}

	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Product{}, nil
	}

	common.DataToStructByTagSql(result, productResult)

	return
}

//
func (p *ProductManager) SelectAll() (productArray []*datamodels.Product, err error) {

	// 判断连接是否存在
	if err = p.Conn(); err != nil {
		fmt.Println("all: 连接异常")
		return nil , err
	}
	fmt.Println("all: step2")
	sql := "Select * from "+p.table
	rows,err := p.mysqlConn.Query(sql)
	defer  rows.Close()
	if err !=nil {
		return nil ,err
	}

	result:= common.GetResultRows(rows)
	if len(result)==0{
		return nil,nil
	}

	for _,v :=range result{
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v,product)
		productArray=append(productArray, product)
	}
	return
}
