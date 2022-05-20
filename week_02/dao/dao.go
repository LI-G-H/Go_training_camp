package dao

import (
	"database/sql"
	// 需要导入，并且准备init连接数据库的方法被其调用
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var Db *sql.DB

// Student 学生实体类
type Student struct {
	id   uint16
	name string
	age  int8
}

// init 初始化数据库链接
func init() {
	var err error
	Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/local_test")
	if err != nil {
		panic(err)
	}
}

// SelectStudentById 查询方法
func SelectStudentById(id uint16) (Student, error) {
	var student Student
	row := Db.QueryRow("select id, name, age from student where id = ?;", id)
	err := row.Scan(&student.id, &student.name, &student.age)
	if err != nil {
		// 将错误包装起来, 并加入错误说明
		return student, errors.Wrap(err, "dao# SelectStudentById err")
	}
	return student, nil
}
