package main

import (
	"log"
	"net/smtp"
)

func endl(str string) string {
	return str + "\r\n";
}

func sendMail(add *AdvertDetails) {
	log.Println("Sending mail");

	from := Cfg().Email.From;
	pass := Cfg().Email.Pass;
	to := Cfg().Email.To;
	
	message := ""
	message += "Subject: " + "Imotbg: Цена:" + add.price + " . " + add.sizeInSquareMtr + " . " + add.location + " . " + add.roomsCount + "\r\n"
	message += "\r\n"
	message += endl(add.url);
	message += endl(add.name);
	message += endl("Цена: " + add.price);
	message += endl("Цена на кв.м:: " + add.pricePerSqMtr);
	message += endl("Квадратура: " + add.sizeInSquareMtr);
	message += endl("Етаж: " + add.flourNumber);
	message += endl("Тип апартамент: " + add.buildingType);
	message += endl("Публикувана на: " + add.publishDate);
	message += endl("Детайли: " + add.features)
	message += "\r\n"
	message += endl(add.message);

	if (Cfg().Email.Simulate) {
		return
	}

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, to, []byte(message))

	if err != nil {
		LogError("smtp error: " + err.Error())
		return
	}
}
