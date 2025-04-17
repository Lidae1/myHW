package main

import (
	"Zadacha/Present"
	"Zadacha/Produce"
	"Zadacha/Service"
)

func main() {
	prod := &Produce.RealProduce{}
	pres := &Present.RealPresent{}
	service := Service.NewService(prod, pres)
	service.Run("URL.txt")
}
