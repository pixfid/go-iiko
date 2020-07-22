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

// StopLists Стоп-листы
type StopLists service

type StopList struct {
	StopList                  []interface{} `json:"stopList"`
	UnregisteredOrganizations []interface{} `json:"unregisteredOrganizations"`
}

// GetDeliveryStopList Получить стоп-лист по сети ресторанов
// Запрос возвращает список продуктов, находящихся в стоп-листе.
// В случае запроса на колл-центра в результате могут находиться позиции стоп-листа из других ресторанов.
func (s *StopLists) GetDeliveryStopList(ctx context.Context, accessToken, organizationId string) (*StopList, error) {
	u := fmt.Sprintf("stopLists/getDeliveryStopList?access_token=%s&organization=%s", accessToken, organizationId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	stopList := new(StopList)
	_, err = s.client.Do(ctx, req, &stopList)
	if err != nil {
		return nil, err
	}

	return stopList, nil
}
