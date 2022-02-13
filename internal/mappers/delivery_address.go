package mappers

import (
	"github.com/AleksK1NG/es-microservice/internal/dto"
	"github.com/AleksK1NG/es-microservice/internal/order/events/v1"
)

func ChangeDeliveryAddressReqDtoToEventData(reqDto dto.ChangeDeliveryAddressReqDto) v1.OrderChangeDeliveryAddress {
	return v1.OrderChangeDeliveryAddress{
		DeliveryAddress: reqDto.DeliveryAddress,
	}
}
