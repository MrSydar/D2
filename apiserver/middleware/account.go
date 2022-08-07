package middleware

import (
	"2corp/d2/apiserver/configs/constants/contextnames"
	"2corp/d2/apiserver/configs/database"
	"2corp/d2/apiserver/configs/log"
	"2corp/d2/apiserver/models"
	"context"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AssociateAccountWithRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		if email == "" {
			return errors.New("email value is empty in the received JWT token")
		}

		account := models.Account{}
		if err := database.Collections.Accounts.FindOne(context.Background(), bson.M{"email": email}).Decode(&account); err != nil {
			if strings.Contains(err.Error(), "no documents in result") { // TODO rewrite
				account.Email = email
				result, err := database.Collections.Accounts.InsertOne(context.Background(), account) // TODO make email field unique in the mongo db
				if err != nil {
					generalErr := errors.New("failed to create new account")
					log.Logger.Errorf("%v: %v", generalErr, err)
					return generalErr
				}
				account.ID = result.InsertedID.(primitive.ObjectID)
			} else {
				generalErr := errors.New("failed to get account")
				log.Logger.Errorf("%v: %v", generalErr, err)
				return generalErr
			}
		}

		c.Set(contextnames.AccountID, account.ID)

		return next(c)
	}
}
