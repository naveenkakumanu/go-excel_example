package main

import (
	"log"

	"github.com/naveenkakumanu/go-excel_example/handlers"
)

func main() {
	fileName := "Book.xlsx"

	fileDetails, err := handlers.CreatSpreadSheet(fileName)
	if err != nil {
		log.Println(err)
	}
	log.Println(fileDetails.Theme.Name)

	err = handlers.ReadSpreadSheet(fileName)
	if err != nil {
		log.Println(err)
	}

	err = handlers.UpdateSpreadSheet(fileName)
	if err != nil {
		log.Println(err)
	}

	err = handlers.DeleteSpreadSheet(fileName)
	if err != nil {
		log.Println(err)
	}

}
