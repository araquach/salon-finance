package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/jinzhu/gorm"
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
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	return db
}

func GetCategories() (c map[string][]string) {
	c = map[string][]string{
		"Wages": {
			"MR J SHARP",
			"JAMES SHARPE",
			"MISS L HALL",
			"GEORGIA LUTTON",
			"LAUREN THOMPSON",
			"LAYLA RELF",
			"ABBI GREEN",
			"MISS V ROWLAND",
			"KELLY REEDY",
			"JOANNE MAHONEY",
			"BRADLEY RYAN",
			"ABIGAIL CLARKE",
			"DAVID RANDLES",
			"LUCY WATSON",
			"LAUREN WATSON",
			"RUBY JOHNSON",
			"SOPHIE YOUDS",
			"BETH BROWN",
			"HARRISON DOOLEY",
			"SARAH CARTWRIGHT",
			"KATE O HALLORAN",
			"DOROTA SOKOLOWSKI",
			"VICTORIA NYLAND",
			"LOUISE BAILEY",
			"LILLY SMITH",
		},
		"PAYE": {
			"HMRC CUSTOMS AND E",
			"FREDRICKSON",
			"ADVANTIS",
		},
		"Pension": {
			"NEST",
		},
		"Freelance": {
			"NATALIE SHARPE",
			"MATHEW LANE",
			"MATTHEW LANE",
			"AMY WOODS",
			"MICHELLE RAILTON",
			"LEON PRITCHARD",
		},
		"VAT": {
			"HMRC VAT",
		},
		"Utilities": {
			"BRIT GAS",
			"BRITISH TELECOM",
			"SCOTTISH POWER",
			"EDF ENERGY",
			"WATER PLUS",
			"EE LIMITED",
			"BT GROUP",
			"BG BUSINESS",
			"ASH WASTE",
			"CATHEDRAL LEASING",
			"WWW.BRITISHGAS.CO.",
			"PASTDUE",
			"SCOTTISHPOWER",
			"O2 DEVICE PLAN",
			"O2 05056477/001",
			"EE & T-MOBILE",
			"E.ON",
		},
		"Building": {
			"MEREHALL ESTATES",
			"BETTERBOOZE LTD",
			"WBC NNDR",
			"JENSON INVESTMENTS",
			"WARRINGTON BOROUGH",
			"WARRINGTON CD",
			"WARRINGTON B.C.",
			"W.B.C MV INTERNET",
			"JENSEN INVESTMENTS",
		},
		"Base": {
			"NJS MAINTENANCE",
			"M SUTTON",
			"MODERN LIGHTING",
			"M A SUTTON",
			"JACOB INTERIORS",
		},
		"Stock": {
			"BEAUTY WORKS",
			"BEAUTYWORKS",
			"ICON CONSULTANCY",
			"SWEET SQUARED",
			"ALAN HOWARD",
			"HENKEL",
			"WWW.ASTONANDFINCHE",
			"ALAN HOWARD(STOCKP CD 6954",
			"SALLY SALON",
			"WWW.GHDHAIR.COM",
			"FEEL FOR HAIR",
			"SALONS DIRECT",
			"GOCARDLESS",
			"FA UK LIMITED",
			"JAMELLA",
			"WWW.SALONSDIRECT.C",
			"BALMAINHAIR.CO.UK",
			"WWW.FEEFORHAIR.CO.",
			"JEMELLA LTD",
			"THE WIGGINS",
			"SIMPLYHAIR",
			"SALONEASY",
			"JEMELLA",
			"CLOUD9",
			"BEAUTY WORX",
			"AMERICAN CREW",
		},
		"Marketing": {
			"RACKSPACE",
			"GOOGLE",
			"TEXTANYWHERE",
			"BUFFER",
			"HEROKU",
			"FACEBK",
			"ADOBE",
			"JetBrains",
			"DIGITALOCEAN.COM",
			"FORGE.LARAVEL.COM",
			"LARACASTS",
			"NDEVOR",
			"COSCHEDULE.COM",
			"COSCHEDULE",
			"THREE BEST RATED",
			"VUEMASTERY.COM",
			"DNSIMPLE",
			"123 REG",
			"WWW.DISCOUNTDISPLA",
			"WINDOWFILMS",
			"THE PRINTING PEOPL",
			"SG MANUFACTURING",
			"POST OFFICE SELF",
			"PENTANGLE CARDS",
			"GRAFENIA",
			"Evernote",
			"DISCOUNT DISPLAYS",
			"CARTRIDGEPEOPLE.CO",
		},
		"Insurance": {
			"CLOSE-COVERSURE",
			"BAUER CONSUMER MED",
			"VLS RE KLARNA",
			"VLS RE CLOSE BROS",
			"GROVE-DEAN.CO.UK",
			"CURRYS  3267567149",
		},
		"Tax": {
			"HMRC NDDS",
			"HMRC - ACCOUNTS OF",
			"HMRC 600000000562153302",
		},
		"Staff": {
			"TRAINLINE",
			"D WRIGHT",
			"PARAGON",
			"Trainline.com",
			"Village Hotel Warr",
			"VIRGIN TRAINS",
			"THE CUMBERLAND",
			"SUITES HOTEL KNOWS",
			"SALON PUNK",
			"PAPA JONES PIZZA",
			"MR LAU'S",
			"LAS RAMBLAS",
			"Just Eat",
			"GREAT CUMBERLAND",
			"FRIAR PENKETH",
			"FIRE EVENTS",
			"DMN/DIRTYMARTINIMA",
			"Circo",
			"AIRBNB",
		},
		"Sundries": {
			"Spotify",
			"PPLPRS",
			"VIMTO OUT",
			"DLT MEDIA",
			"WWW.GOMPELS.CO.UK",
			"DJ DRINK SOLUTIONS",
			"VIKING",
			"TILLROLLSDIRECT.CO",
			"WWW.COSTCO.CO.UK",
			"PPL PRS",
			"POUNDLAND",
			"POUND SUPER STORE",
			"MM NEWSAGENTS",
			"MARTIN MCCOLL",
			"COSTCO",
		},
		"Paypal": {
			"PAYPAL",
		},
		"Amazon": {
			"AMZNMKTPLACE ",
			"AMZNMktplace",
			"Amazon",
			"AMZN",
			"AMZ*BC",
			"AMAZON",
		},
		"Loans": {
			"KENNET",
			"INVESTEC",
			"QUANTUM",
			"JOHN LAMB",
			"D A CARTER",
		},
		"Bank": {
			"NON-GBP TRANS FEE",
			"O/DRAFT INTEREST",
			"SERVICE CHARGES",
			"RETAIL MERCHANT SE",
			"GLOBAL PAYMENTS",
			"EMS",
			"UNAUTH'D BORR. FEE",
			"RETURNED D/D",
			"OVERDRAFT FEE",
			"NON-STG TRANS FEE",
		},
		"Drawings": {
			"ADAM CARTER",
			"MISS I LAMB",
			"NETFLIX.COM",
			"APPLE.COM/BILL",
			"ITUNES.COM/BILL",
			"Audible.co.uk",
			"STARBUCKS",
			"LNK WARRINGTON",
			"Prime",
			"VISION DIRECT",
			"Kindle",
			"Amazon Prime*MB8HF",
			"WWW.ASOS.COM",
			"WWW.TWOSEASONS.CO.",
			"Audible UK",
			"WWW.MISSGUIDED.CO.",
			"VOLCOM SAS",
			"TOPPIK.CO.UK",
			"TICKETMASTER",
			"PropellerheadsCOM",
			"MCDONALDS",
			"GRUMPY MULE",
			"FGT*TICKETMASTER",
			"Etsy.com",
			"EASYJET",
			"DANIEL HANCOCK",
			"CHATURBIL",
			"APPLE.COM/UK",
			"ACOUSTIC CAFE",
		},
	}
	return
}

func loadCosts() {
	var err error
	var costs []Cost

	categories := GetCategories()

	db := dbConn()
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
		db = dbConn()
		db.LogMode(true)
		db.Create(&t)
		if err != nil {
			log.Panic(err)
		}
		db.Close()
	}
}

func loadTakings() {
	var err error
	var files []string
	var takings []Taking

	db := dbConn()
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
		fmt.Println(f)
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

				if !strings.Contains(record[0], "Page") && !strings.Contains(record[0], "Total") {
					s, _ := strconv.ParseFloat(record[2], 8)
					p, _ := strconv.ParseFloat(record[6], 8)

					takings = append(takings, Taking{
						Date:     record[0],
						Name:     stylist,
						Salon:    f[0],
						Services: s,
						Products: p,
					})
				}
			}
		}
		for _, t := range takings {
			db = dbConn()
			db.LogMode(true)
			db.Create(&t)
			if err != nil {
				log.Println(err)
			}
			db.Close()
		}
	}
}