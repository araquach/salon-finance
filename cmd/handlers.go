package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplIndex.Execute(w, nil); err != nil {
		panic(err)
	}
}

func getCategories() (c map[string][]string) {
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

func apiCostsCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var d Totals
	var c GetCost

	db := dbConn()
	db.LogMode(true)
	defer db.Close()

	cat := "Wages"

	db.Table("costs").Select("sum(debit) as d").Where("category = ?", cat).Scan(&c)
	json, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)

}

//func apiAddCategory(w http.ResponseWriter, r *http.Request) {
//	var bankData BankData
//	params := mux.Vars(r)
//	id := params["id"]
//	json.NewDecoder(r.Body).Decode(&bankData)
//	db := dbConn()
//	db.Model(&bankData).Where("id =" + id).Update("category", bankData.Category)
//	db.Close()
//}
//
func apiTakings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var s Total
	// var p Total
	// var d Total

	db := dbConn()
	db.LogMode(true)
	defer db.Close()

	// stylist := "Adam Carter"

	dateFrom := "2019-08-16"
	dateTo := "2020-08-16"

	db.Table("takings").Select("sum(services) as s, sum(products) as p").Where("date >= ? AND date <= ?", dateFrom, dateTo).Where("salon = ?", "Jakata").Scan(&s)

	// db.Table("takings").Select("sum(services) as s, sum(products) as p").Where("name = ?", stylist).Scan(&s)

	//s.C = stylist
	s.D = s.P + s.S

	json, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func apiMonthlyTakings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := dbConn()
	t := []Taking{}

	params := mux.Vars(r)
	month := params["month_year"]

	db.Where("month_year = ?", month).Find(&t)
	db.Close()

	json, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
