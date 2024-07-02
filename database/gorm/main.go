package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

	// db.Create(&Product{
	// 	Name:  "Macbook Pro",
	// 	Price: 23000.00,
	// })

	// products := []Product{
	// 	{Name: "Macbook Pro", Price: 23000.50},
	// 	{Name: "Apple Pencil", Price: 1200.00},
	// 	{Name: "Apple Magic Mouse 2", Price: 800.00},
	// }
	// db.Create(&products)

	// SelectOne
	// var product Product
	// db.First(&product, "name = ?", "Apple Pencil")
	// fmt.Println(product)

	// Select ALl
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// Where
	// var products []Product
	// db.Where("price > ?", 10000).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// Like
	// var products []Product
	// db.Where("name LIKE ?", "%apple%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 1)
	// p.Name = "Samsung Galaxy Book"
	// db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	db.Delete(&p2)

}
