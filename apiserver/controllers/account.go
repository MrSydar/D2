package controllers

import (
	"context"
	"net/http"
	"time"

	"2corp/d2/apiserver/configs"
	"2corp/d2/apiserver/models"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()
var collection = configs.Configs.Database.Collections.Accounts

func CreateAccount(c echo.Context) error {
	var account models.Account

	if err := c.Bind(&account); err != nil {
		msg := "failed to bind account body"
		c.Logger().Error(msg + ": " + err.Error())
		return c.String(http.StatusBadRequest, msg)
	}

	if err := validate.Struct(&account); err != nil {
		msg := "failed to validate account body structure"
		c.Logger().Error(msg + ": " + err.Error())
		return c.String(http.StatusBadRequest, msg)
	}

	account.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, account)
	if err != nil {
		msg := "failed to insert company"
		c.Logger().Error(msg + ": " + err.Error())
		return c.String(http.StatusInternalServerError, msg)
	}

	account.ID = result.InsertedID.(primitive.ObjectID)

	return c.JSON(http.StatusCreated, account)
}

func GetAccount(c echo.Context) error {
	// var item models.Item

	// itemID := c.Param("_id")
	// objId, _ := primitive.ObjectIDFromHex(itemID)

	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	// if err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&item); err != nil {
	// 	if strings.Contains(err.Error(), "no documents in result") {
	// 		return c.String(http.StatusNotFound, "no account was found")
	// 	} else {
	// 		msg := "failed to get account"
	// 		c.Logger().Error(msg + ": " + err.Error())
	// 		return c.String(http.StatusInternalServerError, msg)
	// 	}
	// }

	// return c.JSON(http.StatusOK, item)
	return c.String(http.StatusOK, "your account :)")
}
