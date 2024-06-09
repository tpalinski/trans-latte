// A self updating cache used for storing order info
package rediscache

import (
	"context"
	"os"
	"time"
	"web/messages"

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
}

func GetOrderInfo(id uuid.UUID) (orderInfo messages.OrderStatusInfo, err error) {
	binaryData, err := client.Get(ctx, id.String()).Result();
	if err == redis.Nil { // no key
		// TODO - fetch missing data from backend service. For now, returns mock object
		orderInfo = messages.OrderStatusInfo{Id: id.String(), Email: "mock@mail.com", ClientDescription: "Mock order description", StatusDescription: "Order submitted", DateOrdered: time.DateTime, LastUpdated: time.Now().String()}
	} else if err != nil {
		return orderInfo, err;
	} else {
		err = proto.Unmarshal([]byte(binaryData), &orderInfo);
		if err != nil {
			return orderInfo, err;
		}
	}
	return orderInfo, nil;
}
