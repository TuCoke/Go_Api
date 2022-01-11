package dbops

import (
	"database/sql"
)

var (
	dbConn *sql.DB
	err error
)

func init(){
	_,err:=sql.Open("mysql","goApi:TzjfH3Ay4hSieJDW@tcp(:3306)/goapi?charset=utf8mb4")
	if err!=nil{
		panic(err.Error()) // 中断
	}
	//fmt.Println(dbConn)
}
