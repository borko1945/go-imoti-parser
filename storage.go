package main

import (
	"log"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func New(file string) Db {
	dbo := InitAndMigrateDb(file);
	return Db{db:dbo};
}

type Db struct {
	db *sql.DB
}

func (this Db) Close() {
	this.db.Close();
}

func (this Db) FindMatch(toMatch *AdvertDetails) []AdvertDetails {
	var result []AdvertDetails;

	rows, err := this.db.Query("SELECT * FROM imotbg where url=? AND name LIKE ? AND phone LIKE ?",
		 toMatch.url, "%"+toMatch.name+"%", "%"+toMatch.phone+"%");
		 if (err != nil) {
			 LogError(err.Error())
			 return result;
		 }

	for rows.Next() {
		var advert AdvertDetails;
		err = rows.Scan(&advert.id,
			&advert.name,
			&advert.location,
			&advert.url,
			&advert.price,
			&advert.pricePerSqMtr,
			&advert.roomsCount,
			&advert.sizeInSquareMtr,
			&advert.flourNumber,
			&advert.buildingType,
			&advert.publishDate,
			&advert.phone,
			&advert.features,
			&advert.message,
			&advert.createdOn);

		if (err != nil) {
			LogError(err.Error())
			continue;
		}

		result = append(result, advert);
	}

	rows.Close();

	return result;
}

func (this Db) findByUrl(url string) []AdvertDetails {
	var result []AdvertDetails;
	sqlStr := "SELECT * FROM imotbg where url='" + url+"'";
	// logsql(sqlStr);
	rows, err := this.db.Query(sqlStr);
	if (err != nil) {
		LogError(err.Error())
		return result;
	}

	for rows.Next() {
		var advert AdvertDetails;
		err = rows.Scan(&advert.id,
			&advert.name,
			&advert.location,
			&advert.url,
			&advert.price,
			&advert.pricePerSqMtr,
			&advert.roomsCount,
			&advert.sizeInSquareMtr,
			&advert.flourNumber,
			&advert.buildingType,
			&advert.publishDate,
			&advert.phone,
			&advert.features,
			&advert.message,
			&advert.createdOn);

			if (err != nil) {
				LogError(err.Error())
				continue;
			}

		result = append(result, advert);
	}

	rows.Close();

	return result;
}

func (this Db) Store(advert *AdvertDetails) {
	log.Println("Storing: " + advert.url);
	stmt, err := this.db.Prepare("INSERT INTO imotbg(name,location,url,price,price_per_sqmtr,rooms_count,size_in_square_mtr,flour_number,"+
		"building_type,publish_date,phone,features,message) values(?,?,?,?,?,?,?,?,?,?,?,?,?)");

		if (err != nil) {
			LogError(err.Error())
			return
		}

	_, err = stmt.Exec(advert.name,
		advert.location,
		advert.url,
		advert.price,
		advert.pricePerSqMtr,
		advert.roomsCount,
		advert.sizeInSquareMtr,
		advert.flourNumber,
		advert.buildingType,
		advert.publishDate,
		advert.phone,
		advert.features,
		advert.message );

		if (err != nil) {
			LogError(err.Error())
		}
}

func logsql(query string) {
	log.Println("SQL: ", query);
}

func InitAndMigrateDb(filepath string) *sql.DB {
	db := InitDB(filepath);
	Migrate(db);
	return db;
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if (err != nil) {
		LogError(err.Error())
	}

	return db
}

func Migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS imotbg(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			location VARCHAR NOT NULL,
			url VARCHAR NOT NULL,
			price VARCHAR NOT NULL,
			price_per_sqmtr VARCHAR,
			rooms_count VARCHAR,
			size_in_square_mtr VARCHAR,
			flour_number VARCHAR,
			building_type VARCHAR,
			publish_date VARCHAR,
			phone VARCHAR,
			features VARCHAR,
			message VARCHAR,
			created_on datetime default current_timestamp
	);
	`

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if (nil != err) {
		LogError(err.Error())
	}
}