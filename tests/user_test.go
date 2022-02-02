package tests

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun/dbfixture"

	"github.com/hhuseyinpay/go-generic-repository-pattern/database"
	"github.com/hhuseyinpay/go-generic-repository-pattern/models"
	"github.com/hhuseyinpay/go-generic-repository-pattern/router"
)

func TestUserHandler_GetByName(t *testing.T) {
	db := database.DB()
	db.RegisterModel((*models.User)(nil))
	fixture := dbfixture.New(db, dbfixture.WithTruncateTables())
	err := fixture.Load(context.Background(), os.DirFS("data"), "user.yml")
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	router.Setup(app, db)

	tests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "create user",
			route:         "/api/users?name=hasan",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "hasan",
		},
	}

	for _, test := range tests {
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		resModel := struct {
			Success bool            `json:"success"`
			Data    [][]models.User `json:"data"`
		}{}
		b, err := ioutil.ReadAll(res.Body)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		err = json.Unmarshal(b, &resModel)
		assert.Nilf(t, err, test.description)

		// Verify, that the reponse body equals the expected body
		assert.Equalf(t, test.expectedBody, resModel.Data[0][0].Name, test.description)
	}
}
