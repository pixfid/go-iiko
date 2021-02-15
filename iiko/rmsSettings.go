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

// RmsSettings Настройки ресторана
type RmsSettings service

// OrganizationUserResponse Список всех сотрудников организации
type OrganizationUserResponse struct {
	Users []User `json:"users"`
}

type User struct {
	ID               *string     `json:"id,omitempty"`
	FirstName        *string     `json:"firstName"`
	MiddleName       *string     `json:"middleName"`
	LastName         *string     `json:"lastName"`
	DisplayName      *string     `json:"displayName,omitempty"`
	Phone            *string     `json:"phone"`
	CellPhone        *string     `json:"cellPhone"`
	Email            interface{} `json:"email"`
	Code             *string     `json:"code,omitempty"`
	PinCode          *string     `json:"pinCode,omitempty"`
	Note             interface{} `json:"note"`
	MainRole         *Role       `json:"mainRole,omitempty"`
	Roles            []Role      `json:"roles"`
	Deleted          *bool       `json:"deleted,omitempty"`
	ExternalRevision *int64      `json:"externalRevision,omitempty"`
}

type Role struct {
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
	Code             *string  `json:"code,omitempty"`
	PaymentPerHour   *float64 `json:"paymentPerHour,omitempty"`
	SteadySalary     *float64 `json:"steadySalary,omitempty"`
	Comment          *string  `json:"comment,omitempty"`
	ScheduleType     *string  `json:"scheduleType,omitempty"`
	ExternalRevision *int64   `json:"externalRevision,omitempty"`
	Deleted          *bool    `json:"deleted,omitempty"`
}

// GetCouriers Получить список курьеров организации
// Возвращает список всех сотрудников, которые являются курьерами доставки в заданном ресторане.
// Курьером доставки считается пользователь, обладающий правом “D_DCO” (быть курьером доставки).
func (s *RmsSettings) GetCouriers(ctx context.Context, accessToken, organizationId string) (*OrganizationUserResponse, error) {
	u := fmt.Sprintf("rmsSettings/getCouriers?access_token=%s&organization=%s", accessToken, organizationId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	organizationUserResponse := new(OrganizationUserResponse)
	_, err = s.client.Do(ctx, req, &organizationUserResponse)
	if err != nil {
		return nil, err
	}

	return organizationUserResponse, nil
}

// PaymentTypesResponse Внешние типы оплат
type PaymentTypesResponse struct {
	PaymentTypes []PaymentType `json:"paymentTypes"`
}

// GetPaymentTypes Получить список типов оплат
// Запрос возвращает список внешних типов оплат для заданного ресторана.
// Внешними считаются типы, процессинг которых идет не на стороне iiko.
func (s *RmsSettings) GetPaymentTypes(ctx context.Context, accessToken, organizationId string) (*PaymentTypesResponse, error) {
	u := fmt.Sprintf("rmsSettings/getPaymentTypes?access_token=%s&organization=%s", accessToken, organizationId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	paymentTypesResponse := new(PaymentTypesResponse)
	_, err = s.client.Do(ctx, req, &paymentTypesResponse)
	if err != nil {
		return nil, err
	}

	return paymentTypesResponse, nil
}
