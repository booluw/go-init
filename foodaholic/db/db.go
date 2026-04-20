package db

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"time"
)

/*
* This module reads and writes to the db
* DB is a csv file with rows and columns to match
 */

const FILENAME = "foodaholic.csv"

func initDB(headers []string, fileName string) bool {
	if _, err := os.Stat(fmt.Sprintf("bills/%v.csv", fileName)); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		file, err := os.OpenFile(fmt.Sprintf("bills/%v.csv", fileName), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

		if err != nil {
			panic(err)
		}

		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		writer.Write(headers)
		return true
	} else {
		panic(err)
	}
}

func WriteToDB(order []string, customerName string) bool {
	initDB([]string{"Date", "Order", "Price ($)"}, customerName)

	file, err := os.OpenFile(fmt.Sprintf("bills/%v.csv", customerName), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return false
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(append([]string{time.Now().Format("2006-01-02")}, order...)); err != nil {
		panic(err)
	}

	return true
}
