package main

import (
	"fmt"
	"gorm.io/driver/mysql"
)
import "gorm.io/gorm"

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	ProductID int
	Number    string
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})
	//
	//category := Category{Name: "Eletrônicos"}
	//db.Create(&category)
	//
	//p := Product{Name: "Produto 1", Price: 10.5, CategoryID: category.ID}
	//db.Create(&p)
	//
	//serialNumber := SerialNumber{Number: "123456", ProductID: p.ID}
	//db.Create(&serialNumber)

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)

	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
}
