package learn

import (
	"fmt"
	"github.com/dogeow/learn-go/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Db() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DBNAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	user := models.User{}

	//result := db.Create(&user) // 通过数据的指针来创建
	//result := db.First(&user)]
	//var user models.User
	//result := map[string]interface{}{}
	//result2 := db.Model(&user).First(&result)
	//fmt.Println(result2)

	// 新建会话模式
	stmt := db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement
	fmt.Println(stmt.SQL.String())
	fmt.Println(db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...))

	//result := db.Delete(&user, 2)
	// DELETE FROM users WHERE id = 10;
	// 带额外条件的删除
	//result := db.Where("name = ?", "233").Delete(&user)
	//// DELETE from emails where id = 10 AND name = "jinzhu";

	//fmt.Println(result.RowsAffected)
}
