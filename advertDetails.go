package main

import (
	"strconv"
	"log"
)

type AdvertDetails struct {
	id int
	name string
	location string
	url string
	price string
	pricePerSqMtr string
	roomsCount string
	sizeInSquareMtr string
	flourNumber string
	buildingType string
	publishDate string
	phone string
	features string
	message string
	createdOn string
}

func PrintDetails(details AdvertDetails) {
	log.Println("name: " + details.name);
	log.Println("location: " + details.location);
	log.Println("url: " + details.url);
	log.Println("price: " + details.price);
	log.Println("price_per_sq_met: " + details.pricePerSqMtr);
	log.Println("rooms count: " + details.roomsCount);
	log.Println("size: " + details.sizeInSquareMtr);
	log.Println("flour: " + details.flourNumber);
	log.Println("type: " + details.buildingType);
	log.Println("publish date: " + details.publishDate);
	log.Println("phone: " + details.phone);
	log.Println("features: " + details.features);
	log.Println("message: " + details.message);
	log.Println("id: " + strconv.Itoa(details.id));
	log.Println("created on: " + details.createdOn);
}