package db

import (
	"encoding/csv"
	"fmt"
	"os"
)

/*
* This module reads and writes to the db
* DB is a csv file with rows and columns to match
 */

const FILENAME = "foodaholic.csv"

func init() {
	isInit := initDB("order, price")

	if isInit {
		fmt.Println("DB has been created")
	} else {
		fmt.Println("DB failed")
	}
}

func initDB(headers string) bool {
	file, err := os.Create(FILENAME)

	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()
	defer file.Close()

	writer.Write([]string{ "Order", "Price" })

	return true
}

func WriteToDB(order []string) bool {
	file, err := os.OpenFile(FILENAME, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return false
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(order); err != nil {
		panic(err)
	}

	return true
}
