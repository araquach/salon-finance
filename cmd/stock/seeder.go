package stock

import (
	"bufio"
	"encoding/csv"
	"github.com/araquach/salon-finance/cmd/db"
	"github.com/araquach/salon-finance/cmd/helpers"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func LoadProfessional() {
	db := db.DbConn()

	var files []string
	var stockTransfers []ProductData
	var err error

	root := "data/stock/stock-out/Professional"

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".csv" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	for _, fileName := range files {
		fname := filepath.Base(fileName)
		split := strings.Split(fname, " ")
		salon := split[0]
		date := "01-" + strings.Split(split[1], ".")[0]

		csvFile, _ := os.Open(fileName)

		sod := csv.NewReader(bufio.NewReader(csvFile))

		for {
			sor, err := sod.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
			}

			sl :="data/stock/stock-list/Professional.csv"

			csvFile2, _ := os.Open(sl)

			sld := csv.NewReader(bufio.NewReader(csvFile2))

			for {
				slr, err := sld.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Println(err)
				}

				if slr[0] == sor[0] {
					formattedDate, _ := time.Parse("2006-01-02", helpers.DateFormatDashes(date))
					quantity, _ := strconv.Atoi(sor[1])
					totalPrice, _ := strconv.ParseFloat(sor[2], 8)
					price := totalPrice / float64(quantity)

					stockTransfers = append(stockTransfers, ProductData{Date: formattedDate, Salon: salon, Product: sor[0], Quantity: quantity, Price: price, Supplier: slr[1], Brand: slr[2], Category: slr[3], SubBrand: slr[4], Type: slr[5]})
				}
			}
		}
	}
	for _, r := range stockTransfers {

		db.Create(&r)
		if err != nil {
			log.Println(err)
		}
	}
}

func LoadRetail() {
	db := db.DbConn()

	var files []string
	var stockTransfers []ProductData
	var err error

	root := "data/stock/stock-out/Retail"

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".csv" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	for _, fileName := range files {
		fname := filepath.Base(fileName)
		split := strings.Split(fname, " ")
		salon := split[0]
		date := "01-" + strings.Split(split[1], ".")[0]

		csvFile, _ := os.Open(fileName)

		sod := csv.NewReader(bufio.NewReader(csvFile))

		for {
			sor, err := sod.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
			}

			sl :="data/stock/stock-list/Retail.csv"

			csvFile2, _ := os.Open(sl)

			sld := csv.NewReader(bufio.NewReader(csvFile2))

			for {
				slr, err := sld.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Println(err)
				}

				if slr[0] == sor[0] {

					formattedDate, _ := time.Parse("2006-01-02", helpers.DateFormatDashes(date))
					quantity, _ := strconv.Atoi(sor[4])
					price, _ := strconv.ParseFloat(sor[1], 8)

					stockTransfers = append(stockTransfers, ProductData{Date: formattedDate, Salon: salon, Product: sor[0], Quantity: quantity, Price: price, Supplier: slr[1], Brand: slr[2], Category: slr[3], SubBrand: slr[4], Type: slr[5]})
				}
			}
		}
	}
	for _, r := range stockTransfers {
		db.Create(&r)
		if err != nil {
			log.Println(err)
		}
	}
}