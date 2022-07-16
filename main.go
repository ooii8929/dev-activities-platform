package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)



type Provider interface {
    SessionInit(sid string) (Session, error)
    SessionRead(sid string) (Session, error)
    SessionDestroy(sid string) error
    SessionGC(maxLifeTime int64)
}

type Session interface {
    Set(key, value interface{}) error // set session value
    Get(key interface{}) interface{}  // get session value
    Delete(key interface{}) error     // delete session value
    SessionID() string                // back current sessionID
}


func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //解析參數，預設是不會解析的
    fmt.Println(r.Form)  //這些資訊是輸出到伺服器端的列印資訊
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") //這個寫入到 w 的是輸出到客戶端的
}

func login(w http.ResponseWriter, r *http.Request){
    fmt.Println("method", r.Method)
    if r.Method == "GET" {
        // 顯示登入畫面
        t, _ := template.ParseFiles("./static/login.gtpl")
        log.Println(t.Execute(w, nil))
    }else{
        // 請求為登入資料
        r.ParseForm()
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password", r.Form["password"])
    }
}


// 處理/upload 邏輯
func upload(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //取得請求的方法
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("upload.gtpl")
        t.Execute(w, token)
    } else {
        r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        fmt.Fprintf(w, "%v", handler.Header)
        f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)  // 此處假設當前目錄下已存在 test 目錄
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)
    }
}

func cookie(w http.ResponseWriter, r *http.Request) {
    expiration := time.Now()
    expiration = expiration.AddDate(1, 0, 0)
    cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
    http.SetCookie(w, &cookie)
}

func main() {
    http.HandleFunc("/", sayhelloName) //設定存取的路由
    http.HandleFunc("/login", login)
    http.HandleFunc("/upload", upload)
    http.HandleFunc("/cookie", cookie)

    db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:33360)/test")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    checkErr(err)

    //插入資料
    stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
    checkErr(err)

    res, err := stmt.Exec("astaxie", "研發部門", "2012-12-09")
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)
    //更新資料
    stmt, err = db.Prepare("update userinfo set username=? where uid=?")
    checkErr(err)

    res, err = stmt.Exec("astaxieupdate", id)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    //查詢資料
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }

    //刪除資料
    stmt, err = db.Prepare("delete from userinfo where uid=?")
    checkErr(err)

    res, err = stmt.Exec(id)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    db.Close()

    newErr := http.ListenAndServe(":9090", nil) //設定監聽的埠
    if newErr != nil {
        log.Fatal("ListenAndServe: ", err)
    }


}



func checkErr(err error) {
    if err != nil {
        
        panic(err)
    }
}