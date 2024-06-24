package rabbit

import (
	"backapp/db"
	"backapp/messages"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

type OrderMessageHandler interface {
	HandleMessage(*amqp.Delivery, *chan OrderUpdateChannelPayload) error
}

type DefaultOrderMessageHandler struct {};

func (h *DefaultOrderMessageHandler) HandleMessage(delivery *amqp.Delivery, outChannel *chan OrderUpdateChannelPayload) error {
	var rec messages.NewOrder;
	if outChannel == nil {
		panic("Output channel uninitialized")
	}
	err := proto.Unmarshal(delivery.Body, &rec) 
	if err != nil {
		return err;
	}
	log.Printf("INFO Received order, id: %s\temail: %s", rec.Id, rec.Email);
	info, err := db.HandleNewOrder(&rec);
	if err != nil {
		log.Printf("ERROR db.HandleNewOrder failed to save order with id: %s", rec.Id)
	}
	serialized, err := proto.Marshal(&info);
	if err != nil {
		return err
	}
	*outChannel<-OrderUpdateChannelPayload{Payload: serialized, OrderType: OrderTypeNew};
	return nil;
}

