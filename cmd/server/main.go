package main

import (
	"context"
	"github.com/ashkan-jafarzadeh/delay/api/rest"
	"github.com/ashkan-jafarzadeh/delay/config"
	"github.com/ashkan-jafarzadeh/delay/internal/infra/mysql"
	"github.com/ashkan-jafarzadeh/delay/internal/infra/rabbit"
	mysqlRepo "github.com/ashkan-jafarzadeh/delay/internal/repository/mysql"
	"github.com/ashkan-jafarzadeh/delay/internal/service/assign"
	"github.com/ashkan-jafarzadeh/delay/internal/service/delay"
	"github.com/ashkan-jafarzadeh/delay/internal/service/report"
	"github.com/ashkan-jafarzadeh/delay/pkg/rabbitmq"
	"log"
	"os/signal"
	"syscall"
)

const appName = "snappfood"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.New(appName)
	if err != nil {
		log.Fatal(err)
	}

	db, err := mysql.NewClient(cfg.Mysql)
	if err != nil {
		log.Fatal(err)
	}

	amqpConn, err := rabbit.NewConnection(cfg.Rabbitmq)
	if err != nil {
		log.Fatal(err)
	}

	rMQ := rabbitmq.New(amqpConn, cfg.Rabbitmq)
	err = rMQ.InitChannel()
	if err != nil {
		log.Fatal(err)
	}

	tripRepo := mysqlRepo.NewTrip(db)
	delayRepo := mysqlRepo.NewDelayReport(db)
	orderRepo := mysqlRepo.NewOrder(db)
	agentRepo := mysqlRepo.NewAgent(db)
	vendorRepo := mysqlRepo.NewVendor(db)

	delayService := delay.New(cfg, rMQ, orderRepo, tripRepo, delayRepo)
	assignService := assign.New(cfg, rMQ, agentRepo, delayRepo)
	reportService := report.New(cfg, vendorRepo)

	server := rest.New(cfg, delayService, assignService, reportService)

	//serve http
	if err = server.Serve(ctx); err != nil {
		log.Fatal(err, nil)
	}
}
