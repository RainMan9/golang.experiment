package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleRequest)
	http.ListenAndServe(":8888", nil)
}
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	var strProductName string
	if r.Method == "POST" {
		//获取用户输入的产品名称
		r.ParseForm()
		strProductName = r.FormValue("txtProductName")
	}
	//查询数据
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/db2?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	var rows *sql.Rows
	if strProductName == "" {
		rows, err = db.Query("select ProductName,Price from Product")
	} else {
		strProductName = "%" + strProductName + "%"
		rows, err = db.Query("select ProductName,Price from Product where ProductName like ?", strProductName)
	}
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	//展示数据
	strHTML := "<html><header><title>产品查询</title></header><body>"
	strHTML += "<form action=\"/\" method=\"post\" >产品名称:"
	strHTML += "<input id=\"textProductName\" name=\"txtProductName\" type=\"text\" />"
	strHTML += "<input type=\"submit\" value=\"搜索\" /></form>"
	strHTML += "<table style='width:500px' border='1' cellpadding='2' cellspacing='2' > <tr>"
	strHTML += "<th align='left'>产品名称</th><th align='left' >价格</th></tr>"

	for rows.Next() {
		var name string
		var price float64
		rows.Scan(&name, &price)
		strHTML += "<tr><td>" + html.EscapeString(name) + "</td><td>" + fmt.Sprintf("%v", price) + "</td></tr>"

	}

	rows.Close()
	strHTML += "</table></body></html>"
	w.Write([]byte(strHTML))
}
