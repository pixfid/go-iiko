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
