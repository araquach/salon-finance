package migrations

func categories() (c map[string]map[string][]string) {
	c = map[string]map[string][]string{
		"wages": {
			"stylist": {
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
				"JAMIE BANNING-HIGH",
				"EVE SIMPSON",
				"REBECCA COTTON",
				"SARAH CUDDY",
				"SHAE HOUGHTON",
				"MISS N BAILEY",
				"MAISIE THOMPSON",
				"CALEB BARRIE",
				"LAURA MINETT",
			},
			"freelance": {
				"NATALIE SHARPE",
				"MATHEW LANE",
				"MATTHEW LANE",
				"AMY WOODS",
				"MICHELLE RAILTON",
				"LEON PRITCHARD",
			},
			"paye": {
				"HMRC CUSTOMS AND E",
				"FREDRICKSON",
				"ADVANTIS",
			},
			"pension": {
				"NEST",
				"WWW.THEPENSIONSREG",
			},
		},
		"taxes": {
			"vat": {
				"HMRC VAT",
				"HMRC E VAT",
			},
			"tax": {
				"HMRC NDDS",
				"HMRC - ACCOUNTS OF",
				"HMRC 600000000562153302",
				"HMRC 300000000784948091",
				"HMRC 200000000588802019",
				"HMRC 400000000829117594 3002558756K",
				"1ST LOCATE UK",
			},
		},
		"utilities": {
			"energy": {
				"BRIT GAS",
				"SCOTTISH POWER",
				"EDF ENERGY",
				"BG BUSINESS",
				"WWW.BRITISHGAS.CO.",
				"SCOTTISHPOWER",
				"E.ON",
				"EXTRA ENERGY",
				"GAZPROM ENERGY",
				"EXTRA ENERGY",
			},
			"water": {
				"WATER PLUS",
				"PASTDUE",
			},
			"telecoms": {"BRITISH TELECOM",
				"EE LIMITED",
				"BT GROUP",
				"O2 DEVICE PLAN",
				"O2 05056477/001",
				"EE & T-MOBILE",
				"VIRGIN MEDIA PYMTS",
			},
			"waste": {
				"ASH WASTE",
				"CATHEDRAL LEASING",
				"CATHEDRAL HYGIENE",
			},
			"other": {
				"CHUBB",
				"WWW.MY.CHUBB.CO.UK",
			},
		},
		"building": {
			"rent": {
				"MEREHALL ESTATES",
				"BETTERBOOZE LTD",
				"JENSON INVESTMENTS",
				"JENSEN INVESTMENTS",
			},
			"rates": {
				"WARRINGTON BOROUGH",
				"WARRINGTON CD",
				"WARRINGTON B.C.",
				"W.B.C MV INTERNET",
				"WBC NNDR",
				"WARRINGTON BC",
			},
			"repairs": {
				"CITY ALUMINIUM SHO",
				"MR ROBERT A ANTONO",
				"A6 ROLLER SHUTTERS",
				"SECURITY CENTRES",
				"Selco Trade Centre",
				"WWW.ILOVEWALLPAPER",
				"SCREWFIX",
				"SELCO POOLE",
				"DEREK BYRNE",
				"CITY PLUMBING SUPP",
				"WARRINGTON GLASS",
				"B & Q 1269",
				"PLUMBASE LTD",
				"DELTA GLAZING",
			},
			"salon refits": {
				"MR J SHARP 200000000304140444",
				"NJS MAINTENANCE",
				"M SUTTON",
				"MODERN LIGHTING",
				"M A SUTTON",
				"JACOB INTERIORS",
				"AQUABLISSSA",
				"SOAK.COM",
				"INTERSERVE",
				"RICHER SOUNDS",
				"SALONAQUAFLOW",
				"DISCOUNT FLOORING",
				"Flooring UK",
				"WWW.DIGITECK.NET",
				"WWW.UKPOS.COM",
				"OHNSTONES DEC",
				"HARDWARE WAREHOUSE",
				"AL MURAD DIY",
				"SUK RETAIL",
				"GJM ENGINEERING",
				"WHERETHETRADEBUYS.",
				"SALON AQUA FLOW UK",
				"BESTHEATING",
				"SPARK RISK",
				"QVS ELEC SUPPLIES",
				"DUSKLIGHTS.CO.UK",
				"BBS*ONLINE",
				"THE ELECTRICAL SHO",
				"M and M Direct Ltd",
				"Connox GmbH",
				"WWW.UTILITYDESIGN.",
				"QVS ELEC SUPPLIES",
				"WWW.SHEIN.COM",
				"VICTORIAN PLUMBING",
				"WWW.WINDSORBROWNE.",
				"WWW.QBICWASHROOMS.",
				"WWW.ERS-ONLINE.CO.",
				"WWW.GARDENTRADINGW",
				"WWW.SHEIN.COM",
				"WWW.RICHERSOUNDS.C",
				"INTERNATIONAL LAMP",
				"eBay O*09-06030-06 CD 6954",
				"YESSS ELECTRICAL",
				"THE WAREHOUSE GROU",
				"JAMES SMITH",
			},
		},
		"stock": {
			"tigi": {
				"TIGI",
			},
			"schwarzkopf": {
				"HENKEL",
			},
			"kevin murphy": {
				"ICON CONSULTANCY",
				"SWEET SQUARED",
			},
			"ghd": {
				"WWW.GHDHAIR.COM",
				"JAMELLA",
				"JEMELLA LTD",
				"JEMELLA",
				"GHD",
			},
			"olaplex": {
				"WWW.ASTONANDFINCHE",
			},
			"extensions": {
				"BEAUTY WORKS",
				"BEAUTYWORKS",
				"BALMAINHAIR.CO.UK",
				"THE WIGGINS",
				"SIMPLYHAIR",
				"BEAUTY WORX",
				"SP * HAIRMADEEASIS",
			},
			"gowns": {
				"WWW.FEEFORHAIR.CO.",
				"FEEL FOR HAIR",
			},
			"misc": {
				"WWW.SALONSDIRECT.C",
				"SALONEASY",
				"ALAN HOWARD",
				"ALAN HOWARD(STOCKP CD 6954",
				"SALLY SALON",
				"SALONS DIRECT",
				"WWW.SALON-SERVICES",
				"FA UK LIMITED",
				"CLOUD9",
				"AMERICAN CREW",
				"GOCARDLESS",
			},
		},
		"marketing": {
			"website": {
				"RACKSPACE",
				"HEROKU",
				"DIGITALOCEAN.COM",
				"FORGE.LARAVEL.COM",
				"DNSIMPLE",
				"123 REG",
			},
			"sms": {
				"TEXTANYWHERE",
				"TEXT ANYWHERE",
				"TEXTMAGIC.COM",
			},
			"software": {
				"ADOBE",
				"Evernote",
				"WWW.PHOREST.COM",
				"NDEVOR",
				"JETBRAINS",
				"JetBrains",
				"PADDLE.NET*",
			},
			"courses": {
				"LARACASTS",
				"VUEMASTERY.COM",
				"GROW MY SALON BUSI",
			},
			"social media": {
				"BUFFER",
				"COSCHEDULE.COM",
				"COSCHEDULE",
				"LOOMLY US",
				"HTTPS://HEROPOST.I CA",
				"HOOPLA",
			},
			"advertising": {
				"GOOGLE",
				"FACEBK",
				"THREE BEST RATED",
				"MY ARK",
			},
			"printing": {},
			"signage": {
				"WWW.DISCOUNTDISPLA",
				"WINDOWFILMS",
				"THE PRINTING PEOPL",
				"SG MANUFACTURING",
				"DISCOUNT DISPLAYS",
			},
			"wages": {
				"ROSEMARIE WINTERBU",
			},
			"other": {
				"POST OFFICE SELF",
				"PENTANGLE CARDS",
				"GRAFENIA",
				"INDEED",
				"P.H.A. LTD",
				"Dropbox",
				"Mediapress Limited",
			},
		},
		"insurance": {
			"insurance": {
				"CLOSE-COVERSURE",
				"BAUER CONSUMER MED",
				"VLS RE KLARNA",
				"VLS RE CLOSE BROS",
				"GROVE-DEAN.CO.UK",
				"CURRYS  3267567149",
				"COVERSURE",
			},
		},
		"staff": {
			"transport": {
				"EURO CAR PARKS",
				"TRAINLINE",
				"Trainline.com",
				"VIRGIN TRAINS",
				"TAXI FARE BY VERIF",
				"WARRINGTON CTL",
				"thetrainline.com",
				"IZ *LONDON TAXI",
			},
			"HR": {
				"D WRIGHT",
			},
			"education": {
				"PARAGON",
				"SALON PUNK",
				"FIRE EVENTS",
				"BEHINDTHECHAIR.COM",
				"JON CALHOUN",
				"Udemy",
				"NOTANOTHERACADEMY.",
			},
			"accommodation": {
				"Hotel at Booking.c",
				"Village Hotel Warr",
				"THE CUMBERLAND",
				"SUITES HOTEL KNOWS",
				"GREAT CUMBERLAND",
				"AIRBNB",
				"Hotel on Booking.c",
				"INTERNATIONAL INN",
				"AMBA HOTEL MARBLE",
				"GROSVENOR HOUSE",
				"PREMIER INN4451265",
			},
			"social": {
				"SQ *ESCAPE HUNT",
				"PREZZO",
				"THE GRILL ON THE S",
				"PAPA JONES PIZZA",
				"MR LAU'S",
				"LAS RAMBLAS",
				"Just Eat",
				"FRIAR PENKETH",
				"DMN/DIRTYMARTINIMA",
				"Circo",
				"THE BOTANIST",
				"IAN AUCTION",
				"Wizz Air Hu",
				"GRADBACH MILL",
				"SUBWAY CD 6954",
				"GREGGS",
				"SQ *LOUIES PIZZA",
				"DEAD EYES",
				"THREE TUNS",
				"Acoustic Cafe",
			},
			"incentives": {
				"YUVA MEDISPA",
				"GIFT TIME",
				"THE PERFUME SHOP",
				"SAN LORENZO WARRIN",
				"ODEON",
			},
		},
		"sundries": {
			"music": {
				"Spotify",
				"SPOTIFY",
				"PPLPRS",
				"PPL PRS",
				"PHONOGRAPHIC PERFO",
				"PRS",
			},
			"refreshments": {
				"VIMTO OUT",
				"DJ DRINK SOLUTIONS",
				"MAJESTIC WINE",
			},
			"magazines": {
				"DLT MEDIA",
			},
			"misc": {
				"WWW.GOMPELS.CO.UK",
				"CARTRIDGE SAVE",
				"VIKING",
				"TILLROLLSDIRECT.CO",
				"CARTRIDGEPEOPLE.CO",
				"WWW.COSTCO.CO.UK",
				"POUNDLAND",
				"POUND SUPER STORE",
				"MM NEWSAGENTS",
				"MARTIN MCCOLL",
				"COSTCO",
				"B&M",
				"CHURCH AND DWIGHT",
				"THE POST OFFICE",
				"POUNDWORLD",
				"WWW.HANGERWORLD.CO",
				"HATTERS BLOOM",
				"MEGA BARGAIN STORE",
				"WWW.BENTS.CO.UK",
				"SUPERPOUND PLUS",
			},
		},
		"paypal": {
			"paypal": {
				"PAYPAL",
			},
		},
		"amazon": {
			"amazon": {
				"AMZNMKTPLACE",
				"AMZNMktplace",
				"Amazon",
				"AMZN",
				"AMZ*BC",
				"AMAZON",
				"AMZ*riversidehertf",
			},
		},
		"loans": {
			"family": {
				"JOHN LAMB",
				"D A CARTER",
			},
			"costs": {
				"LOAN",
				"BARCLAYS BANK PLC",
			},
			"finance": {
				"KENNET",
				"INVESTEC",
				"QUANTUM",
				"GE CAPITAL",
			},
		},
		"costs": {
			"costs": {
				"NON-GBP TRANS FEE",
				"O/DRAFT INTEREST",
				"SERVICE CHARGES",
				"RETAIL MERCHANT SE",
				"RMS",
				"GLOBAL PAYMENTS",
				"EMS",
				"UNAUTH'D BORR. FEE",
				"RETURNED D/D",
				"OVERDRAFT FEE",
				"NON-STG TRANS FEE",
				"ARRANGEMENT FEE",
				"JAKATA CD 6939",
				"CARD TXNS",
				"JAKATA CD",
			},
		},
		"drawings": {
			"drawings": {
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
				"SODASTREAM",
				"MADE.COM",
				"TYN-Y-MUR.CO.UK",
				"B & T SKIP",
				"MY OPTIQUE GROUP",
				"ERHM LIMITED",
				"CHILLISAUCE ",
				"MBI-PROBILLER.COM",
				"NCP LIMITED",
				"WWW.TICKETLINE.CO.",
				"IZ *ALIZA FASHIONS",
				"EXPERIAN",
				"TICKETS 0871220026",
				"SCOTTS",
				"UPPER CRUST",
				"WWW.PETSHOPBOWL.CO",
				"VIAGOGO",
				"THE GALLERY CAFE",
				"WWW.VETUK.CO.UK",
				"TICKETS 0871220026",
				"HERON FOODS LTD",
				"DECATHLON",
				"MOJO BAR",
				"Klarna * PSYCHE LT",
				"TOPPIK UK",
				"UBER   *EATS",
				"MYPROTEIN.COM",
				"HUGGLE PETS",
			},
		},
		"equipment": {
			"equipment": {
				"CURRYS CD",
				"TENZY",
				"JM-AC",
				"LADDERS UK",
				"SP * YOISCISSORS",
				"WWW ESPARES",
				"WWW.BUYSPARES.CO.U CD",
				"HOWHIGHBRANDS.CO.U",
				"AMBICOOL",
				"BARBER BLADES",
				"WWW.LEDSUPPLYANDFI",
				"CURRYS ONLINE",
				"IKEA",
				"ANY.DO US",
				"ARGOS RETAIL GROUP",
				"vidaXL UK",
				"WWW.LIGHTS4FUN.CO.",
				"TESCO",
				"COOLBLADES",
				"HOMEBASE LTD",
				"LIGO ELECTRONICS",
				"WILKINSON CAMERAS",
				"VERY.CO.UK",
				"CASTLEGATELIGHTS.C",
				"NEXT DIRECTORY",
			},
		},
		"accountant": {
			"accountant": {
				"CMT ACCOUNTING",
				"POS-HARDWARE LTD",
				"BOWDON ACCOUNTANCY",
				"BOWDEN ACCOUNTANCY",
			},
		},
		"other": {
			"refunds": {
				"RACHEL V PARRY",
				"LAUREN MINCHION",
				"KATE JOHNSON",
				"MRS A M JAMES ",
			},
			"charity": {
				"WWW.HEARINGDIRECT",
			},
			"other": {
				"WWW.SMART-STORAGE",
				"CREDIT SECURITY",
				"SQ *HENRY LEWIS",
				"MRS NICOLA COBLEY",
				"HOME BARGAINS",
				"TESCO STORES",
				"WAYFAIR",
				"WILKO",
				"SUPERDRUG",
				"WH SMITH",
				"CIRCUS STARR",
				"PYRAMID SCREEN PRO",
				"T K MAXX",
				"MPS MALE",
				"ASDA",
				"ENVATOMARKET369606",
				"SAINSBURYS S/MKTS",
				"SIMPLY-4-BUSINESS",
				"IZ *neil johnston",
				"WM MORRISONS",
				"DIANOCHE",
				"WWW.ARGOS.CO.UK",
				"MSFX@GEMINI",
				"WWW.BLUEPOLEWEB.CO",
				"HOBBY CRAFT LTD",
				"CO-OP GROUP",
				"KELLY ETHERIDGE",
				"THE GALLERY CAFE",
				"BELLABARISTA.CO.UK",
				"PESTFORCE",
				"BASE HAIR ACADEMY",
				"SQ *JAX FIRST AID",
				"THE RANGE",
				"UK POINT OF SALE",
				"THE CANAL HOUSE",
				"TP WEBSITE PAY",
				"ELITE CLEANING",
				"WWW.WARRINGTON.PAY",
				"CHSE LFPBACS ACCT",
				"ARGOS LTD",
				"VERIF.COM",
				"DICE.FM",
				"LUL TICKET MACHINE",
			},
		},
	}
	return
}
