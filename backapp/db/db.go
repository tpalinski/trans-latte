package db

import (
	"backapp/db/model"
	"backapp/messages"
	"backapp/utils"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB;

const TIMEOUT time.Duration = 5;
const RETRIES int = 5;

func getConnectionDetails() (addres, user, password, dbname string) {
	addres = utils.GetEnvWithDefault("POSTGRES_ADDRESS", "localhost")
	user = utils.GetEnvWithDefault("POSTGRES_USER", "postgres")
	password = utils.GetEnvWithDefault("POSTGRES_PASSWORD", "postgres")
	dbname = utils.GetEnvWithDefault("POSTGRES_DB", "translatte")
	return 
}

func InitPostgresConnection() error {
	address, user, password, dbname := getConnectionDetails();
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Warsaw", address, user, password, dbname)
	for range RETRIES {
		database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected successfully to postgres")
			db = database
			db.AutoMigrate(&model.OrderInfo {})
			return nil
		} else {
			log.Println("Could not connect to postgres instance, retrying...")
			time.Sleep(TIMEOUT * time.Second)
		}
	}
	return fmt.Errorf("Ran out of retries")
}

func addOrder(order *model.OrderInfo) {
	db.Create(order)
}

func updateOrder(order *model.OrderInfo) {
	db.Model(order).Updates(order)
}

func getOrder(id string) (order model.OrderInfo){
	db.First(&order, "id = ?", id)
	return
}

// Top level function used for adding order to db
func HandleNewOrder(msg *messages.NewOrder) (info messages.OrderStatusInfo, err error) {
	dbStruct := model.ProtoToOrder(msg);
	addOrder(&dbStruct);
	id := dbStruct.Id;
	fetched := getOrder(id);
	if fetched.Id != msg.Id {
		return info, fmt.Errorf("Error while fetching back ");
	}
	info = model.OrderToProto(&fetched);
	return info, nil
}

func HandleOrderInfo(id string) (info messages.OrderStatusInfo, err error) {
	order := getOrder(id)
	if order.Id == "" {
		return info, fmt.Errorf("No such record in db")
	}
	info = model.OrderToProto(&order)
	return info, nil
}

func HandleOrderPricing(msg *messages.OrderStatusInfo) (*messages.OrderStatusInfo, error) {
	msg.Status = messages.OrderStatus_ORDER_STATUS_PRICED;
	dbStruct := model.InfoToOrder(msg);
	updateOrder(&dbStruct)
	return msg, nil
}
