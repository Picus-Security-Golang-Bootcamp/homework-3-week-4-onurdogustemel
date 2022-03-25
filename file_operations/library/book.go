package library

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title          string  `gorm:"type:varchar(100);column:Title"`
	Page           uint64  `gorm:"type:uint;column:Page"`
	Author         string  `gorm:"type:varchar(100);column:Author"`
	NumberOfStocks int     `gorm:"type:int;column:NumberOfStocks"`
	Price          float64 `gorm:"type:float;column:Price"`
	StockCode      string  `gorm:"type:varchar(100);column:StockCode"`
	Isbn           string  `gorm:"type:varchar(100);column:Isbn"`
}

type BookListSlice []Books
