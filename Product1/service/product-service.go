package service

import (
	"fmt"

	"Product1/config"
	"Product1/model"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllAlbums - fetch all Albums
func FetchProductsFromInventory(Inventory *[]model.Inventory) (err error) {
	if err = config.DB.Find(Inventory).Error; err != nil {
		return err
	}
	return nil
}

//CreateAlbum - creates an album
func AddProductToInventory(Inventory *model.Inventory) (err error) {
	if err = config.DB.Create(Inventory).Error; err != nil {
		return err
	}
	return nil
}

//GetAlbum fetch one alb um
func GetInventory(Inventory *model.Inventory, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(Inventory).Error; err != nil {
		return err
	}
	return nil
}

//UpdateAlbum - modifies an album
func UpdateInventory(Inventory *model.Inventory, id string) (err error) {
	fmt.Println(Inventory)
	config.DB.Save(Inventory)
	return nil
}

//DeleteAlbum deletes a given album name
func DeleteInventory(Inventory *model.Inventory, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(Inventory)
	return nil
}

//Searching for an item in inventory
func GetProductDetailsByName(inv *[]model.Inventory, name string) (err error) {
	if err = config.DB.Where("product_name like ?", name+"%").Find(inv).Error; err != nil {
		return err
	}
	return nil
}

func GetProductDetailsById(inv *[]model.Inventory, id string) (err error) {
	if err = config.DB.Where("product_id = ?", id).Find(inv).Error; err != nil {
		return err
	}
	return nil
}
