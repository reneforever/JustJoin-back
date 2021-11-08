package main

import (
	"database/sql"
	"fmt"

	db "JustJoin.com/DB"
	zap "JustJoin.com/Log"
	user "JustJoin.com/User"
	routers "JustJoin.com/routers"
	"github.com/gin-gonic/gin"
)

var currDB *sql.DB

func main() {
	zapInit := zap.InitLogger()
	fmt.Println(zapInit)
	dbInitErr := db.InitDB(currDB)
	if dbInitErr != nil {
		fmt.Println("db init failed, err:", dbInitErr)
		return
	}
	r := gin.Default()
	// 划分路由
	routers.UserLog(r)
	userInfo, queryErr := user.QueryUserInfo(currDB)
	r.Run(":8080")
}
