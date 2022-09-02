package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

//DB 定义全局变量mysql连接且初始化完成,项目中可使用DB
var DB *sql.DB = SetupDatabaseConnectionDB()

func SetupDatabaseConnectionDB() *sql.DB {
	//加载.env配置文件，读取配置数据库连接信息
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Unable to load .env file in SetupDatabaseConnection func!")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True", dbUser, dbPass, dbHost, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("Failed to create connection database!")
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	pErr := db.Ping()
	if pErr != nil {
		fmt.Printf("database DB object is ping to failed,error:%v", pErr)
		return nil
	}
	return db
}

func CloseDatabaseConnection(db *sql.DB) {
	db.Close()
}
