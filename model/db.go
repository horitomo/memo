package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 引っ張ってきたデータ受け取り
type Memo struct {
	ID   string
	Title string
	Content string
}

func DBInit(){
	db, err := sql.Open("mysql", "flow:flow@/gwa")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	sql := "CREATE TABLE memo( id varchar(255),title varchar(255),content varchar(255) )"
	db.Query(sql)
}

func DBInsert(id ,password string){
	//mysqlへ接続。ドライバ名（mysql）と、ユーザー名・データソース(ここではgosample)を指定。
	db, err := sql.Open("mysql", "flow:flow@/gwa")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// insert
	ins, err := db.Prepare("INSERT INTO test VALUES(?,?)")
	if err != nil {
		fmt.Println(err)
	}
	ins.Exec(id, password)
}

func MemoSelect(id string)(string){
	//mysqlへ接続。ドライバ名（mysql）と、ユーザー名・データソース(ここではgosample)を指定。
	db, err := sql.Open("mysql", "flow:flow@/gwa")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	rows, err := db.Query("SELECT * FROM memo WHERE id = ?",id)
	flag := false
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	str := ""
	var memo Memo //構造体Person型の変数personを定義
	for rows.Next() {
		err := rows.Scan(&memo.ID, &memo.Title, &memo.Content)			
		if err != nil {
			panic(err.Error())
		}
		str = str + memo.Title + memo.Content
		flag = true
	}
	if !flag {
		fmt.Println("失敗")
		str = "まだメモはありません"
	}
	return str
}

func MemoInsert(id ,title,content string){
	//mysqlへ接続。ドライバ名（mysql）と、ユーザー名・データソース(ここではgosample)を指定。
	db, err := sql.Open("mysql", "flow:flow@/gwa")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// insert
	ins, err := db.Prepare("INSERT INTO memo VALUES(?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	ins.Exec(id, title, content)
}
