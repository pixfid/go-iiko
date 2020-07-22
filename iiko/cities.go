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

// CitiesService Города, улицы, регионы
type CitiesService service

type CityWithStreets []CityWithStreet

type CityWithStreet struct {
	City    *City  `json:"city,omitempty"`
	Streets []City `json:"streets"`
}

type City struct {
	AdditionalInfo   interface{} `json:"additionalInfo"`
	ClassifierID     *string     `json:"classifierId"`
	Deleted          *bool       `json:"deleted,omitempty"`
	ExternalRevision *int64      `json:"externalRevision,omitempty"`
	ID               *string     `json:"id,omitempty"`
	Name             *string     `json:"name,omitempty"`
	CityID           *string     `json:"cityId,omitempty"`
}

// Cities Получение списка городов и их улиц
// Метод возвращает список всех городов и улиц каждого из городов.
// Эти данные могут быть использовать для задания адреса доставки.
func (s *CitiesService) Cities(ctx context.Context, accessToken, organizationId string) (*CityWithStreets, error) {
	u := fmt.Sprintf("cities/cities?access_token=%s&organization=%s", accessToken, organizationId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	cityWithStreets := new(CityWithStreets)
	_, err = s.client.Do(ctx, req, &cityWithStreets)
	if err != nil {
		return nil, err
	}

	return cityWithStreets, nil
}

type Cities []City

// CitiesList Получение списка городов организации
// Метод возвращает список всех городов заданной организации.
// Эти данные могут быть использовать для задания адреса доставки.
func (s *CitiesService) CitiesList(ctx context.Context, accessToken, organizationId string) (*Cities, error) {
	u := fmt.Sprintf("cities/citiesList?access_token=%s&organization=%s", accessToken, organizationId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	cities := new(Cities)
	_, err = s.client.Do(ctx, req, &cities)
	if err != nil {
		return nil, err
	}

	return cities, nil
}

type Streets []City

// Streets Получение списка улиц города заданной организации
// Метод возвращает список всех городов заданной организации.
// Эти данные могут быть использовать для задания адреса доставки.
func (s *CitiesService) Streets(ctx context.Context, accessToken, organizationId, cityId string) (*Streets, error) {
	u := fmt.Sprintf("streets/streets?access_token=%s&organization=%s&city=%s", accessToken, organizationId, cityId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	streets := new(Streets)
	_, err = s.client.Do(ctx, req, &streets)
	if err != nil {
		return nil, err
	}
	return streets, nil
}

type RegionsList struct {
	Regions []interface{} `json:"regions"` // TODO
}

// Regions Получение списка регионов
// Метод возвращает список всех всех регионов, которые есть в справочнике регионов организации.
// Эти данные могут быть использовать для задания региона в адресе доставки.
func (s *CitiesService) Regions(ctx context.Context, accessToken, organizationId string) (*RegionsList, error) {
	u := fmt.Sprintf("regions/regions?access_token=%s&organization=%s", accessToken, organizationId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	regionsList := new(RegionsList)
	_, err = s.client.Do(ctx, req, &regionsList)
	if err != nil {
		return nil, err
	}

	return regionsList, nil
}
