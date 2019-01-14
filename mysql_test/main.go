// mysql_test project main.go
package main

import (
	// "database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type NewUser struct {
	gorm.Model
	Name     string `gorm:"default:'galeone'"`
	Age      int64  `gorm:"default:18"`
	Birthday time.Time
}

func main() {
	dns := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=true", "root", "lwstar", "tcp", "localhost", 3306, "lw")
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		return
	}
	defer db.Close()

	// db.CreateTable(&user)

	for i := 0; i < 100; i++ {
		user := NewUser{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

		ok := db.NewRecord(user)
		if ok {

		}
	}

	// user := NewUser{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	// db.CreateTable(&user)
	// db.Create(&user)
	// ok := db.NewRecord(user)
	// ok = db.NewRecord(user)
	// if ok {

	// }
	// user1 := NewUser{}
	// db.First(&user)
	// fmt.Println(user1)

	fmt.Println("ok")
}
