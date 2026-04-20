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
	headers := []string{"Date", "Order", "Price"}
	isInit := initDB(headers)

	if isInit {
		fmt.Println("DB has been created")
	} else {
		fmt.Println("DB failed")
	}
}

func initDB(headers []string) bool {
	file, err := os.OpenFile(FILENAME, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(headers)

	return true
}

func WriteToDB(order [][]string) bool {
	file, err := os.OpenFile(FILENAME, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return false
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.WriteAll(order); err != nil {
		panic(err)
	}

	return true
}
