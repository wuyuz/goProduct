package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/imooc?charset=utf8")
	defer db.Close()
	// Query 和 Scan的简单使用
	rows, err := db.Query("SELECT productName FROM product WHERE id > ?", 0)
	if err != nil {
		log.Fatal(err)
	}

	// rows.Next() 类似于迭代器，返回每条记录, 返回多行数据
	for rows.Next() {
		var productName string
		// 通过Query得到的rows, 需要我们设置一个变量进行接受，通过rows.Scan(指针),将值传递给变量
		if err := rows.Scan(&productName); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("多行数据： %s\n ", productName)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// QueryRaw 用于返回单行的查询, 如果是*的话，需要使用多个参数进行接受
	//rows1 := db.QueryRow("SELECT * FROM product WHERE id = ?", 1)
	rows1 := db.QueryRow("SELECT productNum FROM product WHERE id = ?", 1)
	var productNum int64
	if err := rows1.Scan(&productNum); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("单行数据: %d\n", productNum)

	// 预占位符,db.Prepare()返回一个Stmt。Stmt对象可以执行Exec,Query,QueryRow等操作。
	stmt, _ := db.Prepare("insert into product (productName, productNum, productImage, productUrl) values (?,?,?,?)")

	// Exec() 执行函数，常用作执行增删该等操作
	row2, err := stmt.Exec(
		"娃哈哈",
		5,
		"/ha.jpg",
		"/ha",
	)
	idNew, errId := row2.LastInsertId()
	if errId != nil {
		fmt.Println(errId)
	}
	fmt.Printf("数据插入成功！返回操作的数据的ID： %d", idNew)
}
