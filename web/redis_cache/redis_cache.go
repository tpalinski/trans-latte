// A self updating cache used for storing order info
package rediscache

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"web/messages"
	"web/rabbit"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

var client redis.Client;
var ctx context.Context = context.Background();

func InitializeRedisConnection() {
	redisAddress, ok := os.LookupEnv("REDIS_ADDRESS");
	if !ok {
		redisAddress = "localhost:6379"
	}
	client = *redis.NewClient(&redis.Options{
		Addr: redisAddress,
		Password: "",
		DB: 0,
	})
	go func() {
		for info := range rabbit.OrderUpdateChannel {
			go updateOrder(info)
		}
	}()
}

func GetOrderInfo(id uuid.UUID) (orderInfo messages.OrderStatusInfo, err error) {
	binaryData, err := client.Get(ctx, id.String()).Result();
	if err == redis.Nil { // no key
		res := handleMissingKey(id)
		if res == nil {
			return orderInfo, fmt.Errorf("no order with supplied uuid")
		}
		orderInfo = *res;
	} else if err != nil {
		return orderInfo, err;
	} else {
		err = proto.Unmarshal([]byte(binaryData), &orderInfo);
		log.Printf("INFO Cache hit for order: %s", orderInfo.Id)
		if err != nil {
			return orderInfo, err;
		}
	}
	return orderInfo, nil;
}

func updateOrder(order *messages.OrderStatusInfo) error {
	data, err := proto.Marshal(order);
	if err != nil {
		return err;
	}
	client.Set(ctx, order.Id, data, time.Duration(1) * time.Minute)
	return nil;
}

func handleMissingKey(id uuid.UUID) (orderInfo *messages.OrderStatusInfo) {
	requestOrderData(id)
	const timeoutMs time.Duration = 100;
	const retries int = 5;
	for range retries {
		binaryData, err := client.Get(ctx, id.String()).Result();
		if err == redis.Nil {
			time.Sleep(timeoutMs * time.Millisecond)
		} else {
			orderInfo = &messages.OrderStatusInfo{}
			proto.Unmarshal([]byte(binaryData), orderInfo);
			log.Printf("INFO Successfully fetched order info for order: %s", orderInfo.Id)
			return
		}
	}
	return nil
}

func requestOrderData(id uuid.UUID) {
	info := messages.GetOrder{Id: id.String()}
	err := rabbit.SendOrderUpdateRequest(&info)
	if err != nil {
		log.Printf("ERROR: redis_cache.requestOrderData - failed for order: %s", id.String())
	}
}
