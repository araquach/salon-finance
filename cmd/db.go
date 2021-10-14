package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func dbConn() (db *gorm.DB) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func loadCosts() {
	var err error
	var costs []Cost

	categories := GetCategories()

	db := dbConn()
	db.Migrator().DropTable(&Cost{})
	err = db.AutoMigrate(&Cost{})
	if err != nil {
		panic(err)
	}

	var files []string

	root := "data/finance/bank"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		csvFile, _ := os.Open(file)

		reader := csv.NewReader(bufio.NewReader(csvFile))
		for {
			line, error := reader.Read()
			if error == io.EOF {
				break
			} else if error != nil {
				log.Fatal(error)
			}
			d, _ := strconv.ParseFloat(line[5], 8)
			b, _ := strconv.ParseFloat(line[7], 8)
			date, _ := time.Parse("2006-01-02", dateFormat(line[0]))

			if line[5] != "" && line[1] != "TFR" && line[1] != "" {
				costs = append(costs, Cost{
					Date:        date,
					Type:        line[1],
					Account:     line[3],
					Description: line[4],
					Debit:       d,
					Balance:     b,
					Category:    "Other",
					SubCat:      "other",
				})
			}
		}
	}
	for _, t := range costs {
		for i, cat := range categories {
			for j, c := range cat {
				for _, s := range c {
					if strings.Contains(t.Description, s) {
						t.Category = i
						t.SubCat = j
					}
				}
			}
		}
		db.Create(&t)
	}
}

func loadTakings() {
	var err error
	var files []string
	var takings []Taking

	db := dbConn()
	db.Migrator().DropTable(&Taking{})
	err = db.AutoMigrate(&Taking{})
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
					date, _ := time.Parse("2006-01-02", dateFormatYear(record[0]))

					takings = append(takings, Taking{
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
		db.Create(&t)
		if err != nil {
			log.Println(err)
		}
	}
}
