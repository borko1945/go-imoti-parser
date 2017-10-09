package main

import (
	"log"
	"crypto/tls"
	"net/http"
	"fmt"
	"sync"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

func createDoc(url string ) (*goquery.Document, error){
	charset:= "windows-1251"
	// Load the URL

	var tr = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		DisableKeepAlives: true,
	}

	var netClient = &http.Client{Transport: tr}
	res, err := netClient.Get(url)
	if err != nil {
		log.Fatal(err.Error());
		return nil, err
	}
	defer res.Body.Close()

	// Convert the designated charset HTML to utf-8 encoded HTML.
	// `charset` being one of the charsets known by the iconv package.
	utfBody, err := iconv.NewReader(res.Body, charset, "utf-8")
	if err != nil {
		log.Fatal(err.Error());
		return nil, err
	}

	// use utfBody using goquery
	return goquery.NewDocumentFromReader(utfBody)
}

func processDetailsLink(wg *sync.WaitGroup, url string) (details AdvertDetails, res bool) {
	defer wg.Done();

	fmt.Printf("Processing: %s\n", url)
	doc, err := createDoc(url);
	if err != nil {
		return
	}

	par := newDetailsParser(doc);

	details.name = par.getName();
	details.location = par.getLocation();
	details.url = par.getUrl();
	details.price = par.getPrice();
	details.pricePerSqMtr = par.getPricePerSqMeter();
	details.roomsCount = par.getRoomsCount();
	details.sizeInSquareMtr = par.getSizeInSqMeter();
	details.flourNumber = par.getFlourNumber();
	details.buildingType = par.getBuildingType();
	details.publishDate = par.getPublishDate();
	details.phone = par.getPhone();
	details.features = par.getFeatures();
	details.message = par.getMessage();

	return details, true;
}

func processParsedAdverts(db *Db, adverts []AdvertDetails) {
	for _, element := range adverts {
		if (len(db.FindMatch(&element)) == 0) {
			db.Store(&element);
			sendMail(element)
		} 
	}
}

func main() {
	var wg sync.WaitGroup
	var adsList []AdvertDetails;

	db := New("./imotbg.db");
	defer db.Close();

	for i := 1; i <= Cfg().PagesToParse; i++ {
		pageURL := Cfg().URL + strconv.Itoa(i)

		doc, err := createDoc(pageURL);
		if err != nil {
			log.Fatal(err.Error());
		}

		// Find the review items
		doc.Find(".lnk2").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			linkURL, exists := s.Attr("href")
			if (exists) {
				detailsLink :="https:"+linkURL;
				
				wg.Add(1)
				
				addDetails, valid := processDetailsLink(&wg, detailsLink);
				if (valid) {
					adsList = append(adsList, addDetails);
				}
			}
		})
	}

	wg.Wait();

	log.Println("Processed: " + strconv.Itoa(len(adsList)))

	processParsedAdverts(&db, adsList);
}
