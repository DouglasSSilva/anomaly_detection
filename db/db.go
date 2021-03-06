package db

// package main

import (
	"fmt"
	"log"
	"strings"

	"bitbucket.org/liamstask/goose/lib/goose"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	// DBConfPath is to be used for dockerizing the APIs
	DBConfPath = "/etc/dbconf"
	// AlternativeDBConfPath to be used out of docker
	AlternativeDBConfPath = "src/anomaly_detection/db/"
)

// Conn persists routes to SQL
var ConnORM *gorm.DB

func ConnectDB(env, mainDBConf, alternateDBConf string) {
	var err error
	// var openstr string
	// driver := "mysql"

	if strings.Compare(env, "") == 0 {
		env = "development"
	}

	fmt.Println(mainDBConf, alternateDBConf)
	dbconf, err := goose.NewDBConf(mainDBConf, env, "")
	if err != nil {
		log.Println("Info: global db configuration file not found. Using local.")
		dbconf, err = goose.NewDBConf(alternateDBConf, env, "")

		if err != nil {
			panic(err)
		}
	}
	// fmt.Printf("\n Driverf:%+v\nEnv:%+v\nOpenstr:%+v\n ", dbconf.Driver.Name, dbconf.Env, dbconf.Driver.OpenStr)
	driver := dbconf.Driver.Name
	openstr := dbconf.Driver.OpenStr

	ConnORM, err = gorm.Open(mysql.New(mysql.Config{
		DriverName:        driver,
		DSN:               openstr, // data source name
		DefaultStringSize: 256,     // default size for string fields

	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Error: mysql ", err)
	}

}
