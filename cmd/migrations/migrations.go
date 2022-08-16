package migrations

import (
	"encoding/csv"
	"fmt"
	"github.com/araquach/salon-finance/cmd/db"
	"github.com/araquach/salon-finance/cmd/helpers"
	"github.com/araquach/salon-finance/cmd/models"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func LoadCosts() {
	var data []models.Cost

	db.DB.Migrator().DropTable(&models.Cost{})
	db.DB.AutoMigrate(&models.Cost{})

	costs := addCostCategories()
	pp := addPayPalCategories()
	amzn := addAmazonCategories()

	for _, v := range costs {
		sd := time.Date(2019, 01, 07, 0, 0, 0, 0, time.UTC)
		if v.Category == "paypal" && sd.After(v.Date) {
			data = append(data, models.Cost{Date: v.Date, Type: v.Type, Account: v.Account, Description: v.Description, Debit: v.Debit, Category: v.Category, SubCat: v.SubCat})
		}
		if v.Category != "paypal" && v.Category != "amazon" {
			data = append(data, models.Cost{Date: v.Date, Type: v.Type, Account: v.Account, Description: v.Description, Debit: v.Debit, Category: v.Category, SubCat: v.SubCat})
		}
	}

	for _, v := range pp {
		data = append(data, models.Cost{Date: v.Date, Type: v.Type, Account: v.Account, Description: v.Description, Debit: v.Debit, Category: v.Category, SubCat: v.SubCat})
	}

	for _, v := range amzn {
		data = append(data, models.Cost{Date: v.Date, Type: v.Type, Account: v.Account, Description: v.Description, Debit: v.Debit, Category: v.Category, SubCat: v.SubCat})
	}

	db.DB.Create(&data)
}

func LoadTakings() {
	var err error
	var files []string
	var takings []models.Taking

	db.DB.Migrator().DropTable(&models.Taking{})
	err = db.DB.AutoMigrate(&models.Taking{})
	if err != nil {
		panic(err)
	}

	root := "data/finance/takings"

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
		fname := strings.Split(fileName, "/")
		f := strings.Split(fname[3], "_")
		fmt.Println(f)
		fileBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Println(err)
		}

		split := strings.SplitAfter(string(fileBytes), "\n")

		var section []string

		for _, line := range split {
			if string(line[0]) == "T" {
				line = "X"
			}
			if !strings.Contains(line, "Page") {
				section = append(section, line)
			}
		}

		joined := strings.Join(section, "")
		formatted := strings.Split(joined, "X")

		for _, l := range formatted {
			var data []string
			var stylist string

			s := strings.SplitAfter(l, "\n")

			stylistLine := s[0]

			stylistBlock := strings.Split(stylistLine, ",")[0]

			stylistBF := strings.SplitAfter(stylistBlock, " ")

			if len(stylistBF) > 1 {
				if stylistBF[1] == " " {
					stylist = stylistBF[0] + strings.TrimSpace(stylistBF[2])
				} else {
					stylist = stylistBF[0] + strings.TrimSpace(stylistBF[1])
				}
			}

			if len(s) > 1 {
				data = s[3:]
			}

			csvReady := strings.Join(data, "")

			r := csv.NewReader(strings.NewReader(csvReady))

			for {
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Println(err)
				}

				if len(record[0]) > 0 && !strings.Contains(record[0], "Page") && !strings.Contains(record[0], "Total") {
					s, _ := strconv.ParseFloat(record[2], 8)
					p, _ := strconv.ParseFloat(record[6], 8)
					date, _ := time.Parse("2006-01-02", helpers.DateFormatYear(record[0]))

					stylist := helpers.ChangeName(stylist)

					takings = append(takings, models.Taking{
						Date:     date,
						Name:     stylist,
						Salon:    f[0],
						Services: s,
						Products: p,
					})
				}
			}
		}
	}
	for _, t := range takings {
		db.DB.Create(&t)
		if err != nil {
			log.Println(err)
		}
	}
}
