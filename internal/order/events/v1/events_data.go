package v1

import (
	"github.com/AleksK1NG/es-microservice/internal/order/models"
	"time"
)

type OrderCreatedEventData struct {
	ShopItems       []*models.ShopItem `json:"shopItems" bson:"shopItems,omitempty" validate:"required"`
	AccountEmail    string             `json:"accountEmail" bson:"accountEmail,omitempty" validate:"required,email"`
	DeliveryAddress string             `json:"deliveryAddress" bson:"deliveryAddress,omitempty" validate:"required"`
}

type OrderUpdatedEventData struct {
	ShopItems []*models.ShopItem `json:"shopItems" bson:"shopItems,omitempty" validate:"required"`
}

type OrderCanceledEventData struct {
	CancelReason string `json:"cancelReason" validate:"required"`
}

type OrderDeliveredEventData struct {
	DeliveryTimestamp time.Time `json:"deliveryTimestamp" validate:"required"`
}

type OrderChangeDeliveryAddress struct {
	DeliveryAddress string `json:"deliveryAddress" bson:"deliveryAddress,omitempty" validate:"required"`
}