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

	f.SetCellValue("Sheet2", "A5", "Indira")
	f.SetCellValue("Sheet2", "B5", 31)

	f.Save()
	return err
}
