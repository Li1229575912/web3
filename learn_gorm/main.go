package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	log.Println("数据库连接成功！")
	// 后面我们会在这里操作数据库
	db.AutoMigrate(&User{})
	log.Println("创建表成功！")
	email := "zhangsan@example.com"
	users := []User{
		{Name: "zhangsan", Email: &email, Age: 20, MemberNumber: sql.NullString{String: "123456", Valid: true}},
		{Name: "lisi", Email: &email, Age: 21, MemberNumber: sql.NullString{String: "654321", Valid: true}},
		{Name: "wangwu", Email: &email, Age: 22, MemberNumber: sql.NullString{String: "111111", Valid: true}},
	}
	result := db.Create(&users)
	log.Println(result.Error)
	log.Println(result.RowsAffected)

	var selectuser User
	db.First(&selectuser, 1) // 查主键为1的记录
	fmt.Println(selectuser.Name)

	var users1 []User
	db.Find(&users1) // 查所有记录
	fmt.Println(users1)
	// db.Model(&User{}).Where("id = ?", 11).Updates(User{
	// 	Name:  "新名字",
	// 	Email: "xxx@example.com",
	// })
	//db.Model(&user).Update("Email", "new@example.com")
	//log.Println("更新数据成功！")

	//db.Where("email = ?", "zhangsan@example.com").Delete(&User{})
	log.Println("删除数据成功！")

	var users2 []User
	db.Where("name = ?", "张三").Or("email = ?", "li@example.com").Find(&users2)
	//fmt.Println(users2)

	var count int64
	db.Model(&users2).Where("email LIKE ?", "%@example.com").Count(&count)
	fmt.Println(count)

}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Name = strings.ToUpper(u.Name)
	return
}
