package controller_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"test-backend/internal/controller"
	"test-backend/internal/entity"
	"test-backend/internal/model"
	"test-backend/internal/tests"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUserCerts...
func TestGetUser(t *testing.T) {
	testTools := tests.SetupSuite(t, tests.SetupOptions{})
	defer testTools.Teardown(t)

	type input struct {
		UserId string
	}

	type expected struct {
		error          error
		httpStatusCode int
		response       struct {
			Status string      `json:"status"`
			Data   *model.User `json:"data"`
		}
	}
	loc, _ := time.LoadLocation("Asia/Bangkok")
	mockDate := time.Date(2025, 3, 30, 6, 0, 2, 0, loc)
	suite := tests.Testcase{
		"000.Get User.": {
			Mock: &tests.Mock{
				MockFilePaths: []string{"./mocks/user.json"},
			},
			Input: input{
				UserId: "4f6aa9d3-8eca-4045-aa38-4914f8038453",
			},
			Expected: expected{
				error:          nil,
				httpStatusCode: fiber.StatusOK,
				response: entity.GetUserResponse{
					Status: "Success",
					Data: &model.User{
						UserId:    uuid.MustParse("4f6aa9d3-8eca-4045-aa38-4914f8038453"),
						FirstName: "ชัยนา",
						LastName:  "มานอก",
						CreatedAt: &mockDate,
						UpdatedAt: &mockDate,
					},
				},
			},
		},
	}

	for name, tc := range suite {
		t.Run(name, func(t *testing.T) {
			testTools := tests.SetupTest(t, tests.SetupOptions{Mock: tc.Mock})
			defer testTools.Teardown(t)
			testTools.C.RequireInvoke(func(i controller.IUserController) {
				input := tc.Input.(input)
				expected := tc.Expected.(expected)
				app := fiber.New()
				app.Get("/api/v1/user/:user_id", i.GetUser)
				req := httptest.NewRequest(http.MethodGet, "/api/v1/user/"+input.UserId, nil)
				actual, err := app.Test(req)
				if err != nil {
					require.Error(t, err)
					require.Equal(t, expected.httpStatusCode, actual.StatusCode, "HTTP status code mismatch")
				} else {
					defer actual.Body.Close()
					require.NoError(t, err)
					bodyBytes, err := ioutil.ReadAll(actual.Body)
					require.NoError(t, err)
					var actualResponse entity.GetUserResponse
					err = json.Unmarshal(bodyBytes, &actualResponse)
					createdAtStr := actualResponse.Data.CreatedAt.Format("2006-01-02 15:04:05")
					createdAt, err := time.ParseInLocation("2006-01-02 15:04:05", createdAtStr, loc)
					require.NoError(t, err)
					actualResponse.Data.CreatedAt = &createdAt

					updatedAtStr := actualResponse.Data.UpdatedAt.Format("2006-01-02 15:04:05")
					updatedAt, err := time.ParseInLocation("2006-01-02 15:04:05", updatedAtStr, loc)
					require.NoError(t, err)
					actualResponse.Data.UpdatedAt = &updatedAt

					require.NoError(t, err)
					assert.Equal(t, expected.httpStatusCode, actual.StatusCode, "Check HTTP status code")
					assert.Equal(t, expected.response.Status, actualResponse.Status, "Check Response status")
					assert.Equal(t, expected.response.Data, actualResponse.Data, "Check Employee records")
				}
			})
		})
	}
}
