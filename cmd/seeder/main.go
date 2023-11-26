package main

import (
	"context"
	"fmt"
	"github.com/ashkan-jafarzadeh/delay/config"
	"github.com/ashkan-jafarzadeh/delay/internal/infra/mysql"
	"github.com/ashkan-jafarzadeh/delay/internal/model"
	"github.com/ashkan-jafarzadeh/delay/internal/repository"
	mysqlRepo "github.com/ashkan-jafarzadeh/delay/internal/repository/mysql"
	"github.com/ashkan-jafarzadeh/delay/pkg/faker"
	"log"
	"math/rand"
)

const appName = "snappfood"

const agentsCount = 20
const vendorsCount = 20
const ordersCountPerVendor = 10

// ask to seed dummy data for instruments and trades
func main() {
	cfg, err := config.New(appName)
	if err != nil {
		log.Fatal(err)
	}

	db, err := mysql.NewClient(cfg.Mysql)
	if err != nil {
		return
	}

	ctx := context.Background()

	seedAgents(ctx, mysqlRepo.NewAgent(db), agentsCount)
	vendors := seedVendors(ctx, mysqlRepo.NewVendor(db), vendorsCount)

	for _, vendor := range vendors {
		seedOrders(ctx, mysqlRepo.NewOrder(db), mysqlRepo.NewTrip(db), vendor.ID, ordersCountPerVendor)
	}

	log.Println("Seeders Done Successfully.")
}

func seedOrders(ctx context.Context, repo repository.Order, tripRepo repository.Trip, vendorId uint, count int) {
	for i := 0; i < count; i++ {
		order := model.Order{
			VendorId:     vendorId,
			Price:        faker.RandFloat(1000, 50000),
			DeliveryTime: uint(faker.RandInt(5, 25)),
		}

		order.CreatedAt = faker.RandPastTime(0, 120)
		order.UpdatedAt = order.CreatedAt

		err := repo.Create(ctx, &order)
		if err != nil {
			log.Println("Order seed failed")
			continue
		}

		if faker.RandBool() {
			seedTrip(ctx, tripRepo, order.ID)
		}
	}
}

func seedTrip(ctx context.Context, repo repository.Trip, orderId uint) {
	var trip = model.Trip{
		OrderId: orderId,
		Status:  model.TripStatuses()[rand.Intn(3)],
	}

	err := repo.Create(ctx, &trip)
	if err != nil {
		fmt.Println("Trip seed failed")
	}
}

func seedVendors(ctx context.Context, repo repository.Vendor, count int) []model.Vendor {
	var vendors []model.Vendor
	for i := 0; i < count; i++ {
		vendor := model.Vendor{
			Name: fmt.Sprintf("Vendor %d", i),
		}

		err := repo.Create(ctx, &vendor)
		if err != nil {
			fmt.Println("Vendor seed failed")
			continue
		}

		vendors = append(vendors, vendor)
	}

	return vendors
}

func seedAgents(ctx context.Context, repo repository.Agent, count int) {
	for i := 0; i < count; i++ {
		agent := model.Agent{
			Name: fmt.Sprintf("Agent %d", i),
		}

		err := repo.Create(ctx, &agent)
		if err != nil {
			log.Println("Agent seed failed")
		}
	}
}
