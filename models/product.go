package models

import (
	"log"
	"time"
)

var Categories []string

type Product struct {
	ID            uint      `db:"id"`
	Code          string    `db:"code"`
	Title         string    `db:"title"`
	Category      string    `db:"category"`
	Maker         string    `db:"maker"`
	Description   string    `db:"description"`
	Ean           string    `db:"ean"`
	Ncm           string    `db:"ncm"`
	WarrantyMonth uint      `db:"warranty_month"`
	LengthMM      uint      `db:"length_mm"`
	WidthMM       uint      `db:"width_mm"`
	HeightMM      uint      `db:"height_mm"`
	WeightG       uint      `db:"weight_g"`
	IsActive      bool      `db:"is_active"`
	IsAvailable   bool      `db:"is_available"`
	IsSell        bool      `db:"is_sell"`
	StockOrigin   string    `db:"stock_origin"`
	StockCount    uint      `db:"stock_count"`
	PriceBuy      uint      `db:"price_buy"`
	PriceSale     uint      `db:"price_sale"`
	SupplierCode  string    `db:"supplier_code"`
	SupplierName  string    `db:"supplier_name"`
	CreatedAt     time.Time `db:"created_at"`
	ChangedAt     time.Time `db:"changed_at"`
	IsDeleted     bool      `db:"is_deleted"`
}

func GetAllProduct() (products []Product, err error) {
	err = pgDB.Select(&products, "select * from products where is_deleted=$1", false)
	if err != nil {
		log.Printf("[error] %v\n", err)
		return
	}
	return
}

func UpdateCategories() {
	Categories = []string{"Notebook", "Monitor", "HD"}
}
