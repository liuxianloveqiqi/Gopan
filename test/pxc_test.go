package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestName(t *testing.T) {

	// 数据库连接参数
	dbHost := "43.139.195.17"
	dbPort := 6446
	dbUser := "root"
	dbPassword := "root"
	dbName := "gopan"

	// 构建连接字符串
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(dataSourceName)
	// 连接到数据库
	db, err := sql.Open("mysql", "root:root@(43.139.195.17:6446)/gopan?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to ping database:", err)
		return
	}

	fmt.Println("Connected to database successfully!")

	// 进行数据库操作...
}
