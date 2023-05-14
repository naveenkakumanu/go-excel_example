package handlers

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func UpdateSpreadSheet(filename string) error {

	f, err := excelize.OpenFile(filename)
	if err != nil {
		log.Println(err)
		return err
	}
	f.Save()
	return err
}
