// many to many
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
)
import "gorm.io/gorm"

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	//category := Category{Name: "Cozinha"}
	//db.Create(&category)
	//
	//category2 := Category{Name: "Eletrônicos"}
	//db.Create(&category2)
	//
	//p := Product{Name: "Produto 1", Price: 10.5, Categories: []Category{category, category2}}
	//db.Create(&p)

	var categories []Category
	db.Preload("Products").Find(&categories) //has many com has one

	for _, category := range categories {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			fmt.Println(product.Name)
		}
	}
}
