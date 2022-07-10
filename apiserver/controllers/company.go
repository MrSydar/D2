package controllers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"2corp/d2/apiserver/configs"
	"2corp/d2/apiserver/models"
	"2corp/d2/apiserver/responses"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()
var collection = configs.Configs.Database.Collections.Companies

func CreateCompany(c echo.Context) error {
	var company models.Company

	if err := c.Bind(&company); err != nil {
		return responses.BodyValidationFailed(c, err)
	}

	if err := validate.Struct(&company); err != nil {
		return responses.FieldValidationFailed(c, err)
	}

	company.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, company)
	if err != nil {
		err = fmt.Errorf("failed to insert company: %v", err)
		c.Logger().Error(err)
		return responses.InternalServerError(c, err)
	}

	return responses.Created(c, result)
}

func GetCompany(c echo.Context) error {
	var company models.Company

	companyId := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(companyId)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&company); err != nil {
		err = fmt.Errorf("failed to find company: %v", err)
		c.Logger().Error(err)

		if strings.Contains(err.Error(), "no documents in result") {
			return responses.NotFound(c, err)
		} else {
			return responses.InternalServerError(c, err)
		}
	}

	return responses.Created(c, company)
}
