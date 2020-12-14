package finance

import (
	"bufio"
	"encoding/csv"
	"github.com/araquach/salon-finance/cmd/db"
	"github.com/araquach/salon-finance/cmd/helpers"
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
	var err error
	var costs []Cost

	categories := GetCategories()

	db := db.DbConn()
	db.LogMode(true)
	db.AutoMigrate(&Cost{})
	db.Close()

	var files []string

	root := "data/bank"
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
			date, _ := time.Parse("2006-01-02", helpers.DateFormat(line[0]))

			if line[5] != "" && line[1] != "TFR" && line[1] != "" {
				costs = append(costs, Cost{
					Date:        date,
					Type:        line[1],
					Account:     line[3],
					Description: line[4],
					Debit:       d,
					Balance:     b,
					Category:    "Other",
				})
			}
		}
	}
	for _, t := range costs {
		for i, cat := range categories {
			for _, c := range cat {
				if strings.Contains(t.Description, c) {
					t.Category = i
				}
			}
		}
		db.LogMode(true)
		db.Create(&t)
		if err != nil {
			log.Panic(err)
		}
		db.Close()
	}
}

func LoadTakings() {
	var err error
	var files []string
	var takings []Taking

	db := db.DbConn()
	db.LogMode(true)
	db.AutoMigrate(&Taking{})
	db.Close()

	root := "data/takings"

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
		f := strings.Split(fname[2], "_")
		fileBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Println(err)
		}

		split := strings.SplitAfter(string(fileBytes), "\n")[3:]

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

			s := strings.SplitAfter(l, "\n")

			stylist := s[0]

			stylist = strings.Split(stylist, ",")[0]

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
		db.LogMode(true)
		db.Create(&t)
		if err != nil {
			log.Println(err)
		}
		db.Close()
	}
}
