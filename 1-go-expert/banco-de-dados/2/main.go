package main

import (
	"gorm.io/driver/mysql"
)
import "gorm.io/gorm"

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	//db.Create(&Product{Name: "Produto 1", Price: 10.5})

	//products := []Product{
	//	{Name: "Produto 2", Price: 20.5},
	//	{Name: "Produto 3", Price: 30.5},
	//}
	//
	//db.Create(&products)

	//var product Product
	//db.First(&product, 1)
	//fmt.Printf("ID: %d\n", product.ID)
	//fmt.Printf("Name: %s\n", product.Name)
	//fmt.Printf("Price: %.2f\n", product.Price)

	//var products []Product
	//db.Limit(2).Offset(2).Find(&products)
	//for _, product := range products {
	//	fmt.Printf("ID: %d\n", product.ID)
	//	fmt.Printf("Name: %s\n", product.Name)
	//	fmt.Printf("Price: %.2f\n", product.Price)
	//}

	//var products []Product
	//db.Where("price > ?", 20).Find(&products)
	//for _, product := range products {
	//	fmt.Printf("ID: %d\n", product.ID)
	//	fmt.Printf("Name: %s\n", product.Name)
	//	fmt.Printf("Price: %.2f\n", product.Price)
	//	fmt.Println("-----")
	//}

	var product Product
	db.First(&product, 1)
	product.Name = "New mouse 2"

	db.Save(&product)
	//
	//var p2 Product
	//db.First(&p2, 1)
	//fmt.Printf("ID: %d\n", p2.ID)
	//fmt.Printf("Name: %s\n", p2.Name)
	//fmt.Printf("Price: %.2f\n", p2.Price)
	//
	db.Delete(&product)

}
