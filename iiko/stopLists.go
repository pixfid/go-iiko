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
