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

type CustomersDeliveryHistoryResponse struct {
	CustomersDeliveryHistory []CustomersDeliveryHistory `json:"customersDeliveryHistory"`
}

type CustomersDeliveryHistory struct {
	Customer        *Customer         `json:"customer,omitempty"`
	DeliveryHistory []DeliveryHistory `json:"deliveryHistory"`
}

type Customer struct {
	ID                            *string           `json:"id,omitempty"`
	ExternalID                    *string           `json:"externalId"`
	Name                          *string           `json:"name,omitempty"`
	SurName                       *string           `json:"surName"`
	Nick                          *string           `json:"nick"`
	Comment                       *string           `json:"comment,omitempty"`
	Phone                         *string           `json:"phone,omitempty"`
	AdditionalPhones              []AdditionalPhone `json:"additionalPhones"`
	Addresses                     []interface{}     `json:"addresses"`
	CultureName                   interface{}       `json:"cultureName"`
	Birthday                      interface{}       `json:"birthday"`
	Email                         *string           `json:"email"`
	MiddleName                    interface{}       `json:"middleName"`
	ShouldReceivePromoActionsInfo *bool             `json:"shouldReceivePromoActionsInfo,omitempty"`
	Sex                           *int64            `json:"sex,omitempty"`
	CounteragentInfo              interface{}       `json:"counteragentInfo"`
}

type AdditionalPhone struct {
	Phone *string `json:"phone,omitempty"`
}

type DeliveryHistory struct {
	OrderID        *string            `json:"orderId,omitempty"`
	CustomerID     *string            `json:"customerId,omitempty"`
	OrganizationID *string            `json:"organizationId,omitempty"`
	Sum            *float64           `json:"sum,omitempty"`
	Discount       *float64           `json:"discount,omitempty"`
	Number         *string            `json:"number,omitempty"`
	Phone          *string            `json:"phone,omitempty"`
	IsSelfService  *bool              `json:"isSelfService,omitempty"`
	Address        map[string]*string `json:"address"`
	Date           *string            `json:"date,omitempty"`
	Conception     *string            `json:"conception,omitempty"`
	Comment        *string            `json:"comment"`
	Status         *string            `json:"status,omitempty"`
	Items          []Item             `json:"items"`
	Opinion        *Opinion           `json:"opinion,omitempty"`
}

type Item struct {
	ID               *string     `json:"id,omitempty"`
	OrderItemID      *string     `json:"orderItemId,omitempty"`
	Code             *string     `json:"code"`
	Name             *string     `json:"name,omitempty"`
	Category         interface{} `json:"category"`
	Amount           *int64      `json:"amount,omitempty"`
	Size             interface{} `json:"size"`
	Modifiers        []Item      `json:"modifiers"`
	Sum              *int64      `json:"sum,omitempty"`
	Comment          *string     `json:"comment"`
	GuestID          *string     `json:"guestId,omitempty"`
	ComboInformation interface{} `json:"comboInformation"`
	GroupName        *string     `json:"groupName"`
	GroupID          *string     `json:"groupId,omitempty"`
}

type Opinion struct {
	Organization interface{}   `json:"organization"`
	DeliveryID   *string       `json:"deliveryId,omitempty"`
	Comment      interface{}   `json:"comment"`
	Marks        []interface{} `json:"marks"`
}
