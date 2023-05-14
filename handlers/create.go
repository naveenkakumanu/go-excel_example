package handlers

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

// Creating New Excel File and Inserting Data in Sheet1
func CreatSpreadSheet(fileName string) (*excelize.File, error) {
	
	f := excelize.NewFile()

	// Create a Newsheet
	// Sheet1 will automatically created we can insert values
	_, err := f.NewSheet("Sheet2")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Inserting Data in Sheet1 
	f.SetCellValue("Sheet1", "A1", "Name")
	f.SetCellValue("Sheet1", "B1", "Age")

	f.SetCellValue("Sheet1", "A2", "Narendra")
	f.SetCellValue("Sheet1", "B2", 31)

	f.SetCellValue("Sheet1", "A3", "Suresh")
	f.SetCellValue("Sheet1", "B3", 32)

	f.SetCellValue("Sheet1", "A4", "Ramesh")
	f.SetCellValue("Sheet1", "B4", 31)

	// Inserting Data in Sheet2

	f.SetCellValue("Sheet2", "A1", "Name")
	f.SetCellValue("Sheet2", "B1", "Age")

	f.SetCellValue("Sheet2", "A2", "Narendra")
	f.SetCellValue("Sheet2", "B2", 31)

	f.SetCellValue("Sheet2", "A3", "Suresh")
	f.SetCellValue("Sheet2", "B3", 32)

	f.SetCellValue("Sheet2", "A4", "Ramesh")
	f.SetCellValue("Sheet2", "B4", 31)

	// saving file as required Name
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return f, err

}
