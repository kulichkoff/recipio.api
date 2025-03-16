package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"recip.io/api/internal/common"
)

type AuthService interface {
	Authorize(username, password string) (string, error)
}

type service struct {
	logger *zap.Logger
}

func NewService() AuthService {
	logger := common.GetLogger()
	return &service{
		logger: logger,
	}
}

func (s *service) Authorize(username, password string) (string, error) {
	// TODO use UserService and bcrypt
	if username != "test" && password != "tset" {
		return "", ErrUnauthorized
	}

	claims := &JWTUserClaims{
		Name:  username,
		Email: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// TODO get secret from config
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		s.logger.Error("Failed to sign token", zap.Error(err))
		return "", ErrSignIncompleted
	}

	return t, nil
}
