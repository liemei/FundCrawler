package models


import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
	tables []interface{}
)

func init() {
	fmt.Println("开始执行models init")
	tables = append(tables, new(FundInfo)
	engine, _ = xorm.NewEngine("mysql", "root:19900108Lsc!!@(140.143.250.177:3306)/fund?charset=utf8")
	var sqlLog, err = os.Create("sql.log")
	if err != nil {
		fmt.Println(err.Error())
	}
	engine.SetLogger(xorm.NewSimpleLogger(sqlLog))
	engine.ShowSQL(true)
	engine.ShowExecTime(true)
	//创建表
	// for _, table := range tables {
	// 	exit, _ := engine.IsTableExist(table)
	// 	if !exit {
	// 		engine.CreateTables(table)
	// 	}
	// }
	//同步表和表结构
	fmt.Println("表开始同步")
	syncerr := engine.Sync2(new(FundInfo))
	if syncerr != nil {
		fmt.Println("表结构同步报错", syncerr.Error())
	}
}
