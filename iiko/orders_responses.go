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

type DeliveryOrdersResponse struct {
	DeliveryOrders []DeliveryOrder `json:"deliveryOrders"`
}
type Customer struct {
	Sex                           *int64        `json:"sex"`
	ID                            *string       `json:"id"`
	ExternalID                    *string       `json:"externalId"`
	Name                          *string       `json:"name"`
	SurName                       *string       `json:"surName"`
	Nick                          *string       `json:"nick"`
	Comment                       *string       `json:"comment"`
	Phone                         *string       `json:"phone"`
	AdditionalPhones              []interface{} `json:"additionalPhones"`
	Addresses                     []interface{} `json:"addresses"`
	CultureName                   interface{}   `json:"cultureName"`
	Birthday                      interface{}   `json:"birthday"`
	Email                         interface{}   `json:"email"`
	MiddleName                    interface{}   `json:"middleName"`
	ShouldReceivePromoActionsInfo *bool         `json:"shouldReceivePromoActionsInfo"`
	CounteragentInfo              interface{}   `json:"counteragentInfo"`
}
type Address struct {
	City                  *string     `json:"city"`
	Street                *string     `json:"street"`
	StreetID              *string     `json:"streetId"`
	StreetClassifierID    *string     `json:"streetClassifierId"`
	Index                 *string     `json:"index"`
	Home                  *string     `json:"home"`
	Housing               *string     `json:"housing"`
	Apartment             *string     `json:"apartment"`
	Entrance              *string     `json:"entrance"`
	Floor                 *string     `json:"floor"`
	Doorphone             *string     `json:"doorphone"`
	Comment               *string     `json:"comment"`
	RegionID              interface{} `json:"regionId"`
	ExternalCartographyID *string     `json:"externalCartographyId"`
}
type CourierInfo struct {
	CourierID *string     `json:"courierId"`
	Location  interface{} `json:"location"`
}
type MainRole struct {
	ID               *string  `json:"id"`
	Name             *string  `json:"name"`
	Code             *string  `json:"code"`
	PaymentPerHour   *float64 `json:"paymentPerHour"`
	SteadySalary     *float64 `json:"steadySalary"`
	Comment          *string  `json:"comment"`
	ScheduleType     *string  `json:"scheduleType"`
	ExternalRevision *int64   `json:"externalRevision"`
	Deleted          *bool    `json:"deleted"`
}
type Roles struct {
	ID               *string  `json:"id"`
	Name             *string  `json:"name"`
	Code             *string  `json:"code"`
	PaymentPerHour   *float64 `json:"paymentPerHour"`
	SteadySalary     *float64 `json:"steadySalary"`
	Comment          *string  `json:"comment"`
	ScheduleType     *string  `json:"scheduleType"`
	ExternalRevision *int64   `json:"externalRevision"`
	Deleted          *bool    `json:"deleted"`
}
type Operator struct {
	ID               *string     `json:"id"`
	FirstName        *string     `json:"firstName"`
	MiddleName       interface{} `json:"middleName"`
	LastName         *string     `json:"lastName"`
	DisplayName      *string     `json:"displayName"`
	Phone            interface{} `json:"phone"`
	CellPhone        interface{} `json:"cellPhone"`
	Email            interface{} `json:"email"`
	Code             *string     `json:"code"`
	PinCode          *string     `json:"pinCode"`
	Note             interface{} `json:"note"`
	MainRole         MainRole    `json:"mainRole"`
	Roles            []Roles     `json:"roles"`
	Deleted          *bool       `json:"deleted"`
	ExternalRevision *int64      `json:"externalRevision"`
}
type Modifiers struct {
	ID               *string     `json:"id"`
	OrderItemID      *string     `json:"orderItemId"`
	Name             *string     `json:"name"`
	Amount           *float64    `json:"amount"`
	GroupID          *string     `json:"groupId"`
	GroupName        *string     `json:"groupName"`
	Size             interface{} `json:"size"`
	ComboInformation interface{} `json:"comboInformation"`
	Sum              *float64    `json:"sum"`
	Code             interface{} `json:"code"`
	Category         interface{} `json:"category"`
}
type Items struct {
	ID               *string     `json:"id"`
	OrderItemID      *string     `json:"orderItemId"`
	Code             *string     `json:"code"`
	Name             *string     `json:"name"`
	Category         interface{} `json:"category"`
	Amount           *float64    `json:"amount"`
	Size             interface{} `json:"size"`
	Modifiers        []Modifiers `json:"modifiers"`
	Sum              *float64    `json:"sum"`
	Comment          interface{} `json:"comment"`
	GuestID          *string     `json:"guestId"`
	ComboInformation interface{} `json:"comboInformation"`
}
type Guests struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}
type PaymentType struct {
	ID                           *string     `json:"id"`
	Code                         *string     `json:"code"`
	Name                         *string     `json:"name"`
	Comment                      *string     `json:"comment"`
	Combinable                   *bool       `json:"combinable"`
	ExternalRevision             *int64      `json:"externalRevision"`
	ApplicableMarketingCampaigns interface{} `json:"applicableMarketingCampaigns"`
	Deleted                      *bool       `json:"deleted"`
}
type Payments struct {
	Sum                     *float64    `json:"sum"`
	PaymentType             PaymentType `json:"paymentType"`
	AdditionalData          interface{} `json:"additionalData"`
	IsProcessedExternally   *bool       `json:"isProcessedExternally"`
	IsPreliminary           *bool       `json:"isPreliminary"`
	IsExternal              *bool       `json:"isExternal"`
	ChequeAdditionalInfo    interface{} `json:"chequeAdditionalInfo"`
	OrganizationDetailsInfo interface{} `json:"organizationDetailsInfo"`
	IsFiscalizedExternally  *bool       `json:"isFiscalizedExternally"`
}
type OrderType struct {
	ID               *string `json:"id"`
	Name             *string `json:"name"`
	OrderServiceType *string `json:"orderServiceType"`
	ExternalRevision *int64  `json:"externalRevision"`
}
type DeliveryTerminal struct {
	DeliveryTerminalID   *string     `json:"deliveryTerminalId"`
	CrmID                *string     `json:"crmId"`
	RestaurantName       *string     `json:"restaurantName"`
	GroupName            interface{} `json:"groupName"`
	ExternalRevision     *int64      `json:"externalRevision"`
	TechnicalInformation interface{} `json:"technicalInformation"`
	Address              *string     `json:"address"`
	ProtocolVersion      *int64      `json:"protocolVersion"`
}
type Discounts struct {
	DiscountCardTypeID    *string     `json:"discountCardTypeId"`
	DiscountCardSlip      interface{} `json:"discountCardSlip"`
	DiscountOrIncreaseSum *float64    `json:"discountOrIncreaseSum"`
}
type Opinion struct {
	Organization interface{}   `json:"organization"`
	DeliveryID   *string       `json:"deliveryId"`
	Comment      interface{}   `json:"comment"`
	Marks        []interface{} `json:"marks"`
}
type DeliveryOrder struct {
	Customer              Customer         `json:"customer"`
	OrderID               *string          `json:"orderId"`
	CustomerID            *string          `json:"customerId"`
	MovedDeliveryID       interface{}      `json:"movedDeliveryId"`
	CustomerName          *string          `json:"customerName"`
	CustomerPhone         interface{}      `json:"customerPhone"`
	Address               Address          `json:"address"`
	RestaurantID          *string          `json:"restaurantId"`
	Organization          *string          `json:"organization"`
	Sum                   *float64         `json:"sum"`
	Discount              *float64         `json:"discount"`
	Number                *string          `json:"number"`
	Status                *string          `json:"status"`
	StatusCode            *string          `json:"statusCode"`
	DeliveryCancelCause   interface{}      `json:"deliveryCancelCause"`
	DeliveryCancelComment interface{}      `json:"deliveryCancelComment"`
	CourierInfo           CourierInfo      `json:"courierInfo"`
	OrderLocationInfo     interface{}      `json:"orderLocationInfo"`
	DeliveryDate          *string          `json:"deliveryDate"`
	ActualTime            *string          `json:"actualTime"`
	BillTime              *string          `json:"billTime"`
	CancelTime            interface{}      `json:"cancelTime"`
	CloseTime             interface{}      `json:"closeTime"`
	ConfirmTime           *string          `json:"confirmTime"`
	CreatedTime           *string          `json:"createdTime"`
	PrintTime             *string          `json:"printTime"`
	SendTime              *string          `json:"sendTime"`
	Comment               *string          `json:"comment"`
	Problem               interface{}      `json:"problem"`
	Operator              Operator         `json:"operator"`
	Conception            interface{}      `json:"conception"`
	MarketingSource       interface{}      `json:"marketingSource"`
	DurationInMinutes     *int64           `json:"durationInMinutes"`
	PersonsCount          *int64           `json:"personsCount"`
	SplitBetweenPersons   *bool            `json:"splitBetweenPersons"`
	Items                 []Items          `json:"items"`
	Guests                []Guests         `json:"guests"`
	Payments              []Payments       `json:"payments"`
	OrderType             OrderType        `json:"orderType"`
	DeliveryTerminal      DeliveryTerminal `json:"deliveryTerminal"`
	Discounts             []Discounts      `json:"discounts"`
	IikoCard5Coupon       interface{}      `json:"iikoCard5Coupon"`
	CustomData            interface{}      `json:"customData"`
	Opinion               Opinion          `json:"opinion"`
	ReferrerID            interface{}      `json:"referrerId"`
}
