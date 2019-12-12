package main

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/horitomo/memo/model"
)

// 引っ張ってきたデータ受け取り
type Person struct {
	ID   string
	Password string
}

func main(){
	var userId string
	model.DBInit()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/",func(ctx *gin.Context){
		userId = ""
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.POST("/login",func(ctx *gin.Context){

		ctx.Request.ParseForm()
		id := ctx.Request.Form["id"][0]
		password := ctx.Request.Form["password"][0]

		//mysqlへ接続。ドライバ名（mysql）と、ユーザー名・データソース(ここではgosample)を指定。
		db, err := sql.Open("mysql", "flow:flow@/gwa")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
		rows, err := db.Query("SELECT * FROM test WHERE id = ? AND password = ?",id,password)
		flag := false
		defer rows.Close()
		if err != nil {
			panic(err.Error())
		}
		var person Person //構造体Person型の変数personを定義
		for rows.Next() {
			err := rows.Scan(&person.ID, &person.Password)			
			if err != nil {
				panic(err.Error())
			}
			flag = true
		}
		if !flag {
			ctx.HTML(200, "login.html", gin.H{"data2" : person.ID,"data3" : person.Password,"flag" : flag})
		}else {
			userId = person.ID
			ctx.Redirect(302,"/list")
		}
	})
	// 新規登録
	router.GET("/register",func(ctx *gin.Context){
		ctx.HTML(200, "register.html", gin.H{})
	})
	router.POST("/register",func(ctx *gin.Context){

		// postの値を受け取り
		ctx.Request.ParseForm()
		id := ctx.Request.Form["id"][0]
		password := ctx.Request.Form["password"][0]
		model.DBInsert(id,password)
		result := "完了"
		ctx.HTML(200, "register.html", gin.H{"result":result})
	})
	// memo
	router.GET("/list",func(ctx *gin.Context){
		judge := ""
		if userId != judge{
			memo := model.MemoSelect(userId)
			ctx.HTML(200,"list.html",gin.H{"memo":memo})
		}else {
			ctx.Redirect(302,"/")
		}
	})
	// memo登録
	router.GET("/list/register",func(ctx *gin.Context){
		judge := ""
		if userId != judge{
			ctx.HTML(200,"memoRegister.html",gin.H{})
		}else {
			ctx.Redirect(302,"/")
		}
	})
	router.POST("/list/register",func(ctx *gin.Context){
		// postの値を受け取り
		ctx.Request.ParseForm()
		id := userId
		title := ctx.Request.Form["title"][0]
		content := ctx.Request.Form["content"][0]
		model.MemoInsert(id,title,content)
		ctx.Redirect(302,"/list")
	})

	// logout
	router.GET("/logout",func(ctx *gin.Context){
		userId = ""
		ctx.Redirect(302,"/")
	})
	router.Run()
}
