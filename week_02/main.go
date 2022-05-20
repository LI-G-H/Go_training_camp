package main

import (
	"database/sql"
	"fmt"
	"week_02/dao"
)

// 错误堆栈打印在main方法中
func main() {
	defer func(Db *sql.DB) {
		err := Db.Close()
		if err != nil {
			panic(err)
		}
	}(dao.Db)

	student, err := dao.SelectStudentById(2)
	if err != nil {
		fmt.Printf("select student err: %+v", err)
		return
	}
	fmt.Println("student  :  ", student)
}
