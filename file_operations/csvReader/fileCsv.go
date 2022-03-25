package csvReader

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/onurdogustemel/file_operations/file_operations/library"
)

func ReadCSVFile(filename string) (library.BookListSlice, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		return nil, err
	}

	var bookList library.BookListSlice

	for _, values := range records[1:] {

		pageNumber, err := strconv.ParseUint(values[1], 10, 64)
		if err != nil {
			fmt.Println("wrong type", err)
		}

		stocks, err := strconv.Atoi(values[3])
		if err != nil {
			fmt.Println("wrong type", err)
		}

		pr, err := strconv.ParseFloat(values[4], 64)
		if err != nil {
			fmt.Println("wrong type", err)
		}

		bookList = append(bookList, library.Books{
			Title:          values[0],
			Page:           pageNumber,
			Author:         values[2],
			NumberOfStocks: stocks,
			Price:          pr,
			StockCode:      values[5],
			Isbn:           values[6],
		})

	}
	return bookList, err
}

func ReadCSVFileForAuthor(filename string) (library.AuthorSlice, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		return nil, err
	}

	var authorList library.AuthorSlice

	for _, values := range records[1:] {

		authorList = append(authorList, library.Authors{
			Author: values[0],
		})

	}
	return authorList, err
}
