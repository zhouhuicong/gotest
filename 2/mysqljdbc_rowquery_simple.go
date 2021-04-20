package main

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"strconv"

	_ "github.com/go-sql-driver/mysql" //就是你下载的文件地址，如果是自己拷贝的，那么就写自己创建的路径

	xerrors "github.com/pkg/errors"
)

//连接到mysql
func sourceName2(userName, passWord, ip, port, database string) string {
	var connection = userName + ":" + passWord + "@tcp(" + ip + ":" + port + ")/" + database + "?charset=utf8"
	return connection
}

func query(dataSourceName, selfsql string) (string, error) {
	//连接示例
	// db,err := sql.Open("mysql","test2:abc@tcp(192.168.136.136:3306)/test?charset=utf8" )
	conn, err := sql.Open("mysql", dataSourceName)
	if err != nil {

		// log.Println(err)
		xerrors.Wrap(err, "connection is failed")
		log.Println(err)
		panic(err.Error())
	} else {
		fmt.Println("connection mysql succcess ! ")
	}
	defer conn.Close() //只有在前面用了 panic[抛出异常] 这时defer才能起作用，如果链接数据的时候出问题，他会往err写数据。defer:延迟，这里立刻申请了一个关闭sql 链接的草错，defer 后的方法，或延迟执行。在函数抛出异常一会被执行
	funcName, file, line, ok := runtime.Caller(0)
	var temp string
	if ok {
		// fmt.Println("func name: " + runtime.FuncForPC(funcName).Name())
		// fmt.Printf("file: %s, line: %d\n", file, line)
		temp = "file: " + file + " func name: " + runtime.FuncForPC(funcName).Name() + " line : " + strconv.Itoa(line)
	}

	//产生查询语句的Statement
	var name string
	err = conn.QueryRow(selfsql).Scan(&name)

	switch {
	case err == sql.ErrNoRows:
		log.Println(err)
		return "", xerrors.Wrap(err, temp)
		//fmt.Println("sql.ErrNoRows:", err)

	case err != nil:
		fmt.Println(err)
	}
	return name, err

}

func main() {
	//username mysql账号
	var userName = "root"
	//password mysql密码
	var passWord = "123456"
	//ip   mysql数据库的IP
	var ip = "localhost"
	//port  mysql数据库的端口
	var port = "3306"
	// database 需要连接的数据库名称
	var database = "test"

	dataSourceName := sourceName2(userName, passWord, ip, port, database)

	nam, err := query(dataSourceName, "select * from test")
	if err != nil {
		fmt.Println("err is :", err)
	}
	fmt.Println("name is : ", nam)
}
