package controllers

import (
	"context"
	"time"

	"2corp/d2/apiserver/configs"
	"2corp/d2/apiserver/models"
	"2corp/d2/apiserver/responses/badrequest"
	"2corp/d2/apiserver/responses/servererror"
	"2corp/d2/apiserver/responses/success"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()
var collection = configs.Configs.Database.Collections.Companies

func CreateCompany(c echo.Context) error {
	var company models.Company

	if err := c.Bind(&company); err != nil {
		return badrequest.BodyValidationFailed(c, err)
	}

	if err := validate.Struct(&company); err != nil {
		return badrequest.FieldValidationFailed(c, err)
	}

	company.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, company)
	if err != nil {
		return servererror.Internal(c, err)
	}

	return success.Created(c, result)
}
