package main

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func newDetailsParser(doc *goquery.Document) *DetailsParser{
	detailsParser := new(DetailsParser)
	detailsParser.doc = doc;

	return detailsParser;
}

type DetailsParser struct { 
	doc *goquery.Document
}

func (this *DetailsParser) getBuildingType() string {
	return trim(this.getSelectorText("body > div:nth-child(4) > table:nth-child(7) > tbody > tr:nth-child(1) > td:nth-child(1) > form:nth-child(3) > table:nth-child(72) > tbody > tr:nth-child(1) > td:nth-child(3) > table:nth-child(6) > tbody > tr:nth-child(5) > td:nth-child(2) > b"))
}

func (this *DetailsParser) getFlourNumber() string {
	return trim(this.getSelectorText("body > div:nth-child(4) > table:nth-child(7) > tbody > tr:nth-child(1) > td:nth-child(1) > form:nth-child(3) > table:nth-child(72) > tbody > tr:nth-child(1) > td:nth-child(3) > table:nth-child(6) > tbody > tr:nth-child(2) > td:nth-child(2) > b"))
}

func (this *DetailsParser) getName() string {
	return trim(this.getSelectorText("body > div:nth-child(4) > table:nth-child(7) > tbody > tr:nth-child(1) > td:nth-child(1) > form:nth-child(3) > table:nth-child(72) > tbody > tr:nth-child(1) > td:nth-child(3) > span:nth-child(4)"))
}

func (this *DetailsParser) getMessage() string {
	return trim(trim(this.getSelectorText("#description_div")))
}

func (this *DetailsParser) getPrice() string {
	return trim(this.getSelectorText("#cena nobr"))
}

func (this *DetailsParser) getPublishDate() string {
	return trim(this.getSelectorText("body > div:nth-child(4) > table:nth-child(7) > tbody > tr:nth-child(1) > td:nth-child(1) > form:nth-child(3) > table:nth-child(72) > tbody > tr:nth-child(2) > td > div > span"))
}

func (this *DetailsParser) getFeatures() string {
	// var features []string;
	var result string;
	this.getSelection("body > div:nth-child(4) > table:nth-child(7) > tbody > tr:nth-child(1) > td:nth-child(1) > form:nth-child(3) > table:nth-child(77) > tbody > tr > td > div").
	Each(func(i int, s *goquery.Selection) {
		// features = append(features, s.Text())
		// fmt.Println(s.Text())
		result += " " + trim(s.Text());
	});

	return result;
}

func trim(str string) string {
	return strings.Trim(str," ");
}

func (this *DetailsParser) getRoomsCount() string {
	return trim(this.getSelectorText("body > div:nth-child(4) > table:nth-child(7) > tbody > tr:nth-child(1) > td:nth-child(1) > form:nth-child(3) > table:nth-child(72) > tbody > tr:nth-child(1) > td:nth-child(3) > span:nth-child(1)"))
}

func (this *DetailsParser) getPricePerSqMeter() string{
	return trim(this.getSelectorText("#cenakv"))
}

func (this *DetailsParser) getSizeInSqMeter() string{
	return trim(this.getSelectorText("body > div:nth-child(4) > table:nth-child(7) > tbody > tr:nth-child(1) > td:nth-child(1) > form:nth-child(3) > table:nth-child(72) > tbody > tr:nth-child(1) > td:nth-child(3) > table:nth-child(6) > tbody > tr:nth-child(1) > td:nth-child(2) > b"))
}

func (this *DetailsParser) getSelectorText(selector string) string {
	return this.getSelection(selector).First().Text();
}

func (this *DetailsParser) getSelection(selector string) *goquery.Selection {
	return this.doc.Find(selector);
}

func (this* DetailsParser) getPhone() string {
	return trim(this.getSelectorText("body > div:nth-child(4) > table:nth-child(7) > tbody > tr:nth-child(1) > td:nth-child(1) > form:nth-child(3) > div:nth-child(79) > span > span:nth-child(3)"))
}

func (this* DetailsParser) getLocation() string {
	return trim(this.getSelectorText("body > div:nth-child(4) > table:nth-child(7) > tbody > tr:nth-child(2) > td > div:nth-child(2) > h2 > b"))
}

func (this* DetailsParser) getUrl() string {
	url, found := this.getSelection("body > div:nth-child(4) > table:nth-child(7) > tbody > tr:nth-child(1) > td:nth-child(1) > form:nth-child(3) > table:nth-child(72) > tbody > tr:nth-child(2) > td > div > div:nth-child(5) > input").First().Attr("value")
	if (!found) {
		return "";
	}

	return trim(url);
}
