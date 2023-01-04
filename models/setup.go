// untuk membuat koneksi ke database kita
package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	// dsn := "host=localhost user=postgres password=1234 dbname=rest-api-jwt port=5432 sslmode=disable TimeZone=Asia/jakarta"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	fmt.Println("Database tidak terhubug")
	// }

	// // untuk auto migrate ke database postgres
	// db.AutoMigrate(&User{})

	// DB = db

	// Koneksi Ke Database PostgresSQL
	dsn := "host=localhost user=postgres password=1234 dbname=rest-api-jwt port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Tidak terkoneksi ke database")
	}
	// fmt.Println("database sudah terkoneksi")

	db.AutoMigrate(&User{})

	DB = db

}
