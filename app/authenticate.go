package app

import (
	"errors"
	"fmt"

	"github.com/ahmadayub792/twitter-sample-server/model"
	"github.com/ahmadayub792/twitter-sample-server/store"
	"github.com/dgrijalva/jwt-go"
)

func (app *App) GenerateToken(email string, password string) (string, error) {
	user, err := app.UserStore.FindOne(model.User{
		Email: email,
	})

	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return "", ErrAuthFailed
		}
		return "", err
	}

	if err := app.PasswordHasher.ValidateHashPassword(user.Password, password); err != nil {
		return "", ErrIncorrectPassword
	}

	app.setUserSession(user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
	})

	return token.SignedString(app.TokenSecret)
}

// VerifyToken checks token validity
func (app *App) VerifyToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrAuthFailed
		}
		return app.TokenSecret, nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := &model.User{}

		id, ok := claims["userId"].(float64)
		if !ok {
			return fmt.Errorf("verify token: %w", ErrAuthFailed)
		}

		role, ok := claims["role"].(string)
		if !ok {
			return fmt.Errorf("verify token: %w", ErrAuthFailed)
		}

		user.ID = uint(id)
		user.Role = model.Role(role)
		app.setUserSession(user)

		return nil
	}

	return fmt.Errorf("verify token: %w", ErrAuthFailed)
}
