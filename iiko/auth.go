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
	"strings"
)

// AuthService Аутентификация и авторизация
type AuthService service

// AccessToken Получить маркер доступа апи логина
// Маркер доступа выдается на фиксированный интервал времени. По умолчанию это - 15 минут.
func (s *AuthService) AccessToken(ctx context.Context, userId, userSecret string) (string, error) {
	u := fmt.Sprintf("auth/access_token?user_id=%s&user_secret=%s", userId, userSecret)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return "", err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return "", err
	}

	return strings.Trim(string(resp), `"`), nil
}
