package handlers

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func CreatSpreadSheet(fileName string) (*excelize.File, error) {
	// Creating New Excel File
	f := excelize.NewFile()
 
	// Create a Newsheet
	// Sheet1 will automatically created we can insert values
	_, err := f.NewSheet("Sheet2")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// saving file as required Name
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return f, err

}
