package model

import (
	"backapp/messages"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderInfo struct {
	Id			string
	Email			string
	Description		string
	DateOrdered		time.Time
	StatusDescription	string
	Price			int
	CreatedAt		time.Time
	UpdatedAt		time.Time
}

func ProtoToOrder(msg *messages.NewOrder) OrderInfo {
	return OrderInfo{
		Email: msg.Email,
		Id: msg.Id,
		Description: msg.Description,
		DateOrdered: msg.Date.AsTime(),
		StatusDescription: messages.OrderStatus_ORDER_STATUS_CREATED.String(),
		Price: -1,
	}
}

func OrderToProto(model *OrderInfo) (res messages.OrderStatusInfo) {
	res.ClientDescription = model.Description;
	res.Id = model.Id;
	res.Email = model.Email;
	res.Status = messages.OrderStatus(messages.OrderStatus_value[model.StatusDescription]);
	res.LastUpdated = timestamppb.New(model.UpdatedAt)
	res.DateOrdered = timestamppb.New(model.DateOrdered)
	res.ClientDescription = model.Description;
	price := int64(model.Price);
	if price != -1 {
		res.Price = &price
	}
	return 
}
