package main

import (
	"log"
	"crypto/tls"
	"net/http"
	"fmt"
	"sync"
	// "os"

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

func processDetailsLink(wg *sync.WaitGroup, url string) {
	defer wg.Done();

	fmt.Printf("Processing: %s\n", url)
	doc, err := createDoc(url);
	if err != nil {
		return
	}

	par := newDetailsParser(doc);

	fmt.Println("name: " + par.getName());
	fmt.Println("location: " + par.getLocation());
	fmt.Println("url: " + par.getUrl());
	fmt.Println("price: " + par.getPrice());
	fmt.Println("price_per_sq_met: " + par.getPricePerSqMeter());
	fmt.Println("rooms count: " + par.getRoomsCount());
	fmt.Println("size: " + par.getSizeInSqMeter());
	fmt.Println("flour: " + par.getFlourNumber());
	fmt.Println("type: " + par.getBuildingType());
	fmt.Println("publish date: " + par.getPublishDate());
	fmt.Println("phone: " + par.getPhone());
	fmt.Println("features: " + par.getFeatures());
	fmt.Println("message: " + par.getMessage());
	// os.Exit(0);
}

func main() {
	pageURL := "https://www.imot.bg/pcgi/imot.cgi?act=3&slink=32yu5a&f1=1"
	
	doc, err := createDoc(pageURL);
	if err != nil {
		log.Fatal(err.Error());
	}

	var wg sync.WaitGroup
  
	// Find the review items
	doc.Find(".lnk2").Each(func(i int, s *goquery.Selection) {
	  // For each item found, get the band and title
		linkURL, exists := s.Attr("href")
		if (!exists){
		}
		detailsLink :="https:"+linkURL;

		wg.Add(1)
		go processDetailsLink(&wg, detailsLink);
	})

	wg.Wait();
}
