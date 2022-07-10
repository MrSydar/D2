package controllers

// import (
// 	"context"
// 	"time"

// 	"2corp/d2/apiserver/models"
// 	"2corp/d2/apiserver/responses/badrequest"

// 	"github.com/go-playground/validator"
// 	"github.com/labstack/echo"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// var validate = validator.New()

// func CreateCompany(c echo.Context) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	var company models.Company

// 	if err := c.Bind(&company); err != nil {
// 		return badrequest.BodyValidationFailed(c, err)
// 	}

// 	if err := validate.Struct(&company); err != nil {
// 		return badrequest.FieldValidationFailed(c, err)
// 	}

// 	company.ID = primitive.NewObjectID()

// 	// result, err := database.GetCollection(os.Getenv(environment.VariableNames.CompanyCollection))
// }
