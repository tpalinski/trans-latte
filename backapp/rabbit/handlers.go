package rabbit

import (
	"backapp/messages"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

type OrderMessageHandler interface {
	HandleMessage(*amqp.Delivery) error
}

type DefaultOrderMessageHandler struct {};

func (h *DefaultOrderMessageHandler) HandleMessage(delivery *amqp.Delivery) error {
	var rec messages.NewOrder;
	err := proto.Unmarshal(delivery.Body, &rec) 
	if err != nil {
		return err;
	}
	//TODO - do some stuff with the data. Just console log them for now
	log.Printf("Received order, id: %s\temail: %s", rec.Id, rec.Email);
	return nil;
}
