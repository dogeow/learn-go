package learn

import (
	"fmt"
	"github.com/likunyan/learn-go/api/models"
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

	user := models.User{Name: "233", Email: "233@qq.com", Password: "666"}

	result := db.Create(&user) // 通过数据的指针来创建
	fmt.Println(result)
}
