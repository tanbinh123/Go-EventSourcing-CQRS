package events

import (
	"github.com/AleksK1NG/es-microservice/internal/models"
	"github.com/AleksK1NG/es-microservice/pkg/es"
)

const (
	OrderCreated    = "ORDER_CREATED"
	OrderPaid       = "ORDER_PAID"
	OrderSubmitted  = "ORDER_SUBMITTED"
	OrderDelivering = "ORDER_DELIVERING"
	OrderDelivered  = "ORDER_DELIVERED"
	OrderCanceled   = "ORDER_CANCELED"
	OrderUpdated    = "ORDER_UPDATED"
)

type OrderCreatedData struct {
	ShopItems    []*models.ShopItem `json:"shopItems" bson:"shopItems,omitempty" validate:"required"`
	AccountEmail string             `json:"accountEmail" bson:"accountEmail,omitempty" validate:"required,email"`
}

type OrderUpdatedData struct {
	ShopItems []*models.ShopItem `json:"shopItems" bson:"shopItems,omitempty" validate:"required"`
}

func NewCreateOrderEvent(aggregate es.Aggregate, data []byte) es.Event {
	createOrderEvent := es.NewBaseEvent(aggregate, OrderCreated)
	createOrderEvent.SetData(data)
	return createOrderEvent
}

func NewPayOrderEvent(aggregate es.Aggregate, data []byte) es.Event {
	orderPaidEvent := es.NewBaseEvent(aggregate, OrderPaid)
	orderPaidEvent.SetData(data)
	return orderPaidEvent
}

func NewSubmitOrderEvent(aggregate es.Aggregate) es.Event {
	return es.NewBaseEvent(aggregate, OrderSubmitted)
}

func NewOrderUpdatedEvent(aggregate es.Aggregate, data []byte) es.Event {
	orderUpdatedEvent := es.NewBaseEvent(aggregate, OrderUpdated)
	orderUpdatedEvent.SetData(data)
	return orderUpdatedEvent
}