package iiko

import (
	"context"
	"fmt"
)

// Orders Заказы Доставки
type Orders service

// DeliveryHistoryByPhone История доставочных заказов по номеру телефона
// Получение информации о доставочных заказах, которые сделаны с заданного телефона в заданном ресторане.
// Метод работает только для ресторанов с колл-центром (iikoCallCenter)
func (s *Orders) DeliveryHistoryByPhone(ctx context.Context, accessToken, organizationId, phoneNumber string) (*CustomersDeliveryHistoryResponse, error) {
	u := fmt.Sprintf("orders/deliveryHistoryByPhone?access_token=%s&organization=%s&phone=%s", accessToken, organizationId, phoneNumber)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	customersDeliveryHistoryResponse := new(CustomersDeliveryHistoryResponse)
	_, err = s.client.Do(ctx, req, &customersDeliveryHistoryResponse)
	if err != nil {
		return nil, err
	}

	return customersDeliveryHistoryResponse, nil
}

// GetCourierOrders Получить все заказы курьера
func (s *Orders) GetCourierOrders(ctx context.Context, accessToken, organizationId, courierId, requestTimeout string) (*interface{}, error) {
	u := fmt.Sprintf("orders/get_courier_orders?access_token=%s&organization=%s&courier=%s&request_timeout=%s", accessToken, organizationId, courierId, requestTimeout)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	//cityWithStreets := new(CityWithStreets)
	t, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return nil, err
	}

	println(string(t))
	return nil, nil
}
