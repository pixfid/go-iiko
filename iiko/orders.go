/*
 * MIT License
 *
 * Copyright (c) 2020 Semchenko Aleksandr
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package iiko

import (
	"context"
	"fmt"
)

// Orders Заказы Доставки
type Orders service

func (s *Orders) Info(ctx context.Context, accessToken, organizationId, orderId string) (*interface{}, error) {
	u := fmt.Sprintf("orders/info?access_token=%s&organization=%s&order=%s", accessToken, organizationId, orderId)
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

type CheckAddressParams struct {
	City   string `json:"city,omitempty"`
	Street string `json:"street,omitempty"`
	Home   string `json:"home,omitempty"`
}

type CheckAddressResponse struct {
	AddressInZone *bool `json:"addressInZone,omitempty"`
}

// СheckAddress Проверить осуществимость доставки по указанному адресу
// Определяет, возможно ли осуществить доставку по указанному адресу.
func (s *Orders) СheckAddress(ctx context.Context, accessToken, requestTimeout, organizationId string, params CheckAddressParams) (*CheckAddressResponse, error) {
	u := fmt.Sprintf("orders/checkAddress?access_token=%s&request_timeout=%s&organizationId=%s", accessToken, requestTimeout, organizationId)
	req, err := s.client.NewRequest("POST", u, params)
	if err != nil {
		return nil, err
	}

	checkAddressResponse := new(CheckAddressResponse)
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return nil, err
	}

	return checkAddressResponse, nil
}

//TODO
func (s *Orders) DeliveryOrders(ctx context.Context, accessToken, organizationId, dateFrom, dateTo, deliveryStatus, deliveryTerminalId, requestTimeout string) (*interface{}, error) {
	u := fmt.Sprintf("orders/deliveryOrders?access_token=%s&organization=%s&dateFrom=%s&dateTo=%s&deliveryStatus=%s&deliveryTerminalId=%s&request_timeout=%s",
		accessToken, organizationId, dateFrom, dateTo, deliveryStatus, deliveryTerminalId, requestTimeout)
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
