package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "strings"
)

var (
	name     = "root"
	password = "123456"
	host     = "localhost"
	port     = "3306"
	dbname   = "test"
)

type User struct {
	id       int
	username string
	password string
}

var DB *sql.DB

func IntoDB() {
	dsn := name + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	DB = db
}

func userLogin(c *gin.Context) {
	userName := c.Request.URL.Query().Get("username")
	passWord := c.Request.URL.Query().Get("password")
	rows, err := DB.Query("SELECT * FROM test")
	if err != nil {
		fmt.Println("查询失败")
	}
	var s User
	for rows.Next() {
		err = rows.Scan(&s.id, &s.username, &s.password)
		if err != nil {
			fmt.Println(err)
		}
	}
	if userName != s.username {
		c.JSON(200, gin.H{
			"success": false,
			"code":    400,
			"msg":     "无此用户",
		})
	} else {
		us, _ := DB.Query("SELECT password FROM test where username='" + userName + "'")
		for us.Next() {
			var u User
			err = us.Scan(&u.password)
			if err != nil {
				fmt.Println(err)
			}
			if passWord != u.password {
				c.JSON(200, gin.H{
					"success": false,
					"code":    400,
					"msg":     "密码错误",
				})
			} else {
				c.JSON(200, gin.H{
					"success": true,
					"code":    200,
					"msg":     "登录成功",
				})
			}
		}
	}
	rows.Close()
}
func userRegister(c *gin.Context) {
	userName := c.Request.URL.Query().Get("username")
	passWord := c.Request.URL.Query().Get("password")
	userQuestion := c.Request.URL.Query().Get("question")
	rows, err := DB.Query("SELECT * FROM test")
	if err != nil {
		fmt.Println("查询失败")
	}
	for rows.Next() {
		var s User
		err = rows.Scan(&s.id, &s.username, &s.password)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(s.username)
		if userName != s.username {
			result, err := DB.Exec("INSERT INTO test(username,password,question)VALUES (?,?,?)", userName, passWord, userQuestion)
			if err != nil {
				fmt.Println("执行失败")
				return
			} else {
				rows, _ := result.RowsAffected()
				if rows != 1 {
					c.JSON(200, gin.H{
						"success": false,
					})
				} else {
					c.JSON(200, gin.H{
						"success":  true,
						"username": userName,
						"question": userQuestion,
					})
				}
			}
		} else {
			fmt.Println("用户名已被注册")
			c.JSON(200, gin.H{
				"code":    400,
				"success": false,
				"msg":     "用户名已被注册",
			})
		}
	}
	rows.Close()
}
func main() {
	IntoDB()
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/login", userLogin)
		user.POST("/register", userRegister)
	}
	r.Run()
}
