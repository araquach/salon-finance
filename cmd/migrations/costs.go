package migrations

import (
	"bufio"
	"encoding/csv"
	"github.com/araquach/salon-finance/cmd/helpers"
	"github.com/araquach/salon-finance/cmd/models"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Return a list of files from a directory
func readDir(dir string) ([]string, error) {
	var err error
	var files []string

	root := "data/finance/" + dir
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, err
}

// Import costs data into a struct
func costData() []models.Cost {
	var err error
	var costs []models.Cost

	f, err := readDir("bank")
	if err != nil {
		panic(err)
	}

	for _, file := range f {
		csvFile, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		reader := csv.NewReader(bufio.NewReader(csvFile))
		for {
			col, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			date, _ := time.Parse("2006-01-02", helpers.DateFormat(col[0]))
			debit, _ := strconv.ParseFloat(col[5], 8)

			if col[5] != "" && col[1] != "TFR" && col[1] != "" {
				costs = append(costs, models.Cost{
					Date:        date,
					Type:        col[1],
					Account:     col[3],
					Description: col[4],
					Debit:       debit,
					Category:    "uncategorised",
					SubCat:      "uncategorised",
				})
			}
		}
	}
	return costs
}

// import Paypal data into a struct
func payPalData() []models.Cost {
	var err error
	var paypal []models.Cost

	f, err := readDir("paypal")
	if err != nil {
		panic(err)
	}

	for _, file := range f {
		csvFile, _ := os.Open(file)

		reader := csv.NewReader(bufio.NewReader(csvFile))
		for {
			col, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			date, _ := time.Parse("2006-01-02", helpers.DateFormat(col[0]))
			d := strings.Replace(col[5], ",", "", -1)
			debit, _ := strconv.ParseFloat(d[1:], 8)

			if col[11] != "" {
				paypal = append(paypal, models.Cost{
					Date:        date,
					Type:        "PAYPAL",
					Account:     "06517160",
					Description: col[11],
					Debit:       debit,
					Category:    "uncategorised",
					SubCat:      "uncategorised",
				})
			}
		}
	}
	return paypal
}

func amazonData() []models.Cost {
	var err error
	var amazon []models.Cost

	f, err := readDir("amazon")
	if err != nil {
		panic(err)
	}

	for _, file := range f {
		csvFile, _ := os.Open(file)

		reader := csv.NewReader(bufio.NewReader(csvFile))
		for {
			col, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			date, _ := time.Parse("2006-01-02", helpers.DateFormat(col[0]))
			debit, _ := strconv.ParseFloat(col[24], 8)

			if col[37] != "" {
				amazon = append(amazon, models.Cost{
					Date:        date,
					Type:        "AMAZON",
					Account:     "06517160",
					Description: col[37],
					Debit:       debit,
					Category:    "uncategorised",
					SubCat:      "uncategorised",
				})
			}
		}
	}
	return amazon
}

func addCostCategories() []models.Cost {
	costs := costData()
	cats := categories()

	for k, v := range costs {
		for kk, vv := range cats {
			for kkk, vvv := range vv {
				for _, av := range vvv {
					if strings.Contains(v.Description, av) {
						(costs)[k].Category = kk
						(costs)[k].SubCat = kkk
					}
				}
			}
		}
	}
	return costs
}

func addPayPalCategories() []models.Cost {
	pp := payPalData()
	cats := paypalCats()

	for k, v := range pp {
		for kk, vv := range cats {
			for kkk, vvv := range vv {
				for _, av := range vvv {
					if av == v.Description {
						(pp)[k].Category = kk
						(pp)[k].SubCat = kkk
					}
				}
			}
		}
	}
	return pp
}

func addAmazonCategories() []models.Cost {
	amzn := amazonData()
	cats := amazonCats()

	for k, v := range amzn {
		for kk, vv := range cats {
			for kkk, vvv := range vv {
				for _, av := range vvv {
					if av == v.Description {
						(amzn)[k].Category = kk
						(amzn)[k].SubCat = kkk
					}
				}
			}
		}
	}
	return amzn
}
