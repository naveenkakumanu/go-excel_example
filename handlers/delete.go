package handlers

import (
	"log"
	"os"
)

func DeleteSpreadSheet(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		log.Println(err)
	}
	return err

}
