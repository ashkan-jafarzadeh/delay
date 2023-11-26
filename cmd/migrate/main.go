package main

import (
	"github.com/ashkan-jafarzadeh/delay/config"
	"github.com/ashkan-jafarzadeh/delay/internal/infra/mysql"
	"github.com/ashkan-jafarzadeh/delay/internal/model"
	"log"
)

const appName = "snappfood"

func main() {
	cfg, err := config.New(appName)
	if err != nil {
		log.Fatal(err)
	}

	db, err := mysql.NewClient(cfg.Mysql)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&model.Agent{}, &model.Vendor{}, &model.Order{}, &model.Trip{}, &model.DelayReport{}); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Migrations Done Successfully.")
	}
}
